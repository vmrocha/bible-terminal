package bible

// SearchResult is one verse returned by a full-text search.
type SearchResult struct {
	Translation string
	Book        Book
	Chapter     int
	Verse       int
	Text        string
}
