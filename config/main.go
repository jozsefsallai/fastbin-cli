package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

// Config is a structure defining the schema of the configuration (JSON)
type Config struct {
	Server string `json:"server"`
}

// GetConfig returns the configuration object from the user's home directory
func GetConfig() Config {
	home, _ := homedir.Dir()
	rcPath := path.Join(home, ".fastbinrc.json")

	if _, err := os.Stat(rcPath); os.IsNotExist(err) {
		fmt.Println("Configuration file not found. Please create a .fastbinrc.json file in your home directory.")
		os.Exit(1)
	}

	rc, err := os.Open(rcPath)

	if err != nil {
		panic("Failed to open ~/.fastbinrc.json")
	}

	defer rc.Close()

	contents, _ := ioutil.ReadAll(rc)

	var config Config
	json.Unmarshal(contents, &config)

	if err != nil {
		panic("Failed to read configuration file. It is probably broken or invalid.")
	}

	return config
}
