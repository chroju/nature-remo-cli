package commands

import (
	"bufio"
	"io"
	"os"
	"os/user"

	cloud "github.com/chroju/go-nature-remo/cloud"
	"github.com/go-yaml/yaml"
	"github.com/mitchellh/cli"
)

type InitCommand struct {
	UI cli.Ui
}

type Appliance struct {
	Name    string
	ID      string
	Signals []cloud.Signal
}

type Setting struct {
	Credential struct {
		Token string
	}
	Appliances []Appliance
}

func (c *InitCommand) Run(args []string) int {
	if len(args) > 0 {
		c.UI.Error("command \"init\" does not expect any args")
		return 1
	}

	my, err := user.Current()
	if err != nil {
		panic(err)
	}
	configDir := my.HomeDir + "/.config"
	configFile := configDir + "/remo"

	c.UI.Output("Nature Remo OAuth Token:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		token := scanner.Text()

		if _, err := os.Stat(configDir); err == os.ErrNotExist {
			if err := os.Mkdir(configDir, 0755); err != nil {
				panic(err)
			}
		}
		file, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		c.UI.Output("Initializing ...")

		client := cloud.NewClient(token)
		appliances, err := client.GetAppliances()
		if err != nil {
			c.UI.Error(err.Error())
			return 1
		}
		s := Setting{}
		s.Credential.Token = token
		for _, a := range appliances {
			s.Appliances = append(s.Appliances, Appliance{Name: a.Nickname, ID: a.ID, Signals: a.Signals})
		}

		y, err := yaml.Marshal(&s)
		io.WriteString(file, string(y))

		c.UI.Output("Successfully initialized.")
		break
	}

	return 0
}

func (c *InitCommand) Help() string {
	return "initialize command"
}

func (c *InitCommand) Synopsis() string {
	return "remo init"
}
