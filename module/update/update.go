package update

import (
	"runtime"

	"github.com/pterm/pterm"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/cli"
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

				cli.ShowVersion(c)

				if runtime.GOOS == "windows" {
					pterm.Info.Println("To update dops, " +
						"open a new powershell with admin privileges and run:")
					pterm.Println(pterm.LightRed("       iwr -useb dops-cli.com/get/windows | iex"))
				} else if runtime.GOOS == "darwin" {
					pterm.Info.Println("To update dops, open a new terminal and run:")
					pterm.Println(pterm.LightRed(`       /bin/bash -c "$(curl -fsSL https://dops-cli.com/get/linux)"`))
				} else {
					pterm.Info.Println("To update dops, open a terminal and run:")
					pterm.Println(pterm.LightRed("       curl -s https://dops-cli.com/get/linux | sudo bash"))
				}

				return nil
			},
		},
	}
}
