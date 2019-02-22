package commands

import (
	"context"
	"fmt"

	"github.com/chroju/nature-remo-cli/configfile"
	"github.com/mitchellh/cli"
	"github.com/tenntenn/natureremo"
)

type SignalSendCommand struct {
	UI cli.Ui
}

func (c *SignalSendCommand) Run(args []string) int {
	if len(args) != 2 {
		c.UI.Warn(fmt.Sprintf("%s\n\nSpecify appliance and signal name", helpSignalSend))
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
		c.UI.Error(fmt.Sprintf("Failed. Appliance '%s' - Signal '%s' is not found.", applianceName, signalName))
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

	c.UI.Output("Succeeded.")
	return 0
}

func (c *SignalSendCommand) Help() string {
	return helpSignalSend
}

func (c *SignalSendCommand) Synopsis() string {
	return "Send signal"
}

const helpSignalSend = "Usage: remo signal send <appliance> <signal>"
