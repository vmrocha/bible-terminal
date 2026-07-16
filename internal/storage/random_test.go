package storage

import (
	"bytes"
	"context"
	"errors"
	"testing"
	"testing/iotest"
)

func TestRandomCanSelectBibleBoundaries(t *testing.T) {
	tests := []struct {
		name    string
		random  []byte
		book    string
		chapter int
		verse   int
	}{
		{
			name:    "first verse",
			random:  []byte{0x00, 0x00},
			book:    "Genesis",
			chapter: 1,
			verse:   1,
		},
		{
			name:    "last verse",
			random:  []byte{0x79, 0x7e},
			book:    "Revelation",
			chapter: 22,
			verse:   21,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader, err := OpenEmbedded(context.Background())
			if err != nil {
				t.Fatalf("OpenEmbedded: %v", err)
			}
			t.Cleanup(func() { _ = reader.Close() })

			passage, err := reader.Random(context.Background(), bytes.NewReader(test.random))
			if err != nil {
				t.Fatalf("Random: %v", err)
			}
			if passage.Book.Name != test.book ||
				passage.Chapter != test.chapter ||
				passage.StartVerse != test.verse ||
				passage.EndVerse != test.verse ||
				len(passage.Verses) != 1 {
				t.Fatalf("unexpected random passage: %#v", passage)
			}
			if passage.Verses[0].Text == "" {
				t.Fatal("random passage has empty text")
			}
		})
	}
}

func TestRandomRejectsMissingSource(t *testing.T) {
	reader, err := OpenEmbedded(context.Background())
	if err != nil {
		t.Fatalf("OpenEmbedded: %v", err)
	}
	t.Cleanup(func() { _ = reader.Close() })

	if _, err := reader.Random(context.Background(), nil); err == nil {
		t.Fatal("Random unexpectedly accepted a nil source")
	}
}

func TestRandomReportsEntropyFailure(t *testing.T) {
	reader, err := OpenEmbedded(context.Background())
	if err != nil {
		t.Fatalf("OpenEmbedded: %v", err)
	}
	t.Cleanup(func() { _ = reader.Close() })

	sourceErr := errors.New("entropy failed")
	_, err = reader.Random(context.Background(), iotest.ErrReader(sourceErr))
	if !errors.Is(err, sourceErr) {
		t.Fatalf("Random error = %v, want %v", err, sourceErr)
	}
}
