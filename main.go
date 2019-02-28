package main

import (
	"fmt"
	"os"

	"github.com/chroju/nature-remo-cli/commands"
	"github.com/mitchellh/cli"
)

const (
	app     = "remo"
	version = "0.1.0"
)

func main() {
	c := cli.NewCLI(app, version)
	c.Args = os.Args[1:]
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c.Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &commands.InitCommand{UI: ui}, nil
		},
		"aircon": func() (cli.Command, error) {
			return &commands.AirconCommand{UI: &cli.ColoredUi{Ui: ui, ErrorColor: cli.UiColorRed}}, nil
		},
		"aircon list": func() (cli.Command, error) {
			return &commands.AirconListCommand{UI: &cli.ColoredUi{Ui: ui, ErrorColor: cli.UiColorRed}}, nil
		},
		"aircon send": func() (cli.Command, error) {
			return &commands.AirconSendCommand{UI: &cli.ColoredUi{Ui: ui, ErrorColor: cli.UiColorRed}}, nil
		},
		"signal": func() (cli.Command, error) {
			return &commands.SignalCommand{UI: &cli.ColoredUi{Ui: ui, ErrorColor: cli.UiColorRed}}, nil
		},
		"signal list": func() (cli.Command, error) {
			return &commands.SignalListCommand{UI: &cli.ColoredUi{Ui: ui, ErrorColor: cli.UiColorRed}}, nil
		},
		"signal send": func() (cli.Command, error) {
			return &commands.SignalSendCommand{UI: &cli.ColoredUi{Ui: ui, ErrorColor: cli.UiColorRed}}, nil
		},
		"sync": func() (cli.Command, error) {
			return &commands.SyncCommand{UI: &cli.ColoredUi{Ui: ui, ErrorColor: cli.UiColorRed}}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		ui.Error(fmt.Sprintf("Error: %s", err))
	}

	os.Exit(exitStatus)
}
