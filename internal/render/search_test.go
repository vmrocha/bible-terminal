package render

import (
	"bytes"
	"testing"

	"github.com/vmrocha/bible-terminal/internal/bible"
)

func TestSearch(t *testing.T) {
	results := []bible.SearchResult{
		{
			Translation: "WEBP",
			Book:        bible.Book{Name: "John"},
			Chapter:     4,
			Verse:       10,
			Text:        "He would have given you living water.",
		},
		{
			Translation: "WEBP",
			Book:        bible.Book{Name: "John"},
			Chapter:     7,
			Verse:       38,
			Text:        "Rivers of living water will flow.",
		},
	}

	var output bytes.Buffer
	if err := Search(&output, "living water", results, Options{}); err != nil {
		t.Fatalf("Search: %v", err)
	}
	want := "Search results · WEBP\n2 matches for \"living water\"\n\n" +
		"John 4:10  He would have given you living water.\n" +
		"John 7:38  Rivers of living water will flow.\n"
	if output.String() != want {
		t.Fatalf("unexpected search output:\n%s", output.String())
	}
}

func TestStyledSearch(t *testing.T) {
	results := []bible.SearchResult{{
		Translation: "WEBP",
		Book:        bible.Book{Name: "John"},
		Chapter:     3,
		Verse:       16,
		Text:        "For God so loved the world.",
	}}

	var output bytes.Buffer
	if err := Search(&output, "God loved", results, Options{Color: true}); err != nil {
		t.Fatalf("Search: %v", err)
	}
	want := "\x1b[1m\x1b[36mSearch results\x1b[0m · \x1b[2mWEBP\x1b[0m\n" +
		"1 match for \"God loved\"\n\n" +
		"\x1b[1m\x1b[36mJohn 3:16\x1b[0m  For God so loved the world.\n"
	if output.String() != want {
		t.Fatalf("unexpected styled search output:\n%q", output.String())
	}
}

func TestPlainSearch(t *testing.T) {
	results := []bible.SearchResult{{
		Book:    bible.Book{Name: "Psalm"},
		Chapter: 23,
		Verse:   1,
		Text:    "Yahweh is my shepherd.",
	}}

	var output bytes.Buffer
	if err := Search(&output, "shepherd", results, Options{Plain: true}); err != nil {
		t.Fatalf("Search: %v", err)
	}
	if want := "Psalm 23:1\tYahweh is my shepherd.\n"; output.String() != want {
		t.Fatalf("unexpected plain search output: %q", output.String())
	}
}

func TestSearchWithNoResults(t *testing.T) {
	var output bytes.Buffer
	if err := Search(&output, "not found", nil, Options{}); err != nil {
		t.Fatalf("Search: %v", err)
	}
	if want := "No results for \"not found\".\n"; output.String() != want {
		t.Fatalf("unexpected empty search output: %q", output.String())
	}

	output.Reset()
	if err := Search(&output, "not found", nil, Options{Plain: true}); err != nil {
		t.Fatalf("plain Search: %v", err)
	}
	if output.Len() != 0 {
		t.Fatalf("plain empty search output is not empty: %q", output.String())
	}
}
