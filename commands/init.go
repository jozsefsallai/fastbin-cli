package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/jozsefsallai/fastbin-cli/config"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

// InitConfig asks the user for environment configuration details
// and writes the data to ~/.fastbinrc.json
func InitConfig(ctx *cli.Context) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the URL of the fastbin server you want to use: ")
	server, _ := reader.ReadString('\n')

	var conf config.Config
	conf.Server = strings.TrimSpace(server)

	json, _ := json.MarshalIndent(conf, "", "  ")
	home, _ := homedir.Dir()
	rcPath := path.Join(home, ".fastbinrc.json")

	err := ioutil.WriteFile(rcPath, json, 0644)
	if err != nil {
		return err
	}

	return nil
}
