package cli

import (
	"strings"
	"testing"
)

func TestBooksCommand(t *testing.T) {
	output, err := execute(t, "books", "--no-color")
	if err != nil {
		t.Fatalf("execute books: %v", err)
	}
	if !strings.Contains(output, "Old Testament") || !strings.Contains(output, "Revelation") {
		t.Fatalf("unexpected books output:\n%s", output)
	}
}

func TestBooksCommandPlain(t *testing.T) {
	output, err := execute(t, "--plain", "books")
	if err != nil {
		t.Fatalf("execute books --plain: %v", err)
	}
	if !strings.Contains(output, "john\tJohn\tJOH\tJn,Jhn") {
		t.Fatalf("unexpected books output:\n%s", output)
	}
}
