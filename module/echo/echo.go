package echo

import (
	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/utils"
	"strings"
)

// Module returns the created module
type Module struct{}

// GetModuleCommands returns the commands of the module
func (Module) GetModuleCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "echo",
			Usage:       "Displays input text in the console",
			Examples:    nil,
			Description: "Echo displays text from stdin, a file or from the arguments, to the user.",
			Category:    categories.TextProcessing,
			Action: func(context *cli.Context) error {
				var input string

				if context.String("input") != "" {
					input = utils.Input(context.String("input"))
				} else {
					input = strings.Join(context.Args().Slice(), " ")
				}

				say.Text(input)

				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Aliases: []string{"i"},
					Name:    "input",
					Usage:   "Input accepts a file or an URL",
				},
			},
		},
	}
}
