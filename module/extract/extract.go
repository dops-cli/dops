package extract

import (
	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/say"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"regexp"
)

type Module struct{}

func (Module) GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:     "extract-text",
			Usage:    "extracts text using regex from a file",
			Category: categories.TextProcessing,
			Flags: []cli.Flag{&cli.StringFlag{
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
					Name:    "stdout",
					Aliases: []string{"s"},
					Usage:   "prints output to stdout instead of writing to a file",
				}},
			Action: func(c *cli.Context) error {
				regex := c.String("regex")
				input := c.Path("input")
				output := c.String("output")
				stdout := c.Bool("stdout")

				r, err := regexp.Compile(regex)
				if err != nil {
					return err
				}

				file, err := ioutil.ReadFile(input)
				if err != nil {
					return err
				}

				foundStrings := r.FindAllString(string(file), -1)

				if stdout {
					for _, s := range foundStrings {
						say.Text(s)
					}
				} else {
					var out string
					for _, s := range foundStrings {
						out += s + "\n"
					}
					err := ioutil.WriteFile(output, []byte(out), os.ModeAppend)
					if err != nil {
						return err
					}
				}
				return nil
			},
		},
	}
}
