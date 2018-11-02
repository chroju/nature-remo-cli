package configfile

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path"

	cloud "github.com/chroju/go-nature-remo/cloud"
	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
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

type ConfigFile struct {
	path string
}

func New() (*ConfigFile, error) {
	user, err := user.Current()
	if err != nil {
		return nil, errors.Wrap(err, "Unexpected error")
	}
	configDir := user.HomeDir + "/.config"
	configFile := configDir + "/remo"

	return &ConfigFile{path: configFile}, nil
}

func (c *ConfigFile) SyncConfigFile(token string) error {
	if token == "" {
		file, err := c.readFile()
		if err != nil {
			return errors.Wrap(err, "Failed to read token from config file")
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

	client := cloud.NewClient(token)
	appliances, err := client.GetAppliances()
	if err != nil {
		return errors.Wrap(err, "Failed to get appliances from config file")
	}
	for _, a := range appliances {
		s.Appliances = append(s.Appliances, Appliance{Name: a.Nickname, ID: a.ID, Signals: a.Signals})
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
		return "", fmt.Errorf("You have not initialized")
	}

	return s.Credential.Token, nil
}

func (c *ConfigFile) LoadAppliances() ([]Appliance, error) {
	s, err := c.readFile()
	fmt.Println(s.Credential.Token)
	if err != nil {
		return nil, err
	}

	return s.Appliances, nil
}

func (c *ConfigFile) readFile() (*Setting, error) {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read config file")
	}
	s := Setting{}
	yaml.Unmarshal(data, s)
	fmt.Println(s.Credential.Token)

	return &s, nil
}
