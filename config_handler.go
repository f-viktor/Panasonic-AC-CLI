package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ConfigFile struct {
	Username      string `json:"Username"`
	Password      string `json:"Password"`
	Bearer        string `json:"Bearer"`
	DeviceGuid    string `json:"DeviceGuid"`
	RetryAttempts int    `json:"RetryAttempts"`
	Verbose       bool   `json:"Verbose"`
	HttpDebug     bool   `json:"HttpDebug"`
	HttpProxy     string `json:"HttpProxy"`
}

var GlobalConfig ConfigFile

func ReadConfig(path string) {
	//open config file
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)

	//decode config into user structure
	conf := ConfigFile{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Fatalln(err)
	}
	GlobalConfig = conf
}

// unused was only in use when we I tried to use autologin cookies
func OverwriteConfigFile(conf ConfigFile) {

	file, _ := json.MarshalIndent(conf, "", " ")

	_ = ioutil.WriteFile("config.json", file, 0644)
}

func verbosePrint(message string) {
	if GlobalConfig.Verbose {
		fmt.Println(message)
	}
}
