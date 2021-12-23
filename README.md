# Panasonic-AC-CLI
CLI tool for controlling your Panasonic air conditioner (CS-Z25VKEW)

![demo](./demo.gif)

## Setup

You do have to install the [Panasonic app](https://play.google.com/store/apps/details?id=com.panasonic.ACCsmart) and pair with your devices once. It will request that you register on the [Panasonic site](https://csapl.pcpf.panasonic.com/Account/Register001?lang=en). Pair your devices following the in-app instructions, and save your credentials.

Once you have done that, you can set up `config.json`:
```
{
 "Username": "email@panasonicid.com", // Your registered user ID
 "Password": "YourPanasonicPassword", // Password for the registered email
 "Bearer": "",                        // Leave empty, it will be filled by the script
 "DeviceGuid": "",                    // You can fill this later, from the ouput of -list
 "RetryAttempts": 2,                  // If a HTTP request fails, how many times to retry
 "HttpDebug": false,                  // Print HTTP requests and responses
 "HttpProxy": "" ,                    // HTTP proxy for debugging
 "Verbose": false                     // Does nothing, use -q switch instead
}
```
After this you can delete the app from your phone, the script will log you out anyway.

## Build

If you have Go installed you can build the project as as
```
cd Panasonic-AC-CLI
go build .
```

Otherwise, just get the binary from the releases page.

## Usage

List all devices associated with your account
```
./panac -list
```
Select a device
```
./panac -dev <DeviceGUID>
```
Read current settings of a device
```
./panac -status
```
Set new status for device, you can set any combination of settings to be updated
```
./panac -power on -m cool -t 22 -fan auto
```
The rest of the options are
```
-ah string
   Horizontal angle of the air > auto/left/centerleft/center/centerright/right
-av string
   Vertical angle of the air > auto/top/centertop/center/centerbottom/bottom
-config string
   Path to config file (default "config.json")
-dev string
   DeviceGUID use -list to print avaliable
-eco string
   Economy setting > none/quiet/strong
-fan string
   Fan speed > auto/1/2/3/4/5
-h	Display this help text
-help
   Display detailed information
-list
   List devices associated with user
-m string
   AC mode > auto/heat/cool/dry/nanoe
-nanoe string
   Enabling nanoe setting but not nanoe mode > On/Off
-power string
   Power > On/Off
-q	Don't print status update messages
-status
   Display current status of device
-t string
   TemperatureSet in Celsius, supports 1 decimal > 22.5
```
