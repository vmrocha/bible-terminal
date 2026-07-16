package bible

// SearchResult is one verse returned by a full-text search.
type SearchResult struct {
	Translation string
	Book        Book
	Chapter     int
	Verse       int
	Text        string
	Highlights  []TextRange
}

// TextRange identifies a half-open UTF-8 byte range within a result's Text.
type TextRange struct {
	Start int
	End   int
}
