package render

import (
	"bytes"
	"strings"
	"testing"

	"github.com/vmrocha/bible-terminal/internal/canon"
)

func TestBooks(t *testing.T) {
	var output bytes.Buffer
	if err := Books(&output, canon.ProtestantBooks(), false); err != nil {
		t.Fatalf("Books: %v", err)
	}

	for _, expected := range []string{
		"Old Testament",
		" 1  Genesis",
		"39  Malachi",
		"New Testament",
		"40  Matthew",
		"66  Revelation",
		"Jn, Jhn",
	} {
		if !strings.Contains(output.String(), expected) {
			t.Errorf("book output does not contain %q", expected)
		}
	}
}

func TestPlainBooks(t *testing.T) {
	var output bytes.Buffer
	if err := Books(&output, canon.ProtestantBooks(), true); err != nil {
		t.Fatalf("Books: %v", err)
	}
	if !strings.Contains(output.String(), "john\tJohn\tJOH\tJn,Jhn\n") {
		t.Fatalf("unexpected plain output:\n%s", output.String())
	}
}
