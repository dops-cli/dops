package extract

import (
	"github.com/dops-cli/dops/global/options"
	"regexp"

	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/utils"
)

// Text module
func Text() *cli.Command {
	return &cli.Command{
		Name:        "text",
		Aliases:     []string{"t", "string", "strings", "s"},
		Usage:       "Extracts text from data",
		Examples:    nil,
		Description: "This can be used to extract text using a predefined or a custom regex.",
		Subcommands: []*cli.Command{
			Predefined(),
		},
		Action: func(c *cli.Context) error {
			regex := c.String("regex")
			input := utils.Input(c.Path("input"))
			output := c.String("output")
			glob := c.String("glob")
			appendFlag := c.Bool("appendFlag")

			cli.IncompatibleFlags(input, glob)
			cli.IncompatibleFlags(appendFlag, glob)

			var foundStrings []string

			r, err := regexp.Compile(regex)
			if err != nil {
				return err
			}

			if glob != "" {
				err = utils.Glob(glob, func(path string) error {
					input = path
					if options.Verbose {
						foundStrings = append(foundStrings, "\n\n"+input+"\n\n")
					}
					foundStrings = append(foundStrings, r.FindAllString(input, -1)...)
					utils.Output(output, foundStrings, true)
					return nil
				})
				if err != nil {
					return err
				}
			} else if input != "" {
				if options.Verbose {
					foundStrings = append(foundStrings, "\n\n"+input+"\n\n")
				}
				foundStrings = append(foundStrings, r.FindAllString(input, -1)...)
				utils.Output(output, foundStrings, appendFlag)
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "regex",
				Aliases: []string{"r"},
				Usage:   "extracts matching strings with `PATTERN`",
			},
			&cli.StringFlag{
				Name:    "glob",
				Aliases: []string{"g"},
				Usage:   "uses a `GLOB` pattern to input multiple files",
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
	}
}
