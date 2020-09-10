package open

import "github.com/dops-cli/dops/cli"

// Module returns the created module
type Module struct{}

// GetModuleCommands returns the commands of the module
func (Module) GetModuleCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "open",
			Usage:       "Opens a file or URL in the default program assigned to it",
			Description: "",
			Category:    "",
			Subcommands: []*cli.Command{
				URL(),
			},
		},
	}
}
