package update

import (
	"runtime"

	"github.com/dops-cli/dops/cli"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/say/color"
)

// Module returns the created module
type Module struct{}

// GetModuleCommands returns the commands of the module
func (Module) GetModuleCommands() []*cli.Command {
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
					say.Text(color.SHiRed("iwr -useb dops-cli.com/get/windows | iex"))
				} else {
					cli.ShowVersion(c)
					say.Text("To update dops, open a terminal and run")
					say.Text(color.SHiRed("curl -s https://dops-cli.com/get/linux | sudo bash"))
				}

				return nil
			},
		},
	}
}
