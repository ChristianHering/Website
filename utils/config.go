package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

func setupConfig() error {
	Config = defaultConfig
	err := getConfig("config.json", &Config)
	if err != nil {
		return err
	}

	Secrets = defaultSecrets
	err = getConfig("secrets.json", &Secrets)
	if err != nil {
		return err
	}

	return nil
}

var Config ConfigStruct
var Secrets SecretsStruct

type ConfigStruct struct { //TODO populate defaults, and create config/secret struct
	SqlConfig SQLConfig
}

type SQLConfig struct {
	Nodes string
}

type SecretsStruct struct {
	SqlSecrets SQLSecrets
}

type SQLSecrets struct {
	Username string
	Password string
}

var defaultConfig = ConfigStruct{}

var defaultSecrets = SecretsStruct{}

//Gets the configuration from a file name or creates
//a new config file if one doesn't already exist
//
//To use pass a pointer to a struct initialized with default values
func getConfig(configFileName string, configPointer interface{}) error {
	if fileExists(configFileName) { //Get existing configuration from configFileName
		b, err := ioutil.ReadFile(configFileName)
		if err != nil {
			return errors.WithStack(err)
		}

		err = json.Unmarshal(b, configPointer)
		if err != nil {
			fmt.Println("Failed to unmarshal configuration file")
			return errors.WithStack(err)
		}

		return nil
	} else { //If configFileName doesn't exist, create a new config file
		b, err := json.MarshalIndent(configPointer, "", " ")
		if err != nil {
			fmt.Println("Failed to marshal configuration file")
			return errors.WithStack(err)
		}

		err = ioutil.WriteFile(configFileName, b, 0644)
		if err != nil {
			fmt.Println("Failed to write configuration file")
			return errors.WithStack(err)
		}

		return nil
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
