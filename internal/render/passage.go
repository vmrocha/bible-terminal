package render

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/vmrocha/bible-terminal/internal/bible"
)

// Passage writes a Bible passage in human-readable or stable plain form.
func Passage(writer io.Writer, passage bible.Passage, plain bool) error {
	if len(passage.Verses) == 0 {
		return errors.New("render passage: no verses")
	}
	if plain {
		for _, verse := range passage.Verses {
			if _, err := fmt.Fprintf(
				writer,
				"%s %d:%d\t%s\n",
				passage.Book.Name,
				verse.Chapter,
				verse.Number,
				verse.Text,
			); err != nil {
				return err
			}
		}
		return nil
	}

	if _, err := fmt.Fprintf(writer, "%s · %s\n\n", passageTitle(passage), passage.Translation); err != nil {
		return err
	}
	width := len(strconv.Itoa(passage.Verses[len(passage.Verses)-1].Number))
	for _, verse := range passage.Verses {
		if _, err := fmt.Fprintf(writer, "%*d  %s\n", width, verse.Number, verse.Text); err != nil {
			return err
		}
	}
	return nil
}

func passageTitle(passage bible.Passage) string {
	if passage.StartVerse == 0 {
		return fmt.Sprintf("%s %d", passage.Book.Name, passage.Chapter)
	}
	if passage.StartVerse == passage.EndVerse {
		return fmt.Sprintf("%s %d:%d", passage.Book.Name, passage.Chapter, passage.StartVerse)
	}
	return fmt.Sprintf(
		"%s %d:%d–%d",
		passage.Book.Name,
		passage.Chapter,
		passage.StartVerse,
		passage.EndVerse,
	)
}
