package cli

import (
	"strings"
	"testing"
)

func TestCompletionCommand(t *testing.T) {
	tests := []struct {
		shell string
		want  string
	}{
		{"bash", "__start_bible"},
		{"zsh", "#compdef bible"},
		{"fish", "complete -c bible"},
		{"powershell", "Register-ArgumentCompleter"},
	}

	for _, test := range tests {
		t.Run(test.shell, func(t *testing.T) {
			output, err := execute(t, "completion", test.shell)
			if err != nil {
				t.Fatalf("execute completion %s: %v", test.shell, err)
			}
			if !strings.Contains(output, test.want) {
				t.Fatalf("%s completion does not contain %q", test.shell, test.want)
			}
			if strings.Contains(output, "\x1b[") {
				t.Fatalf("%s completion contains ANSI escapes", test.shell)
			}
		})
	}
}

func TestCompletionCommandRejectsUnsupportedShell(t *testing.T) {
	if _, err := execute(t, "completion", "elvish"); err == nil {
		t.Fatal("completion unexpectedly accepted an unsupported shell")
	}
}
