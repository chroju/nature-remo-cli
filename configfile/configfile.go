package configfile

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path"

	"github.com/pkg/errors"
	"github.com/tenntenn/natureremo"
	yaml "gopkg.in/yaml.v2"
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
	Appliances []*Appliance
}

type Config interface {
	LoadAllSetting() (*Setting, error)
	LoadToken() (string, error)
	LoadAppliances() ([]*Appliance, error)
	Sync(string) error
}

type configFile struct {
	path    string
	setting *Setting
}

func New() (Config, error) {
	user, err := user.Current()
	if err != nil {
		return nil, errors.Wrap(err, "Unexpected error")
	}

	path := user.HomeDir + "/.config/remo"

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "You have not correctly initialized. Please execute \"remo init\"")
	}

	setting := &Setting{}
	err = yaml.Unmarshal(data, setting)
	if err != nil {
		return nil, err
	}

	return &configFile{
		path:    path,
		setting: setting,
	}, nil
}

func (c *configFile) Sync(token string) error {
	if token == "" {
		var err error
		token, err = c.LoadToken()
		if err != nil {
			return err
		}
	}

	s := Setting{}
	s.Credential.Token = token

	dirPath := path.Dir(c.path)
	if _, err := os.Stat(dirPath); err == os.ErrNotExist {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to make directory at %s.", dirPath))
		}
	}

	file, err := os.Create(c.path)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to create config file %s.", c.path))
	}
	defer file.Close()

	client := natureremo.NewClient(token)
	ctx := context.Background()
	appliances, err := client.ApplianceService.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("Failed to login.")
	}

	for _, a := range appliances {
		s.Appliances = append(s.Appliances, &Appliance{Name: a.Nickname, ID: a.ID, Type: a.Type, Signals: a.Signals})
	}

	y, err := yaml.Marshal(&s)
	io.WriteString(file, string(y))

	return nil
}

func (c *configFile) LoadToken() (string, error) {
	return c.setting.Credential.Token, nil
}

func (c *configFile) LoadAppliances() ([]*Appliance, error) {
	return c.setting.Appliances, nil
}

func (c *configFile) LoadAllSetting() (*Setting, error) {
	return c.setting, nil
}
