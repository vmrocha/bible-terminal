package storage

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/vmrocha/bible-terminal/internal/reference"
)

func TestEmbeddedDatabase(t *testing.T) {
	digest := sha256.Sum256(embeddedWEBP)
	got := hex.EncodeToString(digest[:])
	const want = "1164fb6d921762958f99c904742ff782d3bf5948d64e82bdcb5c677176ccd337"
	if got != want {
		t.Fatalf("embedded database checksum is %s, want %s", got, want)
	}

	reader, err := OpenEmbedded(context.Background())
	if err != nil {
		t.Fatalf("OpenEmbedded: %v", err)
	}
	t.Cleanup(func() { _ = reader.Close() })

	var books, verses int
	if err := reader.connection.QueryRowContext(
		context.Background(),
		"SELECT (SELECT count(*) FROM books), (SELECT count(*) FROM verses)",
	).Scan(&books, &verses); err != nil {
		t.Fatalf("query embedded totals: %v", err)
	}
	if books != 66 || verses != 31103 {
		t.Fatalf("embedded totals are %d books and %d verses", books, verses)
	}

	if _, err := reader.connection.ExecContext(context.Background(), "DELETE FROM verses"); err == nil {
		t.Fatal("embedded database unexpectedly allowed a write")
	}
}

func TestRead(t *testing.T) {
	reader, err := OpenEmbedded(context.Background())
	if err != nil {
		t.Fatalf("OpenEmbedded: %v", err)
	}
	t.Cleanup(func() { _ = reader.Close() })

	query, err := reference.Parse("Jn 3:16")
	if err != nil {
		t.Fatalf("parse reference: %v", err)
	}
	passage, err := reader.Read(context.Background(), query)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if passage.Book.Name != "John" || passage.Translation != "WEBP" {
		t.Fatalf("unexpected passage metadata: %#v", passage)
	}
	if len(passage.Verses) != 1 || !strings.HasPrefix(passage.Verses[0].Text, "For God so loved") {
		t.Fatalf("unexpected passage verses: %#v", passage.Verses)
	}
}

func TestReadRejectsMissingPassages(t *testing.T) {
	reader, err := OpenEmbedded(context.Background())
	if err != nil {
		t.Fatalf("OpenEmbedded: %v", err)
	}
	t.Cleanup(func() { _ = reader.Close() })

	for _, input := range []string{"Notabook 1", "John 999", "John 3:999", "John 3:35-40"} {
		t.Run(input, func(t *testing.T) {
			query, err := reference.Parse(input)
			if err != nil {
				t.Fatalf("parse reference: %v", err)
			}
			if _, err := reader.Read(context.Background(), query); err == nil {
				t.Fatalf("Read(%q) unexpectedly succeeded", input)
			}
		})
	}
}
