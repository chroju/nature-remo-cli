package commands

import (
	"strings"

	"github.com/mitchellh/cli"
)

type SignalCommand struct {
	UI cli.Ui
}

func (c *SignalCommand) Run(args []string) int {
	c.UI.Output(strings.TrimSpace(helpSignal))

	return 1
}

func (c *SignalCommand) Help() string {
	return strings.TrimSpace(helpSignal)
}

func (c *SignalCommand) Synopsis() string {
	return "Control Signals"
}

const helpSignal = `
Usage: remo signal <subcommand>

SubCommands:
    list    List signals
    send    Send the specified signal
`
