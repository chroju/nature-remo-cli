package commands

import (
	"bytes"
	"context"
	"fmt"

	"github.com/chroju/nature-remo-cli/configfile"
	"github.com/mitchellh/cli"
	"github.com/olekukonko/tablewriter"
	"github.com/tenntenn/natureremo"
)

type AirconListCommand struct {
	UI cli.Ui
}

func (c *AirconListCommand) Run(args []string) int {
	if len(args) != 0 {
		c.UI.Warn(fmt.Sprintf("%s\n\ncommand \"aircon list\" does not expect any args", helpAirconList))
		return 1
	}

	path, err := configfile.GetConfigFilePath()
	if err != nil {
		c.UI.Error(err.Error())
		return 2
	}
	con, err := configfile.New(path)
	if err != nil {
		c.UI.Error(err.Error())
		return 2
	}

	token, err := con.LoadToken()
	if err != nil {
		c.UI.Error(err.Error())
		return 2
	}

	client := natureremo.NewClient(token)
	ctx := context.Background()

	appliances, err := client.ApplianceService.GetAll(ctx)
	if err != nil {
		c.UI.Error(err.Error())
		return 3
	}

	buf := &bytes.Buffer{}
	table := tablewriter.NewWriter(buf)
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
	c.UI.Output(buf.String())

	return 0
}

func (c *AirconListCommand) Help() string {
	return helpAirconList
}

func (c *AirconListCommand) Synopsis() string {
	return "Show current aircon list with their settings."
}

const helpAirconList = "Usage: remo aircon list"
