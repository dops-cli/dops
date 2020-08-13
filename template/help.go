package template

import (
	. "github.com/dops-cli/dops/say/color"
)

// CommandHelp contains the template of dops modules help text.
var CommandHelp = `NAME:
   {{.HelpName}} - {{.Usage}}

USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}

CATEGORY:
   {{.Category}}{{end}}{{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{if .VisibleFlags}}

OPTIONS:
   {{range .VisibleFlags}}{{.}}
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
