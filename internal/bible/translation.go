package bible

// Translation describes one Bible text available to the reader.
type Translation struct {
	ID              string
	Name            string
	Abbreviation    string
	LanguageTag     string
	LanguageName    string
	Edition         string
	Canon           string
	TextEdition     string
	SourcePublisher string
	SourceHomepage  string
	RightsStatus    string
	RightsNoticeURL string
	TrademarkNotice string
	TextPolicy      string
}
