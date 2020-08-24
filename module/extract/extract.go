package extract

import (
	"regexp"

	"github.com/dops-cli/dops/cli"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/utils"
)

// Module returns the created module
type Module struct{}

// GetCommands returns the commands of the module
func (Module) GetModuleCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "extract-text",
			Usage:       "Extracts text using regex from a file",
			Description: `Extract-text can be used to extract text from a file using regex patterns.`,
			Category:    categories.TextProcessing,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "regex",
					Aliases: []string{"r"},
					Usage:   "extracts matching strings with `PATTERN`",
				},
				&cli.PathFlag{
					Name:      "input",
					Aliases:   []string{"i"},
					Usage:     "use `FILE` as input",
					TakesFile: true,
				},
				&cli.StringFlag{
					Name:    "output",
					Aliases: []string{"o"},
					Usage:   "outputs to directory `DIR`",
				},
				&cli.BoolFlag{
					Name:    "append",
					Aliases: []string{"a"},
					Usage:   "append instead of overriding output",
				},
			},
			Action: func(c *cli.Context) error {
				regex := c.String("regex")
				input := c.Path("input")
				output := c.String("output")
				appendFlag := c.Bool("appendFlag")

				var foundStrings []string

				r, err := regexp.Compile(regex)
				if err != nil {
					return err
				}

				foundStrings = r.FindAllString(utils.Input(input), -1)
				utils.Output(output, foundStrings, appendFlag)

				return nil
			},
		},
	}
}
