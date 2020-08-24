package cli

import (
	"github.com/dops-cli/dops/say/color"
)

// AppHelpTemplate is the text template for the Default help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var AppHelpTemplate = color.Primary("\nDOPS - CLI DevOps Toolkit") + `

{{if .VisibleFlags}}` + color.Section(`Global options`) + `
  ` + color.Flag(`{{range $index, $option := .VisibleFlags}}{{if $index}}`) + `
  ` + color.Flag(`{{end}}{{$option}}{{end}}{{end}}`) + `

{{if .VisibleCommands}}` + color.Section(`Modules`) + `{{range .VisibleCategories}}{{if .Name}}
  [` + color.Primary(`{{.Name}}`) + `]{{range .VisibleCommands}}
    · ` + color.Secondary(`{{join .Names ", "}}`) + color.Separator(`{{"\t|\t"}}`) + `{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
    · ` + color.Secondary(`{{join .Names ", "}}`) + color.Separator(`{{"\t|\t"}}`) + `{{.Usage}}{{end}}{{end}}{{end}}{{end}}

` + color.HiRedString("Contribute to this tool here: https://github.com/dops-cli ") + color.RedString("<3\n")

// CommandHelpTemplate is the text template for the command help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var CommandHelpTemplate = color.Primary("\n{{.Name}}") + ` - ` + color.Secondary("{{.Usage}}") + color.R + `

{{if .Description}}` + color.Section("Description") + `
{{.Description}}{{end}}

` + color.Primary("Usage:") + ` {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
{{if .Aliases}}` + color.Primary("Aliases:") + `  {{join .Aliases ", "}}{{end}}
{{if .Category}}` + color.Primary("Category:") + ` {{.Category}}{{end}}

{{if .VisibleFlags}}` + color.Section("Options") + `
   {{range .VisibleFlags}}` + color.Flag("{{.}}") + `
   {{end}}{{end}}
`

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

var FishCompletionTemplate = `# {{ .App.Name }} fish shell completion

function __fish_{{ .App.Name }}_no_subcommand --description 'Test if there has been any subcommand yet'
    for i in (commandline -opc)
        if contains -- $i{{ range $v := .AllCommands }} {{ $v }}{{ end }}
            return 1
        end
    end
    return 0
end

{{ range $v := .Completions }}{{ $v }}
{{ end }}`
