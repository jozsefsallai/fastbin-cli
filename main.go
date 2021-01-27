package main

import (
	"log"
	"os"

	"github.com/jozsefsallai/fastbin-cli/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "fastbin"
	app.Usage = "command line client for fastbin"
	app.Version = "1.1.0"
	app.Action = commands.CreateSnippet
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "full, f",
			Usage: "only print out and copy the full URL",
		},
		cli.BoolFlag{
			Name:  "raw, r",
			Usage: "only print out and copy the raw URL",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "initialize your fastbin CLI environment",
			Action: commands.InitConfig,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
