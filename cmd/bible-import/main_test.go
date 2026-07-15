package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	output := new(bytes.Buffer)
	command := newCommand()
	command.SetOut(output)
	command.SetArgs([]string{"--help"})

	if err := command.Execute(); err != nil {
		t.Fatalf("execute --help: %v", err)
	}
	for _, expected := range []string{"--manifest", "--archive", "--output"} {
		if !strings.Contains(output.String(), expected) {
			t.Errorf("help output does not contain %q", expected)
		}
	}
}

func TestRequiredFlags(t *testing.T) {
	command := newCommand()
	command.SetArgs(nil)

	err := command.Execute()
	if err == nil || !strings.Contains(err.Error(), "are required") {
		t.Fatalf("expected required flags error, got %v", err)
	}
}

func TestRejectsArguments(t *testing.T) {
	command := newCommand()
	command.SetArgs([]string{"unexpected"})

	err := command.Execute()
	if err == nil || !strings.Contains(err.Error(), "unknown command") {
		t.Fatalf("expected argument error, got %v", err)
	}
}
