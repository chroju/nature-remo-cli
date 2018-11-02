package commands

import (
	"bufio"
	"os"

	"github.com/chroju/nature-remo-cli/configfile"
	"github.com/mitchellh/cli"
)

type InitCommand struct {
	UI cli.Ui
}

func (c *InitCommand) Run(args []string) int {
	if len(args) > 0 {
		c.UI.Error("command \"init\" does not expect any args")
		return 1
	}

	c.UI.Output("Nature Remo OAuth Token:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	token := scanner.Text()

	con, err := configfile.New()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	if _, err := con.LoadToken(); err == nil {
		c.UI.Output("You have already initialized remo. Override current settings ? [y/n]")
		for scanner.Scan() {
			if scanner.Text() == "y" {
				break
			} else if scanner.Text() == "n" {
				return 2
			} else {
				c.UI.Output("[y/n]?")
			}
		}
	}

	c.UI.Output("Initializing ...")
	if err := con.SyncConfigFile(token); err != nil {
		c.UI.Error("Failed to initialize !")
		return 1
	}
	c.UI.Output("Successfully initialized.")

	return 0
}

func (c *InitCommand) Help() string {
	return "initialize command"
}

func (c *InitCommand) Synopsis() string {
	return "remo init"
}
