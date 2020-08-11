package modules

import (
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/say"
	"github.com/urfave/cli/v2"
	"regexp"
)

type Module struct{}

func (Module) GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "modules",
			Aliases:     []string{"mods"},
			Usage:       "",
			Description: "List and search modules",
			Action: func(c *cli.Context) error {
				search := c.String("search")
				list := c.Bool("list")

				var foundModules []string

				r, err := regexp.Compile(search)
				if err != nil {
					return err
				}

				if search != "" {
					for _, m := range module.RegisteredModules {
						for _, cmd := range m.GetCommands() {
							if r.MatchString(cmd.Name) {
								foundModules = append(foundModules, cmd.Name)
							}
						}
					}
				} else if list {
					for _, m := range module.RegisteredModules {
						for _, cmd := range m.GetCommands() {
							foundModules = append(foundModules, cmd.Name)
						}
					}

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
					Usage:   "lists all files",
				},
			},
		},
	}
}
