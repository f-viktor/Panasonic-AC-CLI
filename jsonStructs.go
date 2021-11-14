package main

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
