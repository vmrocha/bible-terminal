package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStoreReturnsDefaultsWhenMissing(t *testing.T) {
	store := newTestStore(t)
	got, err := store.Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got != Defaults() {
		t.Fatalf("Load = %#v, want %#v", got, Defaults())
	}
}

func TestStoreSavesAndLoadsPreferences(t *testing.T) {
	store := newTestStore(t)
	want := Preferences{
		Version:     1,
		Translation: "engwebp",
		Plain:       true,
		Color:       false,
	}
	if err := store.Save(want); err != nil {
		t.Fatalf("Save: %v", err)
	}
	got, err := store.Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got != want {
		t.Fatalf("Load = %#v, want %#v", got, want)
	}

	info, err := os.Stat(store.Path())
	if err != nil {
		t.Fatalf("stat configuration: %v", err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("configuration permissions are %o, want 600", info.Mode().Perm())
	}
	directory, err := os.Stat(filepath.Dir(store.Path()))
	if err != nil {
		t.Fatalf("stat configuration directory: %v", err)
	}
	if directory.Mode().Perm() != 0o700 {
		t.Fatalf("configuration directory permissions are %o, want 700", directory.Mode().Perm())
	}
}

func TestStoreUsesDefaultsForOmittedOptionalFields(t *testing.T) {
	store := newTestStore(t)
	if err := os.MkdirAll(filepath.Dir(store.Path()), 0o700); err != nil {
		t.Fatalf("create directory: %v", err)
	}
	if err := os.WriteFile(store.Path(), []byte("{\"version\":1}\n"), 0o600); err != nil {
		t.Fatalf("write configuration: %v", err)
	}

	got, err := store.Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got != Defaults() {
		t.Fatalf("Load = %#v, want %#v", got, Defaults())
	}
}

func TestStoreRejectsInvalidDocuments(t *testing.T) {
	tests := []struct {
		name     string
		document string
		want     string
	}{
		{"malformed JSON", "{", "decode configuration"},
		{"unknown field", "{\"version\":1,\"unknown\":true}", "unknown field"},
		{"missing version", "{}", "version is required"},
		{"multiple values", "{\"version\":1} {}", "multiple JSON values"},
		{"future version", "{\"version\":2}", "unsupported configuration version"},
		{"unknown translation", "{\"version\":1,\"translation\":\"other\"}", "translation is not available"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := newTestStore(t)
			if err := os.MkdirAll(filepath.Dir(store.Path()), 0o700); err != nil {
				t.Fatalf("create directory: %v", err)
			}
			if err := os.WriteFile(store.Path(), []byte(test.document), 0o600); err != nil {
				t.Fatalf("write configuration: %v", err)
			}

			_, err := store.Load()
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("expected error containing %q, got %v", test.want, err)
			}
		})
	}
}

func TestStoreRejectsInvalidPreferencesWithoutChangingFile(t *testing.T) {
	store := newTestStore(t)
	if err := store.Save(Defaults()); err != nil {
		t.Fatalf("Save defaults: %v", err)
	}
	before, err := os.ReadFile(store.Path())
	if err != nil {
		t.Fatalf("read original configuration: %v", err)
	}

	invalid := Defaults()
	invalid.Translation = "other"
	if err := store.Save(invalid); err == nil {
		t.Fatal("Save unexpectedly accepted invalid preferences")
	}
	after, err := os.ReadFile(store.Path())
	if err != nil {
		t.Fatalf("read preserved configuration: %v", err)
	}
	if string(after) != string(before) {
		t.Fatal("invalid save changed the existing configuration")
	}
}

func TestStoreResetRemovesValidOrCorruptConfiguration(t *testing.T) {
	store := newTestStore(t)
	if err := os.MkdirAll(filepath.Dir(store.Path()), 0o700); err != nil {
		t.Fatalf("create directory: %v", err)
	}
	if err := os.WriteFile(store.Path(), []byte("{"), 0o600); err != nil {
		t.Fatalf("write configuration: %v", err)
	}
	if err := store.Reset(); err != nil {
		t.Fatalf("Reset: %v", err)
	}
	if _, err := os.Stat(store.Path()); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("configuration still exists: %v", err)
	}
	if err := store.Reset(); err != nil {
		t.Fatalf("second Reset: %v", err)
	}
}

func TestNewStoreRequiresAbsolutePath(t *testing.T) {
	if _, err := NewStore("relative/config.json"); err == nil {
		t.Fatal("NewStore unexpectedly accepted a relative path")
	}
}

func newTestStore(t *testing.T) *Store {
	t.Helper()
	store, err := NewStore(filepath.Join(t.TempDir(), "nested", "config.json"))
	if err != nil {
		t.Fatalf("NewStore: %v", err)
	}
	return store
}
