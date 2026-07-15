package cli

import "testing"

func TestVersionCommand(t *testing.T) {
	output, err := execute(t, "version")
	if err != nil {
		t.Fatalf("execute version: %v", err)
	}

	expected := "bible v0.1.0-test\ncommit: abc1234\nbuilt: 2026-07-15T12:00:00Z\n"
	if output != expected {
		t.Fatalf("unexpected version output:\n%s", output)
	}
}

func TestVersionCommandRejectsArguments(t *testing.T) {
	_, err := execute(t, "version", "extra")
	if err == nil {
		t.Fatal("expected version arguments to return an error")
	}
}
