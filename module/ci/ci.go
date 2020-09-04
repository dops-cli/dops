package ci

import (
	"io/ioutil"
	"regexp"
	"sort"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/utils"
)

// Module returns the created module
type Module struct{}

// GetModuleCommands returns the commands of the module
func (Module) GetModuleCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "ci",
			Usage:       "Runs on every push to the official GitHub repository of dops",
			Description: "",
			Warning:     "This module should only be used to develop dops further",
			Category:    categories.Dops,
			Action: func(context *cli.Context) error {

				say.Info("Generating documentation for modules...")

				var commands []*cli.Command

				for _, m := range cli.ActiveModules {
					commands = append(commands, m.GetModuleCommands()...)
				}

				sort.Sort(cli.CommandsByName(commands))

				for _, cmd := range commands {
					say.Info(" - Generating documentation for", cmd.Name)
					doc := cli.CommandDocumentation(cmd, nil)
					err := ioutil.WriteFile("./docs/modules/"+cmd.Name+".md", []byte(doc), 0600)
					if err != nil {
						return err
					}
				}

				sidebarPath := "./docs/_sidebar.md"
				sidebarContentByte, err := ioutil.ReadFile(sidebarPath)
				if err != nil {
					return err
				}

				sidebarContent := string(sidebarContentByte)

				beforeRegex := regexp.MustCompile(`(?ms).*<!-- <<<CI-MODULES-START>> -->`)
				afterRegex := regexp.MustCompile(`(?ms)<!-- <<<CI-MODULES-END>> -->.*`)

				before := beforeRegex.FindAllString(sidebarContent, 1)[0]
				after := afterRegex.FindAllString(sidebarContent, 1)[0]

				var newSidebarContent string

				newSidebarContent += before + "\n"

				for _, cmd := range commands {
					newSidebarContent += "  - [" + cmd.Name + "](modules/" + cmd.Name + ".md)\n"
				}

				newSidebarContent += after

				utils.WriteFile(sidebarPath, []byte(newSidebarContent), false)

				say.Success("Documentation successfully generated!")

				return nil
			},
			Hidden: true,
		},
	}
}
