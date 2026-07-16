package storage

import (
	"context"
	"testing"
)

func TestTranslations(t *testing.T) {
	reader, err := OpenEmbedded(context.Background())
	if err != nil {
		t.Fatalf("OpenEmbedded: %v", err)
	}
	t.Cleanup(func() { _ = reader.Close() })

	translations, err := reader.Translations(context.Background())
	if err != nil {
		t.Fatalf("Translations: %v", err)
	}
	if len(translations) != 1 {
		t.Fatalf("Translations returned %d entries, want 1", len(translations))
	}
	got := translations[0]
	if got.ID != "engwebp" ||
		got.Abbreviation != "WEBP" ||
		got.Name != "World English Bible" ||
		got.LanguageTag != "en-US" ||
		got.Edition != "Protestant Edition" ||
		got.RightsStatus != "public-domain" {
		t.Fatalf("unexpected translation: %#v", got)
	}
	if got.SourcePublisher == "" ||
		got.SourceHomepage == "" ||
		got.RightsNoticeURL == "" ||
		got.TrademarkNotice == "" ||
		got.TextPolicy == "" {
		t.Fatalf("translation attribution is incomplete: %#v", got)
	}
}
