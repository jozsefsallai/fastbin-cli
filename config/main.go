package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

// Config is a structure defining the schema of the configuration (JSON)
type Config struct {
	Server string `json:"server"`
	Key    string `json:"key"`
}

// GetConfig returns the configuration object from the user's home directory
func GetConfigFunc() Config {
	home, _ := homedir.Dir()
	rcPath := path.Join(home, ".fastbinrc.json")

	if _, err := os.Stat(rcPath); os.IsNotExist(err) {
		fmt.Println("Configuration file not found. Please create a .fastbinrc.json file in your home directory.")
		os.Exit(1)
	}

	rc, err := os.Open(rcPath)

	if err != nil {
		log.Fatal("Failed to open ~/.fastbinrc.json")
	}

	defer rc.Close()

	contents, err := ioutil.ReadAll(rc)

	var config Config
	err = json.Unmarshal(contents, &config)

	if err != nil {
		log.Fatal("Failed to read configuration file. It is probably broken or invalid.")
	}

	return config
}
