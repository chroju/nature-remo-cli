package commands

import (
	"bufio"
	"os"

	cloud "github.com/chroju/go-nature-remo/cloud"
	"github.com/chroju/nature-remo-cli/controller"
	"github.com/mitchellh/cli"
)

type InitCommand struct {
	UI cli.Ui
}

type Appliance struct {
	Name    string
	ID      string
	Signals []cloud.Signal
}

type Setting struct {
	Credential struct {
		Token string
	}
	Appliances []Appliance
}

func (c *InitCommand) Run(args []string) int {
	if len(args) > 0 {
		c.UI.Error("command \"init\" does not expect any args")
		return 1
	}

	c.UI.Output("Nature Remo OAuth Token:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		token := scanner.Text()

		c.UI.Output("Initializing ...")
		con := controller.NewController()
		con.SetToken(token)
		if err := con.Sync(); err != nil {
			c.UI.Error("Failed to initialize !")
			return 1
		}
		c.UI.Output("Successfully initialized.")
		break
	}

	return 0
}

func (c *InitCommand) Help() string {
	return "initialize command"
}

func (c *InitCommand) Synopsis() string {
	return "remo init"
}
