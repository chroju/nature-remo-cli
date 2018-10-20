package commands

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"strings"

	cloud "github.com/chroju/go-nature-remo/cloud"
	"github.com/mitchellh/cli"
)

type ListCommand struct {
	UI cli.Ui
}

func (c *ListCommand) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error("Select appliances or signals")
		return 1
	}
	target := args[0]

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

	switch target {
	case "appliances":
		appliances, err := client.GetAppliances()
		if err != nil {
			c.UI.Error(err.Error())
			return 1
		}
		for _, a := range appliances {
			c.UI.Output(fmt.Sprintf("%s\t%s", a.Nickname, a.ID))
		}
	case "signals":
		appliances, err := client.GetAppliances()
		if err != nil {
			c.UI.Error(err.Error())
			return 1
		}
		for _, a := range appliances {
			c.UI.Output(fmt.Sprintf("%s\t%s", a.Nickname, a.ID))
			for _, s := range a.Signals {
				c.UI.Output(fmt.Sprintf("\t%s\t%s", s.Name, s.ID))
			}
		}
	}

	return 0
}

func (c *ListCommand) Help() string {
	return "list up appliances or signals"
}

func (c *ListCommand) Synopsis() string {
	return "remo list"
}
