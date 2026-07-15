package translation

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"time"
)

// Manifest records the identity, provenance, rights, and validation
// expectations for one translation source snapshot.
type Manifest struct {
	SchemaVersion int      `json:"schema_version"`
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Abbreviation  string   `json:"abbreviation"`
	Language      Language `json:"language"`
	Edition       string   `json:"edition"`
	Canon         string   `json:"canon"`
	TextEdition   string   `json:"text_edition"`
	Source        Source   `json:"source"`
	Rights        Rights   `json:"rights"`
	Expected      Expected `json:"expected"`
}

// Language identifies the translation language and dialect.
type Language struct {
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

// Source identifies the exact publisher artifact selected for import.
type Source struct {
	Publisher     string `json:"publisher"`
	Homepage      string `json:"homepage"`
	ArchiveURL    string `json:"archive_url"`
	ArchiveMember string `json:"archive_member"`
	RetrievedAt   string `json:"retrieved_at"`
	ArchiveSHA256 string `json:"archive_sha256"`
}

// Rights records the redistribution status and naming constraints.
type Rights struct {
	Status     string `json:"status"`
	NoticeURL  string `json:"notice_url"`
	Trademark  string `json:"trademark"`
	TextPolicy string `json:"text_policy"`
}

// Expected contains structural assertions applied during import.
type Expected struct {
	Books          int    `json:"books"`
	Verses         int    `json:"verses"`
	FirstReference string `json:"first_reference"`
	LastReference  string `json:"last_reference"`
}

// DecodeManifest parses and validates one manifest document.
func DecodeManifest(reader io.Reader) (Manifest, error) {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	var manifest Manifest
	if err := decoder.Decode(&manifest); err != nil {
		return Manifest{}, fmt.Errorf("decode translation manifest: %w", err)
	}

	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		if err == nil {
			return Manifest{}, errors.New("decode translation manifest: multiple JSON values")
		}
		return Manifest{}, fmt.Errorf("decode translation manifest trailer: %w", err)
	}

	if err := manifest.Validate(); err != nil {
		return Manifest{}, err
	}

	return manifest, nil
}

// Validate verifies the provenance fields required before importing text.
func (manifest Manifest) Validate() error {
	if manifest.SchemaVersion != 1 {
		return fmt.Errorf("validate translation manifest: unsupported schema version %d", manifest.SchemaVersion)
	}

	required := []struct {
		name  string
		value string
	}{
		{"id", manifest.ID},
		{"name", manifest.Name},
		{"abbreviation", manifest.Abbreviation},
		{"language.tag", manifest.Language.Tag},
		{"language.name", manifest.Language.Name},
		{"edition", manifest.Edition},
		{"canon", manifest.Canon},
		{"text_edition", manifest.TextEdition},
		{"source.publisher", manifest.Source.Publisher},
		{"source.homepage", manifest.Source.Homepage},
		{"source.archive_url", manifest.Source.ArchiveURL},
		{"source.archive_member", manifest.Source.ArchiveMember},
		{"source.retrieved_at", manifest.Source.RetrievedAt},
		{"rights.status", manifest.Rights.Status},
		{"rights.notice_url", manifest.Rights.NoticeURL},
		{"rights.trademark", manifest.Rights.Trademark},
		{"rights.text_policy", manifest.Rights.TextPolicy},
		{"expected.first_reference", manifest.Expected.FirstReference},
		{"expected.last_reference", manifest.Expected.LastReference},
	}
	for _, field := range required {
		if field.value == "" {
			return fmt.Errorf("validate translation manifest: %s is required", field.name)
		}
	}

	urls := []struct {
		name  string
		value string
	}{
		{"source.homepage", manifest.Source.Homepage},
		{"source.archive_url", manifest.Source.ArchiveURL},
		{"rights.notice_url", manifest.Rights.NoticeURL},
	}
	for _, candidate := range urls {
		parsed, err := url.ParseRequestURI(candidate.value)
		if err != nil || parsed.Scheme != "https" || parsed.Host == "" {
			return fmt.Errorf("validate translation manifest: %s must be an absolute HTTPS URL", candidate.name)
		}
	}

	if _, err := time.Parse(time.DateOnly, manifest.Source.RetrievedAt); err != nil {
		return fmt.Errorf("validate translation manifest: source.retrieved_at: %w", err)
	}

	digest, err := hex.DecodeString(manifest.Source.ArchiveSHA256)
	if err != nil || len(digest) != 32 {
		return errors.New("validate translation manifest: source.archive_sha256 must be 64 hexadecimal characters")
	}

	if manifest.Expected.Books <= 0 {
		return errors.New("validate translation manifest: expected.books must be positive")
	}
	if manifest.Expected.Verses <= 0 {
		return errors.New("validate translation manifest: expected.verses must be positive")
	}

	return nil
}
