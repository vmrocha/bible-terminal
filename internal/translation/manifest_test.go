package translation

import (
	"strings"
	"testing"
)

const validManifest = `{
  "schema_version": 1,
  "id": "example",
  "name": "Example Bible",
  "abbreviation": "EX",
  "language": {"tag": "en", "name": "English"},
  "edition": "Test Edition",
  "canon": "test",
  "text_edition": "1",
  "source": {
    "publisher": "Example",
    "homepage": "https://example.com",
    "archive_url": "https://example.com/source.zip",
    "archive_member": "source.txt",
    "retrieved_at": "2026-07-15",
    "archive_sha256": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  },
  "rights": {
    "status": "public-domain",
    "notice_url": "https://example.com/rights",
    "trademark": "Example naming notice.",
    "text_policy": "Preserve source text."
  },
  "expected": {
    "books": 1,
    "verses": 1,
    "first_reference": "EXA 1:1",
    "last_reference": "EXA 1:1"
  }
}`

func TestDecodeManifest(t *testing.T) {
	manifest, err := DecodeManifest(strings.NewReader(validManifest))
	if err != nil {
		t.Fatalf("DecodeManifest: %v", err)
	}

	if manifest.ID != "example" {
		t.Fatalf("unexpected manifest ID: %q", manifest.ID)
	}
	if manifest.Expected.Verses != 1 {
		t.Fatalf("unexpected verse total: %d", manifest.Expected.Verses)
	}
}

func TestDecodeManifestRejectsUnknownFields(t *testing.T) {
	input := strings.Replace(validManifest, `"schema_version": 1`, `"schema_version": 1, "unexpected": true`, 1)

	_, err := DecodeManifest(strings.NewReader(input))
	if err == nil || !strings.Contains(err.Error(), "unknown field") {
		t.Fatalf("expected unknown field error, got %v", err)
	}
}

func TestValidateRejectsInvalidChecksum(t *testing.T) {
	manifest, err := DecodeManifest(strings.NewReader(validManifest))
	if err != nil {
		t.Fatalf("DecodeManifest: %v", err)
	}

	manifest.Source.ArchiveSHA256 = "not-a-checksum"
	if err := manifest.Validate(); err == nil {
		t.Fatal("expected invalid checksum to return an error")
	}
}

func TestValidateRejectsMissingProvenance(t *testing.T) {
	manifest, err := DecodeManifest(strings.NewReader(validManifest))
	if err != nil {
		t.Fatalf("DecodeManifest: %v", err)
	}

	manifest.Rights.NoticeURL = ""
	if err := manifest.Validate(); err == nil {
		t.Fatal("expected missing rights notice to return an error")
	}
}

func TestValidateRejectsInsecureSourceURL(t *testing.T) {
	manifest, err := DecodeManifest(strings.NewReader(validManifest))
	if err != nil {
		t.Fatalf("DecodeManifest: %v", err)
	}

	manifest.Source.ArchiveURL = "http://example.com/source.zip"
	if err := manifest.Validate(); err == nil {
		t.Fatal("expected insecure source URL to return an error")
	}
}
