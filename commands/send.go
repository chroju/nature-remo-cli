package commands

import (
	"context"
	"fmt"

	"github.com/fatih/color"

	"github.com/chroju/nature-remo-cli/configfile"
	"github.com/mitchellh/cli"
	"github.com/tenntenn/natureremo"
)

type SendCommand struct {
	UI cli.BasicUi
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

	var toSendSignal *natureremo.Signal
	for _, v := range appliances {
		if v.Name == applianceName {
			for _, signal := range v.Signals {
				if signal.Name == signalName {
					toSendSignal = signal
					break
				}
			}
			break
		}
	}
	if toSendSignal == nil {
		c.UI.Error(color.RedString(fmt.Sprintf("Appliance '%s' - Signal '%s' is invalid", applianceName, signalName)))
		return 1
	}

	token, err := con.LoadToken()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	client := natureremo.NewClient(token)
	ctx := context.Background()
	if err := client.SignalService.Send(ctx, toSendSignal); err != nil {
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
