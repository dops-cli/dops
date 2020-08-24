package template

import (
	"strings"
	"text/template"

	"github.com/dops-cli/dops/cli"

	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/say/color"
)

var funcMap = template.FuncMap{"join": strings.Join}

// Modules is a wrapper for cli.Commands
type Modules struct {
	Commands cli.Commands
}

// PrintModules prints all modules to stdout
func PrintModules() error {

	var modules = `{{range .Commands}}` +
		color.Primary("\n{{.Name}}") + ` - ` + color.Secondary("{{.Usage}}") + `

  ` + color.Primary("Usage:") + ` {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
  {{if .Aliases}}` + color.Primary("Aliases:") + `  {{join .Aliases ", "}}{{end}}
  {{if .Category}}` + color.Primary("Category:") + ` {{.Category}}{{end}}{{if .Description}}

` + color.Section("Description") + `
{{.Description}}{{end}}{{if .VisibleFlags}}

` + color.Section("Options") + `
  {{range .VisibleFlags}}` + color.Flag("{{.}}") + `
  {{end}}{{end}}` + "\n\n" + `{{end}}`

	var commands []*cli.Command

	for _, m := range module.ActiveModules {
		commands = append(commands, m.GetCommands()...)
	}

	t := template.Must(template.New("modules").Funcs(funcMap).Parse(modules))

	err := t.Execute(color.Output, Modules{commands})
	if err != nil {
		return err
	}

	return nil
}

// PrintModulesMarkdown prints all modules in markdown format to stdout
func PrintModulesMarkdown() error {
	var modules = `# DOPS - Modules{{range .Commands}}
## {{.Name}}  

> {{.Usage}}  

Usage: {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}<br/>
{{if .Aliases}}Aliases: ` + "`" + `{{join .Aliases "` + "`, `" + `"}}` + "`" + `{{if .Category}}<br/>{{end}}{{end}}
{{if .Category}}Category: {{.Category}}{{end}}

{{if .Description}} ### Description

{{.Description}}{{end}}

{{if .VisibleFlags}}### Options

` + "```" + `
{{range .VisibleFlags}}{{.}}
{{end}}` + "```" + `{{end}}
{{end}}`

	var commands []*cli.Command

	for _, m := range module.ActiveModules {
		commands = append(commands, m.GetCommands()...)
	}

	t := template.Must(template.New("modules").Funcs(funcMap).Parse(modules))

	err := t.Execute(color.Output, Modules{commands})
	if err != nil {
		return err
	}

	return nil
}
