package command

import (
	"strings"
	"testing"

	"github.com/mitchellh/cli"
)

func TestJobHistoryCommand_Implements(t *testing.T) {
	t.Parallel()
	var _ cli.Command = &JobDispatchCommand{}
}

func TestJobHistoryCommand_Fails(t *testing.T) {
	t.Parallel()
	ui := new(cli.MockUi)
	cmd := &JobHistoryCommand{Meta: Meta{Ui: ui}}

	// Fails on misuse
	if code := cmd.Run([]string{"some", "bad", "args"}); code != 1 {
		t.Fatalf("expected exit code 1, got: %d", code)
	}
	if out := ui.ErrorWriter.String(); !strings.Contains(out, cmd.Help()) {
		t.Fatalf("expected help output, got: %s", out)
	}
	ui.ErrorWriter.Reset()

	if code := cmd.Run([]string{"-address=nope", "foo"}); code != 1 {
		t.Fatalf("expected exit code 1, got: %d", code)
	}
	if out := ui.ErrorWriter.String(); !strings.Contains(out, "Error listing jobs") {
		t.Fatalf("expected failed query error, got: %s", out)
	}
	ui.ErrorWriter.Reset()
}
