package extract

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
			Name:        "extract",
			Usage:       "Extracts data from other data",
			Description: `Extract contains multiple data extractors, which can be used to extract data from a file, URL or stdin.`,
			Category:    categories.DataAnalysis,
			Subcommands: []*cli.Command{
				Text(),
			},
		},
	}
}
