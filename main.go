package main

import (
	"strconv"
	"strings"
)

func main() {
	args := parseArgs()

	switch args.Command {
	case "set":
		setParamsOnDevice(args.Params, args.DeviceId)

	case "list":
		GetDeviceList().Print()

	case "status":
		GetDeviceStatus().Print()
	}
}

func setParamsOnDevice(argParams DeviceParamList, deviceId string) {
	var sendParams string

	// This clownery is necessary, because golang can't remove elements from an unmarshal
	// We only want to set the specific arguments we are changing
	if argParams.EcoFunctionData != -999 {
		sendParams += `"ecoFunctionData":` + strconv.Itoa(argParams.EcoFunctionData) + `,`
	}
	if argParams.AirSwingLR != -999 {
		sendParams += `"airSwingLR":` + strconv.Itoa(argParams.AirSwingLR) + `,`
	}
	if argParams.Nanoe != -999 {
		sendParams += `"nanoe":` + strconv.Itoa(argParams.Nanoe) + `,`
	}
	if argParams.EcoNavi != -999 {
		sendParams += `"ecoNavi":` + strconv.Itoa(argParams.EcoNavi) + `,`
	}
	if argParams.EcoMode != -999 {
		sendParams += `"ecoMode":` + strconv.Itoa(argParams.EcoMode) + `,`
	}
	if argParams.OperationMode != -999 {
		sendParams += `"operationMode":` + strconv.Itoa(argParams.OperationMode) + `,`
	}
	if argParams.FanAutoMode != -999 {
		sendParams += `"fanAutoMode":` + strconv.Itoa(argParams.FanAutoMode) + `,`
	}
	if argParams.TemperatureSet != -999 {
		sendParams += `"temperatureSet":` + strconv.FormatFloat(argParams.TemperatureSet, 'f', 1, 64) + `,`
	}
	if argParams.FanSpeed != -999 {
		sendParams += `"fanSpeed":` + strconv.Itoa(argParams.FanSpeed) + `,`
	}
	if argParams.IAuto != -999 {
		sendParams += `"iAuto":` + strconv.Itoa(argParams.IAuto) + `,`
	}
	if argParams.AirQuality != -999 {
		sendParams += `"airQuality":` + strconv.Itoa(argParams.AirQuality) + `,`
	}
	if argParams.Operate != -999 {
		sendParams += `"operate":` + strconv.Itoa(argParams.Operate) + `,`
	}
	if argParams.AirDirection != -999 {
		sendParams += `"airDirection":` + strconv.Itoa(argParams.AirDirection) + `,`
	}
	if argParams.ActualNanoe != -999 {
		sendParams += `"actualNanoe":` + strconv.Itoa(argParams.ActualNanoe) + `,`
	}
	if argParams.AirSwingUD != -999 {
		sendParams += `"airSwingUD":` + strconv.Itoa(argParams.AirSwingUD) + `,`
	}

	// remove last comma
	sendParams = strings.TrimRight(sendParams, ",")

	SetDeviceStatus(sendParams)
}
