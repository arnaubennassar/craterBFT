package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	AppName    = "craterBFT"
	AppVersion = "v0.0.1"
)

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = AppName
	cliApp.Version = AppVersion

	cliApp.Commands = []*cli.Command{
		{
			Name:   "start",
			Usage:  "Start the node",
			Action: StartNode,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "home",
					Usage:    "Home directory `DIR`",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "jwt",
					Usage:    "JWT file for engine API security",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "ethUrl",
					Usage:    "URL to engine's public API",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "engineUrl",
					Usage:    "URL to engine's secure API",
					Required: true,
				},
			},
		},
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
