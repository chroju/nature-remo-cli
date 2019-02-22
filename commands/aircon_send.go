package commands

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/chroju/nature-remo-cli/configfile"
	"github.com/mitchellh/cli"
	flag "github.com/spf13/pflag"
	"github.com/tenntenn/natureremo"
)

type AirconSendCommand struct {
	UI cli.Ui
}

func (c *AirconSendCommand) Run(args []string) int {
	if len(args) == 0 {
		c.UI.Warn(helpAirconSend)
		return 1
	}

	var aircon *(natureremo.Appliance)
	var settings *(natureremo.AirConSettings)
	var on, off bool
	var name, mode, volume, temperature string

	buf := &bytes.Buffer{}
	f := flag.NewFlagSet("aircon_list", flag.ContinueOnError)
	f.SetOutput(buf)
	f.BoolVar(&on, "on", false, "Power on the aircon")
	f.BoolVar(&off, "off", false, "Power off the aircon")
	f.StringVarP(&name, "name", "n", "", "Aircon name to operate")
	f.StringVarP(&mode, "mode", "m", "", "Aircon operation mode")
	f.StringVarP(&volume, "volume", "v", "", "Aircon wind volume")
	f.StringVarP(&temperature, "temperature", "t", "", "Aircon temperature")
	// f.StringVarP(&dir, "dir", "d", "", "Aircon wind direction")
	if err := f.Parse(args); err != nil {
		c.UI.Warn(helpAirconSend)
		return 1
	}

	con, err := configfile.New()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	token, err := con.LoadToken()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	client := natureremo.NewClient(token)
	ctx := context.Background()

	appliances, err := client.ApplianceService.GetAll(ctx)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	for _, v := range appliances {
		if v.Nickname == name {
			settings = v.AirConSettings
			aircon = v
			break
		}
	}

	if settings == nil {
		c.UI.Error(fmt.Sprintf("Not Found Aircon \"%s\"", name))
		return 1
	}

	newSettings := &natureremo.AirConSettings{
		Temperature:   settings.Temperature,
		OperationMode: settings.OperationMode,
		AirVolume:     settings.AirVolume,
		AirDirection:  settings.AirDirection,
		Button:        settings.Button,
	}
	var updateMessage []string

	if on && off {
		c.UI.Error("Cannnot use --on and --off at the same time.")
		return 1
	} else if on && settings.Button == natureremo.ButtonPowerOff {
		newSettings.Button = natureremo.ButtonPowerOn
		updateMessage = append(updateMessage, "OFF -> ON")
	} else if off && settings.Button == natureremo.ButtonPowerOn {
		newSettings.Button = natureremo.ButtonPowerOff
		updateMessage = append(updateMessage, "ON -> OFF")
	}

	switch mode {
	case "cool":
		newSettings.OperationMode = natureremo.OperationModeCool
	case "warm":
		newSettings.OperationMode = natureremo.OperationModeWarm
	case "dry":
		newSettings.OperationMode = natureremo.OperationModeDry
	case "blow":
		newSettings.OperationMode = natureremo.OperationModeBlow
	case "auto":
		newSettings.OperationMode = natureremo.OperationModeAuto
	}
	if newSettings.OperationMode != settings.OperationMode {
		updateMessage = append(updateMessage, fmt.Sprintf("%s -> %s", settings.OperationMode.StringValue(), mode))
	}

	switch volume {
	case "1":
		newSettings.AirVolume = natureremo.AirVolume1
	case "2":
		newSettings.AirVolume = natureremo.AirVolume2
	case "3":
		newSettings.AirVolume = natureremo.AirVolume3
	case "4":
		newSettings.AirVolume = natureremo.AirVolume4
	case "5":
		newSettings.AirVolume = natureremo.AirVolume5
	case "6":
		newSettings.AirVolume = natureremo.AirVolume6
	case "7":
		newSettings.AirVolume = natureremo.AirVolume7
	case "8":
		newSettings.AirVolume = natureremo.AirVolume8
	case "9":
		newSettings.AirVolume = natureremo.AirVolume9
	case "10":
		newSettings.AirVolume = natureremo.AirVolume10
	case "auto":
		newSettings.AirVolume = natureremo.AirVolumeAuto
	}
	if newSettings.AirVolume != settings.AirVolume {
		updateMessage = append(updateMessage, fmt.Sprintf("%s -> %s", settings.AirVolume.StringValue(), volume))
	}

	if len(temperature) > 0 && settings.Temperature != temperature {
		newSettings.Temperature = temperature
		updateMessage = append(updateMessage, fmt.Sprintf("%s -> %s", settings.Temperature, temperature))
	}

	if len(updateMessage) > 0 {
		if err := client.ApplianceService.UpdateAirConSettings(ctx, aircon, newSettings); err != nil {
			c.UI.Error(err.Error())
			return 1
		}
		c.UI.Output(fmt.Sprintf("Updated Aircon \"%s\" settings (%s)", name, strings.Join(updateMessage, ", ")))
	}

	return 0
}

func (c *AirconSendCommand) Help() string {
	return helpAirconSend
}

func (c *AirconSendCommand) Synopsis() string {
	return "Update the aircon settings."
}

const helpAirconSend = `Usage: remo aircon send -n name [OPTION]

  -n, --name string          Aircon name to operate

  --on                       Power on the aircon (exclusive with --off)
  --off                      Power off the aircon (exclusive with --on)
  -m, --mode string          Aircon operation mode
  -t, --temperature string   Aircon temperature
  -v, --volume string        Aircon wind volume
`
