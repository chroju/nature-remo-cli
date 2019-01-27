package configfile

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
	"github.com/tenntenn/natureremo"
)

type Appliance struct {
	Name    string
	ID      string
	Type    natureremo.ApplianceType
	Signals []*natureremo.Signal
}

type Setting struct {
	Credential struct {
		Token string
	}
	Appliances []Appliance
}

type ConfigFile struct {
	path string
}

func New() (*ConfigFile, error) {
	user, err := user.Current()
	if err != nil {
		return nil, errors.Wrap(err, "Unexpected error")
	}

	path := user.HomeDir + "/.config/remo"
	return &ConfigFile{path: path}, nil
}

func (c *ConfigFile) SyncConfigFile(token string) error {
	if token == "" {
		file, err := c.readFile()
		if err != nil {
			return err
		}
		token = file.Credential.Token
	}

	s := Setting{}
	s.Credential.Token = token

	dirPath := path.Dir(c.path)
	if _, err := os.Stat(dirPath); err == os.ErrNotExist {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to make directory at %s", dirPath))
		}
	}

	file, err := os.Create(c.path)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to create config file %s", c.path))
	}
	defer file.Close()

	// client := cloud.NewClient(token)
	client := natureremo.NewClient(token)
	ctx := context.Background()
	appliances, err := client.ApplianceService.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("Failed to login")
	}

	for _, a := range appliances {
		s.Appliances = append(s.Appliances, Appliance{Name: a.Nickname, ID: a.ID, Type: a.Type, Signals: a.Signals})
	}

	y, err := yaml.Marshal(&s)
	io.WriteString(file, string(y))

	return nil
}

func (c *ConfigFile) LoadToken() (string, error) {
	s, err := c.readFile()
	if err != nil {
		return "", err
	}
	if s.Credential.Token == "" {
		return "", fmt.Errorf("You have not correctly initialized. Please execute \"remo init\"")
	}

	return s.Credential.Token, nil
}

func (c *ConfigFile) LoadAppliances() ([]Appliance, error) {
	s, err := c.readFile()
	if err != nil {
		return nil, err
	}

	return s.Appliances, nil
}

func (c *ConfigFile) readFile() (*Setting, error) {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		return nil, errors.Wrap(err, "You have not correctly initialized. Please execute \"remo init\"")
	}

	s := Setting{}
	err = yaml.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
