package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetDeviceList() DeviceGroupList {
	verbosePrint("[+] Getting device list for " + GlobalConfig.Username)
	respBody := performHTTPRequest("GET", "/device/group", nil, GlobalConfig.RetryAttempts)

	var deviceStatus DeviceGroupList
	err := json.Unmarshal(respBody, &deviceStatus)
	if err != nil {
		verbosePrint("[!] Device list could not be read" + string(err.Error()))
	}
	verbosePrint("[+] Device list found for " + GlobalConfig.Username)
	return deviceStatus
}

func SetDeviceStatus(parameters string) {
	getDeviceId()
	verbosePrint("[+] Setting device status for " + GlobalConfig.DeviceGuid)

	controlJson := `{"deviceGuid":"` + GlobalConfig.DeviceGuid + `","parameters":{`
	controlJson += parameters
	controlJson += `}}`

	respBody := performHTTPRequest("POST", "/deviceStatus/control", []byte(controlJson), GlobalConfig.RetryAttempts)

	var result map[string]interface{}
	json.Unmarshal([]byte(respBody), &result)
	if result["result"] != nil && result["result"].(float64) == 0 {
		verbosePrint("[+] Device status set for " + GlobalConfig.DeviceGuid)
	} else if result["message"] != nil {
		verbosePrint("[!] Error getting setting device status: " + result["message"].(string))
	} else {
		verbosePrint("[!] Setting device status update may have failed or was only partial " + GlobalConfig.DeviceGuid)
	}
}

func GetDeviceStatus() DeviceStatusFull {
	getDeviceId()
	verbosePrint("[+] Getting device status for " + GlobalConfig.DeviceGuid)
	deviceId := strings.Replace(GlobalConfig.DeviceGuid, "+", "%2B", -1)
	respBody := performHTTPRequest("GET", "/deviceStatus/now/"+deviceId, nil, GlobalConfig.RetryAttempts)

	var deviceStatus DeviceStatusFull
	err := json.Unmarshal(respBody, &deviceStatus)
	if err != nil {
		verbosePrint("[!] Device status could not be read" + string(err.Error()))
	}
	verbosePrint("[+] Device status found for " + GlobalConfig.DeviceGuid)
	//visualize device status
	return deviceStatus
}

func getDeviceId() {
	if !(GlobalConfig.DeviceGuid == "") {
		return
	}
	verbosePrint("[+] DeviceID not found querying first device in list")

	respBody := performHTTPRequest("GET", "/device/group", nil, GlobalConfig.RetryAttempts)

	var result map[string]interface{}
	json.Unmarshal([]byte(respBody), &result)
	// if we got back a json with at least 1 registered device (this cannot be the right way to assert this)
	if result["groupList"] != nil &&
		result["groupList"].([]interface{})[0] != nil &&
		result["groupList"].([]interface{})[0].(map[string]interface{})["deviceList"] != nil &&
		result["groupList"].([]interface{})[0].(map[string]interface{})["deviceList"].([]interface{})[0] != nil &&
		result["groupList"].([]interface{})[0].(map[string]interface{})["deviceList"].([]interface{})[0].(map[string]interface{})["deviceGuid"] != nil {

		deviceGuid := result["groupList"].([]interface{})[0].(map[string]interface{})["deviceList"].([]interface{})[0].(map[string]interface{})["deviceGuid"].(string)
		verbosePrint("[+] Device ID " + deviceGuid + " chosen")
		GlobalConfig.DeviceGuid = deviceGuid
		overwriteConfigFile(GlobalConfig)
	} else if result["message"] != nil {
		verbosePrint("[!] Error getting DeviceID: " + result["message"].(string))

	} else {
		verbosePrint("[!] Getting DeviceID failed, reason unknown")
	}
}

func refreshLoginAuthToken() {
	verbosePrint("[+] Token Expired or invalid, trying Login as " + GlobalConfig.Username)
	postBody := []byte(`{"language":"0","loginId":"` + GlobalConfig.Username + `","password":"` + GlobalConfig.Password + `"}`)
	//login attempts will be repeated in the main request, should not be retried here, it would cause infinite recursion
	respBody := performHTTPRequest("POST", "/auth/login", postBody, -1)

	var result map[string]interface{}
	json.Unmarshal([]byte(respBody), &result)

	if result["uToken"] != nil {
		verbosePrint("[+] Login successful")
		GlobalConfig.Bearer = result["uToken"].(string)
		overwriteConfigFile(GlobalConfig)
		return
	} else if result["message"] != nil {
		verbosePrint("[!] Login failed")
		verbosePrint("[!] Error: " + result["message"].(string))
	}
	verbosePrint("[!] Login failed, reason unknown")
}

func performHTTPRequest(method string, reqURL string, body []byte, retryCount int) []byte {
	reqURL = "https://accsmart.panasonic.com" + reqURL

	// +1 is here in case login expired, and we need to get a new token
	for i := 0; i <= retryCount+1; i++ {
		req, _ := http.NewRequest(method, reqURL, bytes.NewBuffer(body))
		req.Host = "accsmart.panasonic.com"

		if GlobalConfig.Bearer != "" {
			req.Header.Set("X-User-Authorization", GlobalConfig.Bearer)
		}
		req.Header.Set("X-APP-TYPE", "1") //these two get lowercased when its actually sent
		req.Header.Set("X-APP-VERSION", "1.14.0")
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		req.Header.Set("Accept", "application/json; charset=utf-8")
		req.Header.Set("User-Agent", "G-RAC")

		tr := &http.Transport{}

		if GlobalConfig.HttpProxy != "" {
			/*Debug feature to Turn off cert validation */
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

			// this is for debug proxying
			proxy, _ := url.Parse(GlobalConfig.HttpProxy)
			tr = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				Proxy:           http.ProxyURL(proxy),
			}
		}

		//for avoiding infinite redirect loops
		client := &http.Client{
			Transport: tr,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}

		resp, err := client.Do(req)
		if err != nil {
			verbosePrint("[!] HTTP request failed to " + method + " " + reqURL + " Attempt: " + strconv.Itoa(i))
			verbosePrint(err.Error())
			continue
		}

		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)

		if GlobalConfig.HttpDebug {
			fmt.Printf("%s %s Attempt: %s\nRequest: %s \nResponse: %s \n", method, reqURL, strconv.Itoa(i), string(body), string(respBody))
		}

		if resp.StatusCode != 200 {
			//returned with <status code>
			verbosePrint("[!] HTTP request returned with " + strconv.Itoa(resp.StatusCode) + " : " + method + " " + reqURL)

			// if token was expired, get a new token and retry request (if not already a login request)
			if resp.StatusCode == 401 && retryCount != -1 {
				refreshLoginAuthToken()
			}

			// 500 error code would mean the request was wrong, but did not time out
			if resp.StatusCode != 500 {
				verbosePrint("[!] Retrying previous HTTP request")
				continue
			}
		}

		return []byte(respBody)
	}
	return nil
}
