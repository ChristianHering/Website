package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func setConfig() error {
	value, err := getConfig("config.json", defaultConfig)
	if err != nil {
		return err
	}
	Config = value.Interface().(ConfigStruct)

	value, err = getConfig("secrets.json", defaultSecrets)
	if err != nil {
		return err
	}
	Secrets = value.Interface().(SecretsStruct)

	return nil
}

var Config ConfigStruct
var Secrets SecretsStruct

type ConfigStruct struct { //TODO populate defaults, and create config/secret struct
}

type SecretsStruct struct {
}

var defaultConfig = ConfigStruct{}

var defaultSecrets = SecretsStruct{}

//Gets the configuration from a file name or creates
//a new config file if one doesn't already exist
//
//To use fill a struct or use the return value, call
//configuration.Interface().(TheTypeOfDefaultConfig)
func getConfig(configFileName string, defaultConfig interface{}) (configuration reflect.Value, err error) {
	if fileExists(configFileName) { //Get existing configuration from configFileName
		b, err := ioutil.ReadFile(configFileName)

		config := reflect.ValueOf(defaultConfig) //We're only doing this for it's type

		err = json.Unmarshal(b, &config)
		if err != nil {
			fmt.Println("Failed to unmarshal configuration file")
			return reflect.Value{}, err
		}

		return config, nil
	} else { //If configFileName doesn't exist, create a new config file
		config := reflect.ValueOf(defaultConfig)

		b, err := json.MarshalIndent(config.Interface(), "", " ")
		if err != nil {
			fmt.Println("Failed to marshal configuration file")
			return reflect.Value{}, err
		}

		err = ioutil.WriteFile(configFileName, b, 0644)
		if err != nil {
			fmt.Println("Failed to write configuration file")
			return reflect.Value{}, err
		}

		return config, nil //Return default configuration
	}
}

//Check to see if a file exists by name. Return bool
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
