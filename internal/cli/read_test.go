package cli

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/vmrocha/bible-terminal/internal/bible"
	"github.com/vmrocha/bible-terminal/internal/reference"
)

type stubReader struct {
	passage bible.Passage
	err     error
	closed  bool
}

func (reader *stubReader) Read(context.Context, reference.Query) (bible.Passage, error) {
	return reader.passage, reader.err
}

func (reader *stubReader) Close() error {
	reader.closed = true
	return nil
}

func TestReadCommand(t *testing.T) {
	reader := &stubReader{passage: bible.Passage{
		Translation: "WEBP",
		Book:        bible.Book{Name: "John"},
		Chapter:     3,
		StartVerse:  16,
		EndVerse:    16,
		Verses: []bible.Verse{
			{Chapter: 3, Number: 16, Text: "For God so loved…"},
		},
	}}
	factory := func(context.Context) (PassageReader, error) { return reader, nil }

	output, err := executeWithOptions(t, []Option{WithReaderFactory(factory)}, "read", "John", "3:16")
	if err != nil {
		t.Fatalf("execute read: %v", err)
	}
	if !strings.Contains(output, "John 3:16 · WEBP") || !strings.Contains(output, "For God so loved") {
		t.Fatalf("unexpected read output:\n%s", output)
	}
	if !reader.closed {
		t.Fatal("reader was not closed")
	}
}

func TestReadCommandReportsReaderError(t *testing.T) {
	reader := &stubReader{err: errors.New("passage failed")}
	factory := func(context.Context) (PassageReader, error) { return reader, nil }

	_, err := executeWithOptions(t, []Option{WithReaderFactory(factory)}, "read", "John", "3")
	if err == nil || !strings.Contains(err.Error(), "passage failed") {
		t.Fatalf("expected passage error, got %v", err)
	}
	if !reader.closed {
		t.Fatal("reader was not closed after an error")
	}
}
