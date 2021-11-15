package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

// used as a response when querying the device status at https://accsmart.panasonic.com/device/now/<deviceID>
type DeviceStatusFull struct {
	DryTempMin  int `json:"dryTempMin"`
	ModeAvlList struct {
		AutoMode int `json:"autoMode"`
		FanMode  int `json:"fanMode"`
	} `json:"modeAvlList"`
	AirSwingLR       bool            `json:"airSwingLR"`
	Nanoe            bool            `json:"nanoe"`
	AutoMode         bool            `json:"autoMode"`
	AutoSwingUD      bool            `json:"autoSwingUD"`
	EcoNavi          bool            `json:"ecoNavi"`
	HeatTempMax      int             `json:"heatTempMax"`
	TemperatureUnit  int             `json:"temperatureUnit"`
	IAutoX           bool            `json:"iAutoX"`
	CoolTempMin      int             `json:"coolTempMin"`
	AutoTempMin      int             `json:"autoTempMin"`
	QuietMode        bool            `json:"quietMode"`
	PowerfulMode     bool            `json:"powerfulMode"`
	Timestamp        int64           `json:"timestamp"`
	FanMode          bool            `json:"fanMode"`
	CoolMode         bool            `json:"coolMode"`
	SummerHouse      int             `json:"summerHouse"`
	CoolTempMax      int             `json:"coolTempMax"`
	Permission       int             `json:"permission"`
	DryMode          bool            `json:"dryMode"`
	HeatMode         bool            `json:"heatMode"`
	FanSpeedMode     int             `json:"fanSpeedMode"`
	DryTempMax       int             `json:"dryTempMax"`
	AutoTempMax      int             `json:"autoTempMax"`
	FanDirectionMode int             `json:"fanDirectionMode"`
	EcoFunction      int             `json:"ecoFunction"`
	HeatTempMin      int             `json:"heatTempMin"`
	PairedFlg        bool            `json:"pairedFlg"`
	Parameters       DeviceParamList `json:"parameters"`
}

type DeviceParamList struct {
	EcoFunctionData   int     `json:"ecoFunctionData"`
	AirSwingLR        int     `json:"airSwingLR"`
	Nanoe             int     `json:"nanoe"`
	EcoNavi           int     `json:"ecoNavi"`
	EcoMode           int     `json:"ecoMode"`
	OperationMode     int     `json:"operationMode"`
	FanAutoMode       int     `json:"fanAutoMode"`
	TemperatureSet    float64 `json:"temperatureSet"`
	FanSpeed          int     `json:"fanSpeed"`
	IAuto             int     `json:"iAuto"`
	AirQuality        int     `json:"airQuality"`
	InsideTemperature int     `json:"insideTemperature"`
	OutTemperature    int     `json:"outTemperature"`
	Operate           int     `json:"operate"`
	AirDirection      int     `json:"airDirection"`
	ActualNanoe       int     `json:"actualNanoe"`
	AirSwingUD        int     `json:"airSwingUD"`
}

// used as a response when listing devices at https://accsmart.panasonic.com/device/group
type DeviceGroupList struct {
	IaqStatus struct {
		StatusCode int `json:"statusCode"`
	} `json:"iaqStatus"`
	UIFlg      bool `json:"uiFlg"`
	GroupCount int  `json:"groupCount"`
	GroupList  []struct {
		GroupID    int    `json:"groupId"`
		GroupName  string `json:"groupName"`
		DeviceList []struct {
			DeviceGUID         string `json:"deviceGuid"`
			DeviceType         string `json:"deviceType"`
			DeviceName         string `json:"deviceName"`
			Permission         int    `json:"permission"`
			DeviceModuleNumber string `json:"deviceModuleNumber"`
			DeviceHashGUID     string `json:"deviceHashGuid"`
			SummerHouse        int    `json:"summerHouse"`
			IAutoX             bool   `json:"iAutoX"`
			Nanoe              bool   `json:"nanoe"`
			AutoMode           bool   `json:"autoMode"`
			HeatMode           bool   `json:"heatMode"`
			FanMode            bool   `json:"fanMode"`
			DryMode            bool   `json:"dryMode"`
			CoolMode           bool   `json:"coolMode"`
			EcoNavi            bool   `json:"ecoNavi"`
			PowerfulMode       bool   `json:"powerfulMode"`
			QuietMode          bool   `json:"quietMode"`
			AirSwingLR         bool   `json:"airSwingLR"`
			AutoSwingUD        bool   `json:"autoSwingUD"`
			EcoFunction        int    `json:"ecoFunction"`
			TemperatureUnit    int    `json:"temperatureUnit"`
			ModeAvlList        struct {
				AutoMode int `json:"autoMode"`
				FanMode  int `json:"fanMode"`
			} `json:"modeAvlList"`
			CoordinableFlg bool            `json:"coordinableFlg"`
			Parameters     DeviceParamList `json:"parameters"`
		} `json:"deviceList"`
	} `json:"groupList"`
}

