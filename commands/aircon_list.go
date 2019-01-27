package commands

import (
	"context"

	"github.com/chroju/nature-remo-cli/configfile"
	"github.com/mitchellh/cli"
	"github.com/olekukonko/tablewriter"
	"github.com/tenntenn/natureremo"
)

type AirconListCommand struct {
	UI cli.BasicUi
}

func (c *AirconListCommand) Run(args []string) int {
	if len(args) != 0 {
		c.UI.Warn(helpAirconList)
		return 1
	}

	con, err := configfile.New()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	token, err := con.LoadToken()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	client := natureremo.NewClient(token)
	ctx := context.Background()

	appliances, err := client.ApplianceService.GetAll(ctx)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	table := tablewriter.NewWriter(c.UI.Writer)
	table.SetHeader([]string{"NAME", "POWER", "TEMP", "MODE", "VOLUME", "DIRECTION"})
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetHeaderLine(false)
	table.SetColumnSeparator("")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, v := range appliances {
		if v.Type == "AC" {
			a := v.AirConSettings
			var button string
			if a.Button == natureremo.ButtonPowerOff {
				button = "OFF"
			} else {
				button = "ON"
			}
			table.Append([]string{v.Nickname,
				button,
				a.Temperature,
				a.OperationMode.StringValue(),
				a.AirVolume.StringValue(),
				a.AirDirection.StringValue()})
		}
	}
	table.Render()

	return 0
}

func (c *AirconListCommand) Help() string {
	return helpAirconList
}

func (c *AirconListCommand) Synopsis() string {
	return "Show current aircon list with their settings."
}

const helpAirconList = "Usage: remo aircon list"
