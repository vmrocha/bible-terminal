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
	query   reference.Query
	next    reference.Query
}

func (reader *stubReader) Read(_ context.Context, query reference.Query) (bible.Passage, error) {
	reader.query = query
	return reader.passage, reader.err
}

func (reader *stubReader) Navigate(context.Context, reference.Query, int) (reference.Query, error) {
	return reader.next, nil
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

func TestReadCommandNavigates(t *testing.T) {
	reader := &stubReader{
		next: reference.Query{Book: "john", Chapter: 4},
		passage: bible.Passage{
			Translation: "WEBP",
			Book:        bible.Book{Name: "John"},
			Chapter:     4,
			Verses:      []bible.Verse{{Chapter: 4, Number: 1, Text: "After these things…"}},
		},
	}
	factory := func(context.Context) (PassageReader, error) { return reader, nil }

	_, err := executeWithOptions(t, []Option{WithReaderFactory(factory)}, "read", "John", "3", "--next")
	if err != nil {
		t.Fatalf("execute read --next: %v", err)
	}
	if reader.query.Book != "john" || reader.query.Chapter != 4 {
		t.Fatalf("read query is %#v", reader.query)
	}
}

func TestReadCommandRejectsInvalidNavigationFlags(t *testing.T) {
	for _, args := range [][]string{
		{"read", "John", "3", "--next", "--previous"},
		{"read", "John", "3:16", "--next"},
	} {
		_, err := executeWithOptions(t, nil, args...)
		if err == nil {
			t.Fatalf("execute %v unexpectedly succeeded", args)
		}
	}
}
