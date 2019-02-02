package commands

import (
	"strings"

	"github.com/mitchellh/cli"
)

type AirconCommand struct {
	UI cli.Ui
}

func (c *AirconCommand) Run(args []string) int {
	c.UI.Output(strings.TrimSpace(helpAircon))

	return 1
}

func (c *AirconCommand) Help() string {
	return strings.TrimSpace(helpList)
}

func (c *AirconCommand) Synopsis() string {
	return "Control Air Conditionar"
}

const helpAircon = `
Usage: remo aircon <subcommand>

SubCommands:
	list	List air conditionars
	send	Send signal to change the air conditionar setting
`
