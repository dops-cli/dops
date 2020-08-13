package template

import (
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/say/color"
	. "github.com/dops-cli/dops/say/color"
	"github.com/urfave/cli/v2"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{"join": strings.Join}

type Modules struct {
	Commands cli.Commands
}

func PrintModules() error {

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

	var commands []*cli.Command

	for _, m := range module.ActiveModules {
		for _, cmd := range m.GetCommands() {
			commands = append(commands, cmd)
		}
	}

	t := template.Must(template.New("modules").Funcs(funcMap).Parse(modules))

	err := t.Execute(color.Output, Modules{commands})
	if err != nil {
		return err
	}

	return nil
}

func PrintModulesMarkdown() error {
	var modules = `{{range .Commands}}
### {{.Name}}  

> {{.Usage}}  

Usage: {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}<br/>
{{if .Aliases}}Aliases: ` + "`" + `{{join .Aliases "` + "`, `" + `"}}` + "`" + `{{if .Category}}\{{end}}{{end}}
{{if .Category}}Category: {{.Category}}{{end}}

{{if .Description}} #### Description

{{.Description}}{{end}}

{{if .VisibleFlags}}#### Options

` + "```" + `
{{range .VisibleFlags}}{{.}}
{{end}}` + "```" + `{{end}}
{{end}}`

	var commands []*cli.Command

	for _, m := range module.ActiveModules {
		for _, cmd := range m.GetCommands() {
			commands = append(commands, cmd)
		}
	}

	t := template.Must(template.New("modules").Funcs(funcMap).Parse(modules))

	err := t.Execute(color.Output, Modules{commands})
	if err != nil {
		return err
	}

	return nil
}
