package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

//Config global configuration var
var Config ConfigStruct

//ConfigStruct Main Configuration Struct
type ConfigStruct struct { //TODO populate defaults, and create config/secret struct
	SQLConfig            SQLConf
	AuthenticationConfig AuthenticationConf
	EmailConfig          EmailConf
	MaxCacheAge          string //Cache Control setting in seconds
}

//SQLConf Configuration struct for SQL
type SQLConf struct {
	Host     string
	Username string
	Password string
	Database string
}

//AuthenticationConf Configuration struct for Auth0
type AuthenticationConf struct {
	Auth0Domain       string //OpenID Provider URL
	Auth0ClientID     string
	CookieStoreKeys   [][]byte
	Auth0ClientSecret string
}

//EmailConf Stores SMTP authentication details
type EmailConf struct {
	Host     string
	Port     int
	Username string
	Password string
}

func setupConfig() error {
	Config = ConfigStruct{}
	err := getConfig("config.json", &Config)
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
	}

	//If configFileName doesn't exist, create a new config file
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

//Check to see if a file exists by name. Return bool
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
