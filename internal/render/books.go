package render

import (
	"fmt"
	"io"
	"strings"

	"github.com/vmrocha/bible-terminal/internal/canon"
)

// Books writes the accepted book names, source codes, and aliases.
func Books(writer io.Writer, entries []canon.Entry, plain bool) error {
	for _, entry := range entries {
		aliases := strings.Join(entry.Aliases, ", ")
		if plain {
			if _, err := fmt.Fprintf(
				writer,
				"%s\t%s\t%s\t%s\n",
				entry.Book.ID,
				entry.Book.Name,
				entry.Book.SourceCode,
				strings.Join(entry.Aliases, ","),
			); err != nil {
				return err
			}
			continue
		}

		if entry.Book.Position == 1 {
			if _, err := fmt.Fprintln(writer, "Old Testament"); err != nil {
				return err
			}
		}
		if entry.Book.Position == 40 {
			if _, err := fmt.Fprintln(writer, "\nNew Testament"); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintf(
			writer,
			"%2d  %-17s %-3s  %s\n",
			entry.Book.Position,
			entry.Book.Name,
			entry.Book.SourceCode,
			aliases,
		); err != nil {
			return err
		}
	}
	return nil
}
