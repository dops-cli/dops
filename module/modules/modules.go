package modules

import (
	"regexp"
	"strconv"

	"github.com/dops-cli/dops/cli"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/template"
)

// Module returns the created module
type Module struct{}

// GetCommands returns the commands of the module
func (Module) GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "modules",
			Aliases: []string{"mods"},
			Usage:   "List and search modules",
			Description: `The 'modules' command, is used to list and search modules in dops.
Furthermore, 'modules' can output all modules and their descriptions at the same time. 
With the 'markdown' flag, the output text is parsed in markdown. This is for example used in the DOPS-CI toolchain to generate the 'MODULES.md' file.`,
			Category: categories.Dops,
			Action: func(c *cli.Context) error {
				search := c.String("search")
				list := c.Bool("list")
				describe := c.Bool("describe")
				markdown := c.Bool("markdown")
				count := c.Bool("count")

				var foundModules []string

				r, err := regexp.Compile(search)
				if err != nil {
					return err
				}

				if search != "" {
					for _, m := range module.ActiveModules {
						for _, cmd := range m.GetCommands() {
							if r.MatchString(cmd.Name) {
								foundModules = append(foundModules, cmd.Name)
							}
						}
					}
				} else if list {
					for _, m := range module.ActiveModules {
						for _, cmd := range m.GetCommands() {
							foundModules = append(foundModules, cmd.Name)
						}
					}
				} else if describe {
					err := template.PrintModules()
					if err != nil {
						return err
					}
					return nil
				} else if markdown {
					err := template.PrintModulesMarkdown()
					if err != nil {
						return err
					}
					return nil
				} else if count {
					say.Text(strconv.Itoa(len(module.ActiveModules) + 2))
					return nil
				}

				for _, name := range foundModules {
					say.Text(name)
				}

				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "search",
					Aliases: []string{"s"},
					Usage:   "searches for `MODULE` using regex",
				},
				&cli.BoolFlag{
					Name:    "list",
					Aliases: []string{"l", "ls"},
					Usage:   "lists all modules",
				},
				&cli.BoolFlag{
					Name:    "describe",
					Aliases: []string{"d"},
					Usage:   "describes all modules",
				},
				&cli.BoolFlag{
					Name:    "markdown",
					Aliases: []string{"m", "md"},
					Usage:   "describes all modules with markdown output",
				},
				&cli.BoolFlag{
					Name:    "count",
					Aliases: []string{"c"},
					Usage:   "counts all modules",
				},
			},
		},
	}
}
