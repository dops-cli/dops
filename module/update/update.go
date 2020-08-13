package update

import (
	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/say"
	"github.com/urfave/cli/v2"
)

type Module struct{}

func (Module) GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "update",
			Usage:       "Updates the dops tool",
			Description: "NOTICE: This module is in progress. But you can already see it's usage for further use!",
			Category:    categories.Dops,
			Action: func(c *cli.Context) error {
				say.Text("Automatic updates are not supported currently.")
				say.Text("Please visit https://github.com/dops-cli/dops/releases to download the current version.")
				cli.ShowVersion(c)
				return nil
			},
		},
	}
}
