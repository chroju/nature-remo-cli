package commands

import (
	"os"
	"testing"

	"github.com/mitchellh/cli"
)

func TestRunAirconSendCommand(t *testing.T) {
	cases := []struct {
		args     []string
		preFunc  func()
		expected int
	}{
		{
			[]string{},
			func() {},
			1,
		},
		{
			[]string{"--dummy"},
			func() {},
			1,
		},
		{
			[]string{"--on", "--off"},
			func() {},
			1,
		},
		{
			[]string{"-n", "dummy"},
			func() { os.Setenv("HOME", "dummy") },
			2,
		},
	}
	ui := cli.NewMockUi()
	command := AirconSendCommand{UI: ui}

	for _, c := range cases {
		c.preFunc()
		if gotten := command.Run(c.args); gotten != c.expected {
			t.Errorf("want: %v\nget : %v", c.expected, gotten)
		}
	}
}
