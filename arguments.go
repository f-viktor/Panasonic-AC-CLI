package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Arguments struct {
	Params   DeviceParamList
	Command  string
	DeviceId string
}

// parse command line arguments
func parseArgs() Arguments {

	config := flag.String("config", "config.json", "Path to config file")
	quiet := flag.Bool("q", false, "Don't print status update messages")
	help := flag.Bool("h", false, "Display this help text")
	helpDetails := flag.Bool("help", false, "Display detailed information")
	listDevices := flag.Bool("list", false, "List devices associated with user")
	deviceId := flag.String("dev", "", "DeviceGUID use -list to print avaliable")
	statDevice := flag.Bool("status", false, "Display current status of device")

	power := flag.String("power", "", "Power > On/Off ")
	mode := flag.String("m", "", "AC mode > auto/heat/cool/dry/nanoe")
	eco := flag.String("eco", "", "Economy setting > none/quiet/strong")
	fanSpeed := flag.String("fan", "", "Fan speed > auto/1/2/3/4/5")
	angleVertical := flag.String("av", "", "Vertical angle of the air > auto/top/centertop/center/centerbottom/bottom")
	angleHorizontal := flag.String("ah", "", "Horizontal angle of the air > auto/left/centerleft/center/centerright/right")
	nanoe := flag.String("nanoe", "", "Enabling nanoe setting but not nanoe mode > On/Off")
	temperatureSet := flag.String("t", "", "TemperatureSet in Celsius, supports 1 decimal > 22.5 ")

	flag.Parse()
	readConfig(*config)
	GlobalConfig.Verbose = !*quiet
	if *deviceId != "" {
		GlobalConfig.DeviceGuid = *deviceId
	}
	overwriteConfigFile(GlobalConfig)

	if *help {
		flag.Usage()
		os.Exit(2)
	}

	if *helpDetails {
		printHelpDetails()
		os.Exit(2)
	}
	args := Arguments{}
	args.Command = "set"
	args.Params = setDeviceFlags(strings.ToLower(*power), strings.ToLower(*mode), strings.ToLower(*eco), strings.ToLower(*fanSpeed), strings.ToLower(*angleVertical), strings.ToLower(*angleHorizontal), strings.ToLower(*nanoe), strings.ToLower(*temperatureSet))

	if *listDevices {
		args.Command = "list"
	}
	if *statDevice {
		args.Command = "status"
	}

	return args
}

func setDeviceFlags(power string, mode string, eco string, fanSpeed string, angleVertical string, angleHorizontal string, nanoe string, temperatureSet string) DeviceParamList {
	//-999 is used to mark entries that are not set
	params := DeviceParamList{-999, -999, -999, -999, -999, -999, -999, -999, -999, -999, -999, -999, -999, -999, -999, -999, -999}

	switch power {
	case "on":
		params.Operate = 1
	case "off":
		params.Operate = 0
	}

	switch mode {
	case "auto":
		params.OperationMode = 0
	case "heat":
		params.OperationMode = 3
	case "cool":
		params.OperationMode = 2
	case "dry":
		params.OperationMode = 1
	case "nanoe":
		params.OperationMode = 4
		params.Nanoe = 4
	}

	switch eco {
	case "none":
		params.EcoNavi = 0
		params.EcoMode = 0
	case "quiet":
		params.EcoNavi = 1
		params.EcoMode = 2
		params.IAuto = 1
	case "strong":
		params.EcoNavi = 1
		params.EcoMode = 1
		params.IAuto = 1
	}

	switch fanSpeed {
	case "auto":
		params.FanSpeed = 0
	case "1":
		params.FanSpeed = 1
	case "2":
		params.FanSpeed = 2
	case "3":
		params.FanSpeed = 3
	case "4":
		params.FanSpeed = 4
	case "5":
		params.FanSpeed = 5
	}

	switch angleVertical {
	case "top":
		params.AirSwingUD = 0
	case "bottom":
		params.AirSwingUD = 1
	case "center":
		params.AirSwingUD = 2
	case "centertop":
		params.AirSwingUD = 3
	case "centerbottom":
		params.AirSwingUD = 4
	}

	switch angleHorizontal {
	case "left":
		params.AirSwingLR = 0
	case "centerleft":
		params.AirSwingLR = 4
	case "center":
		params.AirSwingLR = 2
	case "centerright":
		params.AirSwingLR = 5
	case "right":
		params.AirSwingLR = 1
	}

	switch nanoe {
	case "on":
		params.Nanoe = 4
	case "off":
		params.Nanoe = 1
	}

	if temperatureSet != "" {
		params.TemperatureSet, _ = strconv.ParseFloat(temperatureSet, 64)
	}

	// if either is explicitly set, lets assume they are not set to auto
	if angleHorizontal != "" && angleVertical != "" {
		params.FanAutoMode = 1
	}
	if angleHorizontal == "auto" {
		params.FanAutoMode = 3
	}
	if angleVertical == "auto" {
		params.FanAutoMode = 2
	}
	if angleVertical == "auto" && angleHorizontal == "auto" {
		params.FanAutoMode = 0
	}

	return params
}

