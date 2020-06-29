package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

var Config ConfigStruct
var Secrets SecretsStruct

//Main Configuration Struct

type ConfigStruct struct { //TODO populate defaults, and create config/secret struct
	SqlConfig  SQLConfig
	AuthConfig AuthenticationConfig
}

type SQLConfig struct {
	Nodes string
}

type AuthenticationConfig struct {
	Auth0Domain   string //OpenID Provider URL
	Auth0ClientID string
}

//Main Secrets Struct

type SecretsStruct struct {
	SqlSecrets  SQLSecrets
	AuthSecrets AuthenticationSecrets
}

type SQLSecrets struct {
	Username string
	Password string
}

type AuthenticationSecrets struct {
	CookieStoreKeys   [][]byte
	Auth0ClientSecret string
}

var defaultConfig = ConfigStruct{}

var defaultSecrets = SecretsStruct{}

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

		return errors.New("Configuration file not set")
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
