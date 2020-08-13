package template

import (
	. "github.com/dops-cli/dops/say/color"
)

var SectionColor = New(FgHiYellow, Underline)

// CommandHelp contains the template of dops modules help text.
var CommandHelp = HiCyanString("\n{{.HelpName}}") + ` - ` + HiMagentaString("{{.Usage}}") + `

` + HiCyanString("Usage:") + ` {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
{{if .Aliases}}` + HiCyanString("Aliases:") + `  {{join .Aliases ", "}}{{end}}
{{if .Category}}` + HiCyanString("Category:") + ` {{.Category}}{{end}}{{if .Description}}

` + SectionColor.Sprint("Description") + `
{{.Description}}{{end}}{{if .VisibleFlags}}

` + SectionColor.Sprint("Options") + `
   {{range .VisibleFlags}}` + YellowString("{{.}}") + `
   {{end}}{{end}}
`

// AppHelp contains the template of dops help text.
var AppHelp = HiCyanString("\nDOPS - CLI DevOps Toolkit") + `

{{if .VisibleFlags}}` + New(FgHiYellow, Underline).Sprint(`Global options`) + `
  ` + YellowString(`{{range $index, $option := .VisibleFlags}}{{if $index}}`) + `
  ` + YellowString(`{{end}}{{$option}}{{end}}{{end}}`) + `

{{if .VisibleCommands}}` + New(FgHiYellow, Underline).Sprint(`Modules`) + `{{range .VisibleCategories}}{{if .Name}}
  [` + HiCyanString(`{{.Name}}`) + `]{{range .VisibleCommands}}
    · ` + HiMagentaString(`{{join .Names ", "}}`) + HiRedString(`{{"\t|\t"}}`) + `{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
    · ` + HiMagentaString(`{{join .Names ", "}}`) + HiRedString(`{{"\t|\t"}}`) + `{{.Usage}}{{end}}{{end}}{{end}}{{end}}

` + HiRedString("Contribute to this tool here: https://github.com/dops-cli ") + RedString("<3\n")

// SubcommandHelpTemplate is the text template for the subcommand help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var SubcommandHelpTemplate = `NAME:
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

var MarkdownDocTemplate = `% {{ .App.Name }} 8

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