func addJsonElement(name string, value string) string {
	return `"` + name + `":"` + value + `"`
}

func (dlg DeviceGroupList) Print() {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "|\tGroup\t|\tDevice Name\t|\tDeviceGUID\t|")
	fmt.Fprintln(w, "|\t-\t|\t-\t|\t-\t|")

	for _, group := range dlg.GroupList {
		for _, device := range group.DeviceList {
			fmt.Fprintln(w, "|\t"+group.GroupName+"\t|\t"+device.DeviceName+"\t|\t"+device.DeviceGUID+"\t|")
		}
	}
	w.Flush()
}

func (dsf DeviceStatusFull) Print() {
	t := time.Unix(dsf.Timestamp/1000, 0)
	fmt.Println("Device Clock: " + t.Format("2006.01.02 15:04:05"))

	//color print
	var power string
	if dsf.Parameters.Operate == 0 {
		power = "Off"
	} else {
		power = "On"
	}

	var nanoe string
	if dsf.Parameters.Nanoe == 1 {
		nanoe = "Off"
	} else {
		nanoe = "On"
	}

	var eco string
	if dsf.Parameters.EcoMode == 0 {
		eco = "None"
	} else if dsf.Parameters.EcoMode == 1 {
		eco = "Strong"
	} else if dsf.Parameters.EcoMode == 2 {
		eco = "Quiet"
	}

	var fanSpeed string
	switch dsf.Parameters.FanSpeed {
	case 0:
		fanSpeed = "auto"
	case 1:
		fanSpeed = "1"
	case 2:
		fanSpeed = "2"
	case 3:
		fanSpeed = "3"
	case 4:
		fanSpeed = "4"
	case 5:
		fanSpeed = "5"
	}

	var airSwingUD string
	switch dsf.Parameters.FanSpeed {
	case 0:
		airSwingUD = "Top"
	case 1:
		airSwingUD = "Bottom"
	case 2:
		airSwingUD = "Center"
	case 3:
		airSwingUD = "Center-High"
	case 4:
		airSwingUD = "Center-Low"
	}

	var airSwingLR string
	switch dsf.Parameters.FanSpeed {
	case 0:
		airSwingLR = "Left"
	case 1:
		airSwingLR = "Right"
	case 2:
		airSwingLR = "Center"
	case 4:
		airSwingLR = "Center-Left"
	case 5:
		airSwingLR = "Center-Right"
	}

	switch dsf.Parameters.FanAutoMode {
	case 3:
		airSwingLR = "auto"
	case 2:
		airSwingUD = "auto"
	case 0:
		airSwingLR = "auto"
		airSwingUD = "auto"
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	//Graphs with bars or something
	fmt.Fprintln(w, "|\tPower: "+power+" \t|\t Nanoe: "+nanoe+" \t|\t Eco mode: "+eco+"\t|")

	fmt.Fprintln(w, "|\tFan Speed: "+fanSpeed+" \t|\t Vertical Angle: "+airSwingUD+" \t|\t Horizontal Angle: "+airSwingLR+"\t|")

	fmt.Fprintln(w, "|\tTarget temp: "+strconv.FormatFloat(dsf.Parameters.TemperatureSet, 'f', 1, 64)+"°C \t|\t Inside temp: "+strconv.Itoa(dsf.Parameters.InsideTemperature)+"°C  \t|\t Outside temp: "+strconv.Itoa(dsf.Parameters.OutTemperature)+"°C\t|")
	w.Flush()
}
