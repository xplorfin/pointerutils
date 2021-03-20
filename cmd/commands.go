package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/xplorfin/go-template/internal"
)

var startCommand = &cli.Command{
	Name:    "start",
	Aliases: []string{"s"},
	Action: func(c *cli.Context) error {
		internal.Setup()
		return nil
	},
}
