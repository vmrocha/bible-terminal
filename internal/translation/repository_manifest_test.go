package translation

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestRepositoryManifest(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate test source")
	}

	path := filepath.Join(filepath.Dir(filename), "..", "..", "data", "translations", "engwebp", "manifest.json")
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("open repository manifest: %v", err)
	}
	t.Cleanup(func() { _ = file.Close() })

	manifest, err := DecodeManifest(file)
	if err != nil {
		t.Fatalf("DecodeManifest: %v", err)
	}

	if manifest.ID != "engwebp" {
		t.Fatalf("unexpected manifest ID: %q", manifest.ID)
	}
	if manifest.Expected.Books != 66 || manifest.Expected.Verses != 31103 {
		t.Fatalf(
			"unexpected WEBP structure: %d books, %d verses",
			manifest.Expected.Books,
			manifest.Expected.Verses,
		)
	}
}
