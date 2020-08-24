package update

import (
	"runtime"

	"github.com/urfave/cli/v2"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/say/color"
)

// Module returns the created module
type Module struct{}

// GetCommands returns the commands of the module
func (Module) GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "update",
			Usage:       "Updates the dops tool",
			Description: "NOTICE: This module is in progress. But you can already see it's usage for further use!",
			Category:    categories.Dops,
			Action: func(c *cli.Context) error {

				if runtime.GOOS == "windows" {
					cli.ShowVersion(c)
					say.Text(color.Primary("To update dops, open a new powershell with admin privileges and run:"))
					say.Text(color.New(color.BgBlack).Sprint("iwr -useb dops-cli.com/get/windows | iex"))
				} else {
					cli.ShowVersion(c)
					say.Warning("Automatic updates are not supported for " + runtime.GOOS + ", yet.")
					say.Text("Please visit https://github.com/dops-cli/dops/releases to download the latest version.")
				}

				return nil
			},
		},
	}
}
