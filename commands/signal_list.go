package commands

import (
	"fmt"

	"github.com/chroju/nature-remo-cli/configfile"
	"github.com/mitchellh/cli"
)

type SignalListCommand struct {
	UI cli.Ui
}

func (c *SignalListCommand) Run(args []string) int {
	if len(args) != 0 {
		c.UI.Warn(fmt.Sprintf("%s\n\ncommand \"signal list\" does not expect any args", helpSignalList))
		return 1
	}

	con, err := configfile.New()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	appliances, err := con.LoadAppliances()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	for _, a := range appliances {
		for _, s := range a.Signals {
			c.UI.Output(fmt.Sprintf("%s %s", a.Name, s.Name))
		}
	}

	return 0
}

func (c *SignalListCommand) Help() string {
	return helpSignalList
}

func (c *SignalListCommand) Synopsis() string {
	return "Show all appliance and signal names"
}

const helpSignalList = "Usage: remo signal list"
