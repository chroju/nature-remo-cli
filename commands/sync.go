package commands

import (
	"fmt"

	"github.com/chroju/nature-remo-cli/configfile"
	"github.com/mitchellh/cli"
)

type SyncCommand struct {
	UI cli.Ui
}

func (c *SyncCommand) Run(args []string) int {
	if len(args) != 0 {
		c.UI.Warn(fmt.Sprintf("%s\n\ncommand \"sync\" does not expect any args", helpSync))
		return 1
	}

	path, err := configfile.GetConfigFilePath()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	con, err := configfile.New(path)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	if _, err := con.LoadToken(); err == nil {
		reply, err := c.UI.Ask("Override current settings ? [y/n]")
		if err != nil {
			c.UI.Error(err.Error())
			return 1
		}
		for {
			if reply == "y" {
				break
			} else if reply == "n" {
				return 0
			} else {
				reply, err = c.UI.Ask("[y/n]?")
				if err != nil {
					c.UI.Error(err.Error())
					return 1
				}
			}
		}
	}

	if err := con.Sync(""); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output("Synced!")
	return 0
}

func (c *SyncCommand) Help() string {
	return helpSync
}

func (c *SyncCommand) Synopsis() string {
	return "Sync local config with your latest one"
}

const helpSync = "Usage: remo sync"
