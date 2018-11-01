package commands

import (
	"fmt"

	cloud "github.com/chroju/go-nature-remo/cloud"
	"github.com/chroju/nature-remo-cli/controller"
	"github.com/mitchellh/cli"
)

type SendCommand struct {
	UI cli.Ui
}

func (c *SendCommand) Run(args []string) int {
	if len(args) != 2 {
		c.UI.Error("Specify appliance and signal name")
		return 1
	}
	applianceName := args[0]
	signalName := args[1]

	con := controller.NewController()
	con.Read()

	var signalID string
	for _, v := range con.Setting.Appliances {
		if v.Name == applianceName {
			for _, signal := range v.Signals {
				if signal.Name == signalName {
					signalID = signal.ID
					break
				}
			}
			break
		}
	}
	if signalID == "" {
		c.UI.Error(fmt.Sprintf("Appliance '%s' - Signal '%s' not exist", applianceName, signalName))
		return 1
	}

	client := cloud.NewClient(con.Setting.Credential.Token)
	if _, err := client.SendSignal(signalID); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	return 0
}

func (c *SendCommand) Help() string {
	return "send signals"
}

func (c *SendCommand) Synopsis() string {
	return "remo send"
}
