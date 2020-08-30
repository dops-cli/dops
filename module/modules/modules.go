package modules

import (
	"regexp"
	"sort"
	"strconv"

	"github.com/dops-cli/dops/cli"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/say"
)

// Module returns the created module
type Module struct{}

// GetModuleCommands returns the commands of the module
func (Module) GetModuleCommands() []*cli.Command {
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
				markdown := c.Bool("markdown")
				count := c.Bool("count")
				gd := c.Bool("generate-docs")

				cli.IncompatibleFlags(search, list, markdown, count, gd)

				var foundModules []string

				r, err := regexp.Compile(search)
				if err != nil {
					return err
				}

				if search != "" {
					for _, m := range cli.ActiveModules {
						for _, cmd := range m.GetModuleCommands() {
							if r.MatchString(cmd.Name) {
								foundModules = append(foundModules, cmd.Name)
							}
						}
					}
				} else if list {
					for _, m := range cli.ActiveModules {
						for _, cmd := range m.GetModuleCommands() {
							foundModules = append(foundModules, cmd.Name)
						}
					}
				} else if markdown {
					err := cli.PrintModulesMarkdown()
					if err != nil {
						return err
					}
					return nil
				} else if count {
					say.Text(strconv.Itoa(len(cli.ActiveModules) + 2))
					return nil
				} else if gd {
					// err := cli.GenerateDocs()
					// if err != nil {
					// 	return err
					// }
					// return nil
				}

				sort.Strings(foundModules)

				for _, name := range foundModules {
					say.Text(name)
				}

				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "search",
					Aliases: []string{"s"},
					Usage:   "Searches for `MODULE` using regex",
				},
				&cli.BoolFlag{
					Name:    "list",
					Aliases: []string{"l", "ls"},
					Usage:   "Lists all modules",
				},
				&cli.BoolFlag{
					Name:    "markdown",
					Aliases: []string{"m", "md"},
					Usage:   "Describes all modules with markdown output",
				},
				&cli.BoolFlag{
					Name:    "generate-docs",
					Aliases: []string{"gen-doc", "gd"},
					Usage:   "Generate the markdown for the official documentation of dops",
				},
				&cli.BoolFlag{
					Name:    "count",
					Aliases: []string{"c"},
					Usage:   "Returns the total module count of dops",
				},
			},
		},
	}
}
