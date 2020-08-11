package update

import (
	"github.com/dops-cli/dops/say"
	"github.com/urfave/cli/v2"
)

type Module struct{}

func (Module) GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "update",
			Usage:       "",
			Description: "Updates dops",
			Action: func(c *cli.Context) error {
				say.Text("Automatic updates are not supported currently.")
				say.Text("Please visit https://github.com/dops-cli/dops/releases to download the current version.")
				cli.ShowVersion(c)
				return nil
			},
		},
	}
}
