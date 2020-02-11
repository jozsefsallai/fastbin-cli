package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "fastbin"
	app.Usage = "command line client for fastbin"
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		fmt.Println("This app is a work in progress.")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
