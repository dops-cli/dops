package extract

import (
	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/utils"
	"regexp"

	"github.com/dops-cli/dops/categories"
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
				{
					Name:        "text",
					Aliases:     []string{"t", "string", "strings", "s"},
					Usage:       "Extracts text from data",
					Examples:    nil,
					Description: "This can be used to extract text using a predefined or a custom regex.",
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
				},
			},
		},
	}
}