func printHelpDetails() {
	fmt.Println(`
++++++++++++++++++++++++++++++++++++
|App  buttons and how to mimic them|
++++++++++++++++++++++++++++++++++++

Turn off: {"operate":0,"operationMode":1}
Turn on: {"operate":1,"operationMode":1}

Heat Mode: {"fanSpeed":0,"operationMode":3,"temperatureSet":23.0}
Auto mode: {"fanSpeed":5,"operationMode":0,"temperatureSet":22.0}
Cool mode: {"fanSpeed":0,"operationMode":2,"temperatureSet":27.0}
Dry  Mode: {"fanSpeed":0,"operationMode":1,"temperatureSet":25.0}
(dryy mode is incompatible with setting temperature and eco modes)
"nanoe Mode" ... whatever that is: {"ecoMode":0,"ecoNavi":1,"fanSpeed":0,"nanoe":4,"operationMode":4,"temperatureSet":27.0}

Setting temperature: {"operationMode":0,"temperatureSet":24.0}

Quiet button On: {"ecoMode":2,"ecoNavi":1,"iAuto":1,"operationMode":0}
Quiet button Off: {"ecoMode":0}

Strong button On: {"ecoMode":1,"ecoNavi":1,"iAuto":1,"operationMode":0}
Strong button Off: {"ecoMode":0,"operationMode":0}

Fan speed config: {"fanSpeed":0,"operationMode":0}
"fanSpeed":0 = auto
"fanSpeed":1 = lowest
"fanSpeed":2 = medium-low
"fanSpeed":3 = medium
"fanSpeed":4 = medium-high
"fanSpeed":5 = highest

Air Up/Down angle: {"airSwingUD":0,"fanAutoMode":3,"operationMode":0}
"airSwingUD":0 = horizontal / highest
"airSwingUD":1 = vertical / straigth down
"airSwingUD":2 = middle / 45 angle
"airSwingUD":3 = middle-high
"airSwingUD":4 = middle-low

Air Left/Right angle: {"airSwingLR":2,"ecoNavi":1,"fanAutoMode":2,"operationMode":0}
"airSwingLR":0 = left
"airSwingLR":4 = center-left
"airSwingLR":2 = center
"airSwingLR":5 = center-right
"airSwingLR":1 = right

All automatic off: {"airSwingUD":2,"fanAutoMode":1,"operationMode":3}
Only Left/Right auto:  {"fanAutoMode":3,"operationMode":3}
Only Up/Down auto:  {"airSwingLR":2,"ecoNavi":1,"fanAutoMode":2,"operationMode":3}
Left/Right & Up/Down auto: {"fanAutoMode":0,"operationMode":3}

setting nanoe works differently, you first have to enable it:
you dont set it as device parameter but inested as a separate param: {"deviceGuid":"CS-Z25VKEW+4962315197","nanoe":true}
Then you can tourn it on as a param: "parameters":{"nanoe":4,"operationMode":3}
and off : {"nanoe":1,"operationMode":3}

there are also other optons in the menu but they dont seem to actually do anything or even save.

The rest of the options are not configurable on my specific device`)

}
