package open

import (
	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/cli"
)

// Module returns the created module
type Module struct{}

// GetModuleCommands returns the commands of the module
func (Module) GetModuleCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "open",
			Usage:       "Opens a file or URL in the default program assigned to it",
			Description: "Open finds the standard program, assigned to a file, and opens it with the found programm.",
			Category:    categories.Execute,
			Subcommands: []*cli.Command{
				URL(),
			},
		},
	}
}
