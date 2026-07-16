package render

import (
	"fmt"
	"io"

	"github.com/vmrocha/bible-terminal/internal/bible"
)

// Search writes ranked verse matches in human-readable or stable plain form.
func Search(writer io.Writer, query string, results []bible.SearchResult, options Options) error {
	if options.Plain {
		for _, result := range results {
			if _, err := fmt.Fprintf(
				writer,
				"%s %d:%d\t%s\n",
				result.Book.Name,
				result.Chapter,
				result.Verse,
				result.Text,
			); err != nil {
				return err
			}
		}
		return nil
	}

	if len(results) == 0 {
		_, err := fmt.Fprintf(writer, "No results for %q.\n", query)
		return err
	}

	title := styled("Search results", ansiBold+ansiCyan, options.Color)
	translation := styled(results[0].Translation, ansiDim, options.Color)
	if _, err := fmt.Fprintf(writer, "%s · %s\n", title, translation); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(writer, "%d %s for %q\n\n", len(results), matchLabel(len(results)), query); err != nil {
		return err
	}
	for _, result := range results {
		reference := fmt.Sprintf("%s %d:%d", result.Book.Name, result.Chapter, result.Verse)
		reference = styled(reference, ansiBold+ansiCyan, options.Color)
		if _, err := fmt.Fprintf(writer, "%s  %s\n", reference, result.Text); err != nil {
			return err
		}
	}
	return nil
}

func matchLabel(count int) string {
	if count == 1 {
		return "match"
	}
	return "matches"
}
