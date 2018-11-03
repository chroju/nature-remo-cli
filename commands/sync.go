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
		c.UI.Warn(fmt.Sprintf("%s\ncommand \"sync\" does not expect any args", helpSync))
		return 1
	}

	con, err := configfile.New()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	if err := con.SyncConfigFile(""); err != nil {
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
