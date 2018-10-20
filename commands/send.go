package commands

import (
	"io/ioutil"
	"os/user"
	"strings"

	cloud "github.com/chroju/go-nature-remo/cloud"
	"github.com/mitchellh/cli"
)

type SendCommand struct {
	UI cli.Ui
}

func (c *SendCommand) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error("Specify signal ID")
		return 1
	}
	signalID := args[0]

	my, err := user.Current()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(my.HomeDir + "/.config/remo")
	if err != nil {
		c.UI.Error("Please execute init command at first")
		return 1
	}
	token := strings.TrimRight(string(data), "\r\n")

	client := cloud.NewClient(token)
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
