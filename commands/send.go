package commands

import (
	"fmt"

	"github.com/fatih/color"

	cloud "github.com/chroju/go-nature-remo/cloud"
	"github.com/chroju/nature-remo-cli/configfile"
	"github.com/mitchellh/cli"
)

type SendCommand struct {
	UI cli.Ui
}

func (c *SendCommand) Run(args []string) int {
	if len(args) != 2 {
		c.UI.Warn(fmt.Sprintf("%s\nSpecify appliance and signal name", helpSend))
		return 1
	}
	applianceName := args[0]
	signalName := args[1]

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

	var signalID string
	for _, v := range appliances {
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
		c.UI.Error(color.RedString(fmt.Sprintf("Appliance '%s' - Signal '%s' is invalid", applianceName, signalName)))
		return 1
	}

	token, err := con.LoadToken()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	client := cloud.NewClient(token)
	if _, err := client.SendSignal(signalID); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output("Success.")
	return 0
}

func (c *SendCommand) Help() string {
	return helpSend
}

func (c *SendCommand) Synopsis() string {
	return "Send signal"
}

const helpSend = "Usage: remo send <appliance> <signal>"
