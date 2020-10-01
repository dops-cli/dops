package image

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
			Name:        "image",
			Usage:       "image modification",
			Description: `This module has a list of modules to modify images.`,
			Category:    categories.ImageProcessing,
			Subcommands: []*cli.Command{
				Watermark(),
			},
		},
	}
}
