package main

import (
	"fmt"
	"os"

	"github.com/chroju/nature-remo-cli/commands"
	"github.com/mitchellh/cli"
)

const (
	app     = "remo"
	version = "0.0.1"
)

var UI cli.Ui

func main() {
	c := cli.NewCLI(app, version)
	c.Args = os.Args[1:]
	UI = &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c.Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &commands.InitCommand{UI: UI}, nil
		},
		"aircon": func() (cli.Command, error) {
			return &commands.AirconCommand{UI: UI}, nil
		},
		"list": func() (cli.Command, error) {
			return &commands.ListCommand{UI: UI}, nil
		},
		"send": func() (cli.Command, error) {
			return &commands.SendCommand{UI: UI}, nil
		},
		"sync": func() (cli.Command, error) {
			return &commands.SyncCommand{UI: UI}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		UI.Error(fmt.Sprintf("Error: %s", err))
	}

	os.Exit(exitStatus)
}
