package commands

import (
	"bufio"
	"io"
	"os"
	"os/user"

	"github.com/mitchellh/cli"
)

type InitCommand struct {
	UI cli.Ui
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

		io.WriteString(file, scanner.Text())
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
