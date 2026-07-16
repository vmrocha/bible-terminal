package storage

import (
	"context"
	"strings"
	"testing"
)

func TestSearchExpression(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"living water", "\"living\" AND \"water\""},
		{"  Faith, hope & LOVE! ", "\"Faith\" AND \"hope\" AND \"LOVE\""},
		{"king's", "\"king\" AND \"s\""},
		{"água viva", "\"água\" AND \"viva\""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got, err := searchExpression(test.input)
			if err != nil {
				t.Fatalf("searchExpression: %v", err)
			}
			if got != test.want {
				t.Fatalf("searchExpression = %q, want %q", got, test.want)
			}
		})
	}
}

func TestSearchExpressionRejectsPunctuationOnly(t *testing.T) {
	if _, err := searchExpression(" -- !!! "); err == nil {
		t.Fatal("searchExpression unexpectedly succeeded")
	}
}

func TestSearch(t *testing.T) {
	reader, err := OpenEmbedded(context.Background())
	if err != nil {
		t.Fatalf("OpenEmbedded: %v", err)
	}
	t.Cleanup(func() { _ = reader.Close() })

	results, err := reader.Search(context.Background(), "faith hope love", 2)
	if err != nil {
		t.Fatalf("Search: %v", err)
	}
	if len(results) != 2 {
		t.Fatalf("Search returned %d results, want 2", len(results))
	}
	first := results[0]
	if first.Book.Name != "1 Corinthians" || first.Chapter != 13 || first.Verse != 13 {
		t.Fatalf("first result is %#v", first)
	}
	for _, result := range results {
		text := strings.ToLower(result.Text)
		for _, token := range []string{"faith", "hope", "love"} {
			if !strings.Contains(text, token) {
				t.Errorf("%s %d:%d does not contain %q", result.Book.Name, result.Chapter, result.Verse, token)
			}
		}
	}
}

func TestSearchReturnsNoResults(t *testing.T) {
	reader, err := OpenEmbedded(context.Background())
	if err != nil {
		t.Fatalf("OpenEmbedded: %v", err)
	}
	t.Cleanup(func() { _ = reader.Close() })

	results, err := reader.Search(context.Background(), "zyxwvutsrqponmlkjihgfedcba", 20)
	if err != nil {
		t.Fatalf("Search: %v", err)
	}
	if len(results) != 0 {
		t.Fatalf("Search returned %d results, want none", len(results))
	}
}

func TestSearchRejectsInvalidLimit(t *testing.T) {
	reader, err := OpenEmbedded(context.Background())
	if err != nil {
		t.Fatalf("OpenEmbedded: %v", err)
	}
	t.Cleanup(func() { _ = reader.Close() })

	if _, err := reader.Search(context.Background(), "water", 0); err == nil {
		t.Fatal("Search unexpectedly accepted a zero limit")
	}
}

func BenchmarkSearch(b *testing.B) {
	reader, err := OpenEmbedded(context.Background())
	if err != nil {
		b.Fatalf("OpenEmbedded: %v", err)
	}
	b.Cleanup(func() { _ = reader.Close() })

	for b.Loop() {
		if _, err := reader.Search(context.Background(), "living water", 20); err != nil {
			b.Fatalf("Search: %v", err)
		}
	}
}
