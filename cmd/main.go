package cmd

import (
	"log"

	"github.com/urfave/cli/v2"
)

func Start(args []string) {
	app := cli.App{
		Name:        "start",
		Description: "go-template app",
		Commands: []*cli.Command{
			startCommand,
		},
	}

	if err := app.Run(args); err != nil {
		log.Println(err)
	}
}
