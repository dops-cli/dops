package template

import (
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/say/color"
	. "github.com/dops-cli/dops/say/color"
	"github.com/urfave/cli/v2"
	"strings"
	"text/template"
)

type Modules struct {
	Commands cli.Commands
}

var modules = `{{range .Commands}}` +
	Primary("\n{{.Name}}") + ` - ` + Secondary("{{.Usage}}") + `

  ` + Primary("Usage:") + ` {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
  {{if .Aliases}}` + Primary("Aliases:") + `  {{join .Aliases ", "}}{{end}}
  {{if .Category}}` + Primary("Category:") + ` {{.Category}}{{end}}{{if .Description}}

` + Section("Description") + `
  {{.Description}}{{end}}{{if .VisibleFlags}}

` + Section("Options") + `
  {{range .VisibleFlags}}` + Flag("{{.}}") + `
  {{end}}{{end}}` + "\n\n" + `{{end}}`

func PrintModules() error {

	var commands []*cli.Command

	for _, m := range module.ActiveModules {
		for _, cmd := range m.GetCommands() {
			commands = append(commands, cmd)
		}
	}

	funcMap := template.FuncMap{"join": strings.Join}

	t := template.Must(template.New("modules").Funcs(funcMap).Parse(modules))

	err := t.Execute(color.Output, Modules{commands})
	if err != nil {
		return err
	}

	return nil
}
