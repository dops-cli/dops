package extract

import (
	"fmt"
	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/utils"
	"regexp"
)

type PredefinedRegexCommand struct {
	Name    string
	Usage   string
	Aliases []string
	Regex   string
	Matches []string
	Fails   []string
}

func GeneratePredefinedRegexCommands() []*cli.Command {
	var list []*cli.Command

	for _, c := range RegexList {
		cmd := &cli.Command{
			Name:    c.Name,
			Aliases: c.Aliases,
			Usage:   c.Usage,
			Examples: []cli.Example{
				{
					ShortDescription: fmt.Sprintf("Extract all %s´s from INPUT", c.Name),
					Usage:            fmt.Sprintf("dops extract text predefined %s --input file.txt", c.Name),
				},
			},
			Description: fmt.Sprintf("The %s command finds all %s`s in the input and returns them.\n\n", c.Name, c.Name) + "Regex: \n" + c.Regex,
			Category:    categories.TextProcessing,
			Action: func(context *cli.Context) error {
				input := utils.Input(context.Path("input"))
				output := context.String("output")
				appendFlag := context.Bool("appendFlag")

				var foundStrings []string

				r, err := regexp.Compile(c.Regex)
				if err != nil {
					return err
				}

				foundStrings = r.FindAllString(input, -1)
				utils.Output(output, foundStrings, appendFlag)

				return nil
			},
			Flags: InputOutputFlags(),
		}
		list = append(list, cmd)
	}

	return list
}

func InputOutputFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Aliases: []string{"i"},
			Name:    "input",
			Usage:   "Input accepts a file, URL or stdin if not set",
		},
		&cli.StringFlag{
			Aliases: []string{"o"},
			Name:    "output",
			Usage:   "Writes to a file, if not set it writes to stdout",
		},
		&cli.BoolFlag{
			Name:    "append",
			Aliases: []string{"a"},
			Usage:   "append instead of overriding output",
		},
	}
}
