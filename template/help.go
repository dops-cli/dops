package template

import (
	. "github.com/dops-cli/dops/say/color"
)

// CommandHelp contains the template of dops modules help text.
var CommandHelp = Primary("\n{{.Name}}") + ` - ` + Secondary("{{.Usage}}") + `

` + Primary("Usage:") + ` {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
{{if .Aliases}}` + Primary("Aliases:") + `  {{join .Aliases ", "}}{{end}}
{{if .Category}}` + Primary("Category:") + ` {{.Category}}{{end}}{{if .Description}}

` + Section("Description") + `
{{.Description}}{{end}}{{if .VisibleFlags}}

` + Section("Options") + `
   {{range .VisibleFlags}}` + Flag("{{.}}") + `
   {{end}}{{end}}
`

// AppHelp contains the template of dops help text.
var AppHelp = Primary("\nDOPS - CLI DevOps Toolkit") + `

{{if .VisibleFlags}}` + Section(`Global options`) + `
  ` + Flag(`{{range $index, $option := .VisibleFlags}}{{if $index}}`) + `
  ` + Flag(`{{end}}{{$option}}{{end}}{{end}}`) + `

{{if .VisibleCommands}}` + Section(`Modules`) + `{{range .VisibleCategories}}{{if .Name}}
  [` + Primary(`{{.Name}}`) + `]{{range .VisibleCommands}}
    · ` + Secondary(`{{join .Names ", "}}`) + Separator(`{{"\t|\t"}}`) + `{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
    · ` + Secondary(`{{join .Names ", "}}`) + Separator(`{{"\t|\t"}}`) + `{{.Usage}}{{end}}{{end}}{{end}}{{end}}

` + HiRedString("Contribute to this tool here: https://github.com/dops-cli ") + RedString("<3\n")

// SubcommandHelp is the text template for the subcommand help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var SubcommandHelp = `NAME:
   {{.HelpName}} - {{.Usage}}

USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} command{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}

COMMANDS:{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{range .VisibleCommands}}
     {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
   {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`

var MarkdownDoc = `% {{ .App.Name }} 8

# NAME

{{ .App.Name }}{{ if .App.Usage }} - {{ .App.Usage }}{{ end }}

# SYNOPSIS

{{ .App.Name }}
{{ if .SynopsisArgs }}
` + "```" + `
{{ range $v := .SynopsisArgs }}{{ $v }}{{ end }}` + "```" + `
{{ end }}{{ if .App.UsageText }}
# DESCRIPTION

{{ .App.UsageText }}
{{ end }}
**Usage**:

` + "```" + `
{{ .App.Name }} [GLOBAL OPTIONS] command [COMMAND OPTIONS] [ARGUMENTS...]
` + "```" + `
{{ if .GlobalArgs }}
# GLOBAL OPTIONS
{{ range $v := .GlobalArgs }}
{{ $v }}{{ end }}
{{ end }}{{ if .Commands }}
# COMMANDS
{{ range $v := .Commands }}
{{ $v }}{{ end }}{{ end }}`
