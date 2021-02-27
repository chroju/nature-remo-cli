package commands

import (
	"os"
	"testing"

	"github.com/mitchellh/cli"
)

func TestRunSyncCommand(t *testing.T) {
	cases := []struct {
		args     []string
		preFunc  func()
		expected int
	}{
		{
			[]string{"dummy"},
			func() {},
			1,
		},
		{
			[]string{},
			func() { os.Setenv("HOME", "dummy") },
			2,
		},
	}
	ui := cli.NewMockUi()
	command := SyncCommand{UI: ui}

	for _, c := range cases {
		c.preFunc()
		if gotten := command.Run(c.args); gotten != c.expected {
			t.Errorf("want: %v\nget : %v", c.expected, gotten)
		}
	}
}
