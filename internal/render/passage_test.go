package render

import (
	"bytes"
	"testing"

	"github.com/vmrocha/bible-terminal/internal/bible"
)

func TestPassage(t *testing.T) {
	passage := bible.Passage{
		Translation: "WEBP",
		Book:        bible.Book{Name: "John"},
		Chapter:     3,
		StartVerse:  16,
		EndVerse:    17,
		Verses: []bible.Verse{
			{Chapter: 3, Number: 16, Text: "For God so loved…"},
			{Chapter: 3, Number: 17, Text: "For God didn’t send…"},
		},
	}

	var output bytes.Buffer
	if err := Passage(&output, passage, false); err != nil {
		t.Fatalf("Passage: %v", err)
	}
	want := "John 3:16–17 · WEBP\n\n16  For God so loved…\n17  For God didn’t send…\n"
	if output.String() != want {
		t.Fatalf("unexpected output:\n%s", output.String())
	}
}

func TestPlainPassage(t *testing.T) {
	passage := bible.Passage{
		Book:    bible.Book{Name: "Luke"},
		Chapter: 17,
		Verses: []bible.Verse{
			{Chapter: 17, Number: 36, Text: ""},
			{Chapter: 17, Number: 37, Text: "They asked…"},
		},
	}

	var output bytes.Buffer
	if err := Passage(&output, passage, true); err != nil {
		t.Fatalf("Passage: %v", err)
	}
	want := "Luke 17:36\t\nLuke 17:37\tThey asked…\n"
	if output.String() != want {
		t.Fatalf("unexpected output:\n%s", output.String())
	}
}

func TestPassageRejectsEmptyContent(t *testing.T) {
	var output bytes.Buffer
	if err := Passage(&output, bible.Passage{}, false); err == nil {
		t.Fatal("expected empty passage to return an error")
	}
}
