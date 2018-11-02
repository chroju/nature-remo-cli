package commands

import (
	"fmt"

	"github.com/chroju/nature-remo-cli/configfile"
	"github.com/mitchellh/cli"
)

type ListCommand struct {
	UI cli.Ui
}

func (c *ListCommand) Run(args []string) int {
	if len(args) != 0 {
		c.UI.Error("command \"list\" does not expect any args")
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

func (c *ListCommand) Help() string {
	return "list up appliances or signals"
}

func (c *ListCommand) Synopsis() string {
	return "remo list"
}
