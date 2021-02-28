package commands

import (
	"os"
	"testing"

	"github.com/mitchellh/cli"
)

func TestRunSignalSendCommand(t *testing.T) {
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
			[]string{"dummy", "dummy", "dummy"},
			func() {},
			1,
		},
		{
			[]string{"light", "on"},
			func() { os.Setenv("HOME", "dummy") },
			2,
		},
	}
	ui := cli.NewMockUi()
	command := SignalSendCommand{UI: ui}

	for _, c := range cases {
		c.preFunc()
		if gotten := command.Run(c.args); gotten != c.expected {
			t.Errorf("want: %v\nget : %v", c.expected, gotten)
		}
	}
}
