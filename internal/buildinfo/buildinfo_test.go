package buildinfo

import "testing"

func TestCurrent(t *testing.T) {
	info := Current()

	if info.Version == "" {
		t.Error("Version must not be empty")
	}
	if info.Commit == "" {
		t.Error("Commit must not be empty")
	}
	if info.Date == "" {
		t.Error("Date must not be empty")
	}
}
