package bible

// Dataset is the normalized content produced by a translation importer.
type Dataset struct {
	Books  []Book
	Verses []Verse
}

// Book identifies one book within a translation and its canonical position.
type Book struct {
	ID         string
	SourceCode string
	Position   int
	Name       string
}

// Verse is one addressable verse with publisher-provided text.
type Verse struct {
	BookID  string
	Chapter int
	Number  int
	Text    string
}
