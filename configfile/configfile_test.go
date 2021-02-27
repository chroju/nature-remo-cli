package configfile

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/tenntenn/natureremo"
)

var (
	testFilePath       = "./test.yaml"
	expectedToken      = "_TESTTOKEN"
	expectedAppliances = []*Appliance{
		{
			Name:    "TV",
			ID:      "TV-id",
			Type:    "TV",
			Signals: []*natureremo.Signal{},
		},
		{
			Name: "light",
			ID:   "light-id",
			Signals: []*natureremo.Signal{
				{
					ID:    "signal-1",
					Name:  "brighten",
					Image: "ico_foo",
				},
				{
					ID:    "signal-2",
					Name:  "darken",
					Image: "ico_bar",
				},
			},
		},
	}
	expectedSetting = &Setting{
		Appliances: expectedAppliances,
	}
)

func TestNew(t *testing.T) {
	_, err := New(testFilePath)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestLoadToken(t *testing.T) {
	config, err := New(testFilePath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	gottenToken, err := config.LoadToken()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if gottenToken != expectedToken {
		t.Errorf("want: %s\nget : %s", expectedToken, gottenToken)
	}
}

func TestLoadAppliances(t *testing.T) {
	config, err := New(testFilePath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	gottenAppliances, err := config.LoadAppliances()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !reflect.DeepEqual(gottenAppliances, expectedAppliances) {
		t.Errorf("want: %v\nget : %v", expectedAppliances, gottenAppliances)
	}
}

func TestLoadAllSetting(t *testing.T) {
	expectedSetting.Credential.Token = expectedToken

	config, err := New(testFilePath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	gottenSetting, err := config.LoadAllSetting()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !reflect.DeepEqual(gottenSetting, expectedSetting) {
		t.Errorf("want: %v\nget : %v", expectedSetting, gottenSetting)
	}
}

func TestGetConfigFilePath(t *testing.T) {
	expected := filepath.Join(os.Getenv("HOME"), ".config", "remo")

	gotten, err := GetConfigFilePath()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if gotten != expected {
		t.Errorf("want: %s\nget : %s", expected, gotten)
	}
}
