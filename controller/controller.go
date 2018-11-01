package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"

	cloud "github.com/chroju/go-nature-remo/cloud"
	"github.com/go-yaml/yaml"
)

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

type Controller struct {
	path    string
	Setting Setting
}

func NewController() Controller {
	my, err := user.Current()
	if err != nil {
		panic(err)
	}
	configDir := my.HomeDir + "/.config"
	configFile := configDir + "/remo"

	if _, err := os.Stat(configDir); err == os.ErrNotExist {
		if err := os.Mkdir(configDir, 0755); err != nil {
			panic(err)
		}
	}

	return Controller{path: configFile}
}

func (c *Controller) SetToken(token string) {
	c.Setting.Credential.Token = token
}

func (c *Controller) Sync() error {
	if c.Setting.Credential.Token == "" {
		return fmt.Errorf("token is nothing")
	}

	file, err := os.Create(c.path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	client := cloud.NewClient(c.Setting.Credential.Token)
	appliances, err := client.GetAppliances()
	if err != nil {
		return err
	}
	for _, a := range appliances {
		c.Setting.Appliances = append(c.Setting.Appliances, Appliance{Name: a.Nickname, ID: a.ID, Signals: a.Signals})
	}

	y, err := yaml.Marshal(&c.Setting)
	io.WriteString(file, string(y))

	return nil
}

func (c *Controller) Read() error {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		return fmt.Errorf("Please execute init command at first")
	}
	yaml.Unmarshal(data, &c.Setting)

	return nil
}
