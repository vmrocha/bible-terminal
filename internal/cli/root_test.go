package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/vmrocha/bible-terminal/internal/buildinfo"
)

var testBuild = buildinfo.Info{
	Version: "v0.1.0-test",
	Commit:  "abc1234",
	Date:    "2026-07-15T12:00:00Z",
}

func execute(t *testing.T, args ...string) (string, error) {
	return executeWithOptions(t, nil, args...)
}

func executeWithOptions(t *testing.T, options []Option, args ...string) (string, error) {
	t.Helper()

	output := new(bytes.Buffer)
	command := New(testBuild, options...)
	command.SetOut(output)
	command.SetErr(output)
	command.SetArgs(args)

	err := command.Execute()
	return output.String(), err
}

func TestHelp(t *testing.T) {
	output, err := execute(t, "--help")
	if err != nil {
		t.Fatalf("execute --help: %v", err)
	}

	for _, expected := range []string{
		"Read the Bible from your terminal",
		"read",
		"version",
		"--help",
	} {
		if !strings.Contains(output, expected) {
			t.Errorf("help output does not contain %q", expected)
		}
	}
}

func TestVersionFlag(t *testing.T) {
	output, err := execute(t, "--version")
	if err != nil {
		t.Fatalf("execute --version: %v", err)
	}

	if output != "bible v0.1.0-test\n" {
		t.Fatalf("unexpected version output: %q", output)
	}
}

func TestUnknownCommand(t *testing.T) {
	_, err := execute(t, "unknown")
	if err == nil {
		t.Fatal("expected unknown command to return an error")
	}
}
