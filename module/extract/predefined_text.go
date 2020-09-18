package extract

import (
	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/cli"
)

func Predefined() *cli.Command {
	return &cli.Command{
		Name:        "predefined",
		Usage:       "Use a predefined regex to extract strings",
		Examples:    []cli.Example{},
		Description: `Use the predefined submodule to choose from a set of regexes, which you can use to extract strings either from a website, a file or stdin.`,
		Category:    categories.TextProcessing,
		Subcommands: GeneratePredefinedRegexCommands(),
	}
}
