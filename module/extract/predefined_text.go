package extract

import (
	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/cli"
)

func Predefined() *cli.Command {
	return &cli.Command{
		Name:        "predefined",
		Usage:       "",
		Examples:    []cli.Example{},
		Description: "",
		Category:    categories.TextProcessing,
		Subcommands: GeneratePredefinedRegexCommands(),
	}
}
