package commands

import (
	"fmt"

	"github.com/chroju/nature-remo-cli/configfile"
	"github.com/mitchellh/cli"
)

type InitCommand struct {
	UI cli.Ui
}

func (c *InitCommand) Run(args []string) int {
	if len(args) > 0 {
		c.UI.Error(fmt.Sprintf("%s\ncommand \"init\" does not expect any args", helpInit))
		return 1
	}

	con, err := configfile.New()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	if _, err := con.LoadToken(); err == nil {
		reply, err := c.UI.Ask("You have already initialized remo. Override current settings ? [y/n]")
		if err != nil {
			c.UI.Error(err.Error())
			return 1
		}
		for {
			if reply == "y" {
				break
			} else if reply == "n" {
				return 2
			} else {
				reply, err = c.UI.Ask("[y/n]?")
				if err != nil {
					c.UI.Error(err.Error())
					return 1
				}
			}
		}
	}

	token, err := c.UI.AskSecret("Nature Remo OAuth Token:")
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output("Initializing ...")
	if err := con.SyncConfigFile(token); err != nil {
		c.UI.Error("Failed to initialize!")
		return 1
	}
	c.UI.Output("Successfully initialized.")

	return 0
}

func (c *InitCommand) Help() string {
	return helpInit
}

func (c *InitCommand) Synopsis() string {
	return "Initialize remo with your OAuth token"
}

const helpInit = "Usage: remo init"
