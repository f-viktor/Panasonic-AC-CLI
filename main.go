package main

import (
	"strconv"
	"strings"
)

func main() {
	args := ParseArgs()

	var params string

	// This clownery is necessary, because golang can't remove elements from an unmarshal
	// We only want to set the specific arguments we are changing
	if args.EcoFunctionData != -999 {
		params += `"ecoFunctionData":` + strconv.Itoa(args.EcoFunctionData) + `,`
	}
	if args.AirSwingLR != -999 {
		params += `"airSwingLR":` + strconv.Itoa(args.AirSwingLR) + `,`
	}
	if args.Nanoe != -999 {
		params += `"nanoe":` + strconv.Itoa(args.Nanoe) + `,`
	}
	if args.EcoNavi != -999 {
		params += `"ecoNavi":` + strconv.Itoa(args.EcoNavi) + `,`
	}
	if args.EcoMode != -999 {
		params += `"ecoMode":` + strconv.Itoa(args.EcoMode) + `,`
	}
	if args.OperationMode != -999 {
		params += `"operationMode":` + strconv.Itoa(args.OperationMode) + `,`
	}
	if args.FanAutoMode != -999 {
		params += `"fanAutoMode":` + strconv.Itoa(args.FanAutoMode) + `,`
	}
	if args.TemperatureSet != -999 {
		params += `"temperatureSet":` + strconv.FormatFloat(args.TemperatureSet, 'f', 1, 64) + `,`
	}
	if args.FanSpeed != -999 {
		params += `"fanSpeed":` + strconv.Itoa(args.FanSpeed) + `,`
	}
	if args.IAuto != -999 {
		params += `"iAuto":` + strconv.Itoa(args.IAuto) + `,`
	}
	if args.AirQuality != -999 {
		params += `"airQuality":` + strconv.Itoa(args.AirQuality) + `,`
	}
	if args.Operate != -999 {
		params += `"operate":` + strconv.Itoa(args.Operate) + `,`
	}
	if args.AirDirection != -999 {
		params += `"airDirection":` + strconv.Itoa(args.AirDirection) + `,`
	}
	if args.ActualNanoe != -999 {
		params += `"actualNanoe":` + strconv.Itoa(args.ActualNanoe) + `,`
	}
	if args.AirSwingUD != -999 {
		params += `"airSwingUD":` + strconv.Itoa(args.AirSwingUD) + `,`
	}

	// remove last comma
	params = strings.TrimRight(params, ",")

	SetDeviceStatus(params)
}
