package commands

import (
	"context"
	"fmt"

	"github.com/chroju/nature-remo-cli/configfile"
	"github.com/mitchellh/cli"
	"github.com/tenntenn/natureremo"
)

type AirconCommand struct {
	UI cli.Ui
}

func (c *AirconCommand) Run(args []string) int {
	if len(args) == 2 {
		c.UI.Warn(helpAircon)
		return 1
	}
	subcommand := args[0]

	con, err := configfile.New()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	switch subcommand {
	case "list":
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

		for _, v := range appliances {
			if v.Type == "AC" {
				fmt.Println(v.AirConSettings)
			}
		}
	}

	return 0
}

func (c *AirconCommand) Help() string {
	return helpAircon
}

func (c *AirconCommand) Synopsis() string {
	return "aircon"
}

const helpAircon = "Usage: remo aircon"
