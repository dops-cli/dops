package update

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

type Module struct{}

func (Module) GetFlags() []cli.Flag {
	return []cli.Flag{}
}

func (Module) GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "update",
			Usage:       "",
			Description: "Updates dops",
			Action: func(c *cli.Context) error {
				fmt.Println("Automatic updates are not supported currently.")
				fmt.Println("Please visit https://github.com/dops-cli/dops/releases to download the current version.")
				cli.ShowVersion(c)
				return nil
			},
		},
	}
}
