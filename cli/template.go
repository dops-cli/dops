package cli

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/pterm/pterm"

	"github.com/dops-cli/dops/global/options"
	"github.com/dops-cli/dops/theme"
	"github.com/dops-cli/dops/utils"
)

// AppHelpTemplate is the text template for the Default help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var AppHelpTemplate = theme.Primary("\nDOPS - CLI DevOps Toolkit") + `

{{if .VisibleFlags}}` + theme.Section(`Global options`) + `
  ` + theme.Flag(`{{range $index, $option := .VisibleFlags}}{{if $index}}`) + `
  ` + theme.Flag(`{{end}}{{$option}}{{end}}{{end}}`) + `

{{if .VisibleCommands}}` + theme.Section(`Modules`) + `{{range .VisibleCategories}}{{if .Name}}
  [` + theme.Primary(`{{.Name}}`) + `]{{range .VisibleCommands}}
    · ` + theme.Secondary(`{{join .Names ", "}}`) + theme.Separator(`{{"\t|\t"}}`) + `{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
    · ` + theme.Secondary(`{{join .Names ", "}}`) + theme.Separator(`{{"\t|\t"}}`) + `{{.Usage}}{{end}}{{end}}{{end}}{{end}}

` + pterm.LightRed("Contribute to this tool here: https://github.com/dops-cli ") + pterm.Red("<3\n")

// CommandHelpTemplate is the text template for the command help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var CommandHelpTemplate = theme.Primary("\n{{.Name}}") + ` - ` + theme.Secondary("{{.Usage}}") + pterm.Normal() + `

{{if .Description}}` + theme.Section("Description") + `
{{.Description}}{{end}}

` + theme.Primary("Usage:") + ` {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
{{if .Aliases}}` + theme.Primary("Aliases:") + `  {{join .Aliases ", "}}{{end}}
{{if .Category}}` + theme.Primary("Category:") + ` {{.Category}}{{end}}

{{if .VisibleFlags}}` + theme.Section("Options") + `
   {{range .VisibleFlags}}` + theme.Flag("{{.}}") + `
   {{end}}{{end}}
`

// SubcommandHelpTemplate is the text template for the subcommand help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var SubcommandHelpTemplate = theme.Primary("\n{{.Name}}") + ` - ` + theme.Secondary("{{.Usage}}") + pterm.Normal() + `

{{if .Description}}` + theme.Section("Description") + `
{{.Description}}{{end}}

` + theme.Primary("Usage:") + ` {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} command{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}

{{if .Aliases}}` + theme.Primary("Aliases:") + `  {{join .Aliases ", "}}{{end}}
{{if .Category}}` + theme.Primary("Category:") + ` {{.Category}}{{end}}

{{if .VisibleCommands}}` + theme.Section(`Commands`) + `{{range .VisibleCategories}}{{if .Name}}
  [` + theme.Primary(`{{.Name}}`) + `]{{range .VisibleCommands}}
    · ` + theme.Secondary(`{{join .Names ", "}}`) + theme.Separator(`{{"\t|\t"}}`) + `{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
    · ` + theme.Secondary(`{{join .Names ", "}}`) + theme.Separator(`{{"\t|\t"}}`) + `{{.Usage}}{{end}}{{end}}{{end}}{{end}}

{{if .VisibleFlags}}` + theme.Section("Options") + `
   {{range .VisibleFlags}}` + theme.Flag("{{.}}") + `
   {{end}}{{end}}
`

// MarkdownDocTemplate is the template used for markdown documentation
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

// FishCompletionTemplate is the template, which resolves to fish autocompletion
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

var funcMap = template.FuncMap{"join": strings.Join}

// Modules is a wrapper for cli.Commands
type Modules struct {
	Commands Commands
}

// PrintModules prints all modules to stdout
func PrintModules() error {

	var modules = `{{range .Commands}}` +
		theme.Primary("\n{{.Name}}") + ` - ` + theme.Secondary("{{.Usage}}") + `

  ` + theme.Primary("Usage:") + ` {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
  {{if .Aliases}}` + theme.Primary("Aliases:") + `  {{join .Aliases ", "}}{{end}}
  {{if .Category}}` + theme.Primary("Category:") + ` {{.Category}}{{end}}{{if .Description}}

` + theme.Section("Description") + `
{{.Description}}{{end}}{{if .VisibleFlags}}

` + theme.Section("Options") + `
  {{range .VisibleFlags}}` + theme.Flag("{{.}}") + `
  {{end}}{{end}}` + "\n\n" + `{{end}}`

	var commands []*Command

	for _, m := range ActiveModules {
		commands = append(commands, m.GetModuleCommands()...)
	}

	t := template.Must(template.New("modules").Funcs(funcMap).Parse(modules))

	err := t.Execute(os.Stdout, Modules{commands})
	if err != nil {
		return err
	}

	return nil
}

// CommandDocumentation returns the documentation used at https://dops-cli.com for a module
func CommandDocumentation(cmd *Command, parent *Command, level int) string {

	var docs string

	var levelPrefix string

	for i := 0; i < level; i++ {
		levelPrefix += "#"
	}

	docs += levelPrefix + "# " + cmd.Name + "\n\n"

	docs += "> " + cmd.Usage + "\n\n"

	docs += cmd.Description + "\n\n"
	if cmd.Warning != "" {
		docs += "!> **WARNING**\n"
		docs += cmd.Warning + "  \n\n"
	}
	if cmd.Tip != "" {
		docs += "?> **TIP**\n"
		docs += cmd.Tip + "  \n\n"
	}
	if cmd.Note != "" {
		docs += "?> **NOTE**\n"
		docs += cmd.Note + "  \n\n"
	}

	docs += levelPrefix + "## Usage\n\n"
	docs += "> `dops [options] "
	if parent != nil {
		docs += parent.Name + " "
		if len(parent.Flags) > 0 {
			docs += "[options] "
		}
	}
	docs += cmd.Name + " "
	if cmd.UsageText != "" {
		docs += cmd.UsageText + " "
	} else if cmd.HelpName != "" {
		docs += cmd.HelpName + " "
	}
	if len(cmd.VisibleFlags()) > 0 {
		docs += "[options] "
	}
	if len(cmd.Subcommands) > 0 {
		docs += "subcommand "
	}
	if cmd.ArgsUsage != "" {
		docs += cmd.ArgsUsage + " "
	} else {
		docs += "[arguments...]"
	}

	docs += "`\n\n"

	docs += "**Category:** " + cmd.Category + "  \n"
	if len(cmd.Aliases) > 0 {
		docs += "**Aliases:** `" + strings.Join(cmd.Aliases, ", ") + "`  \n"
	}
	if len(cmd.Flags) > 0 {
		docs += "\n" + levelPrefix + "### Options\n"
		docs += "```flags\n"
		for _, flag := range cmd.Flags {
			docs += flag.String() + "  \n"
		}
		docs += "```\n"
	}
	if len(cmd.Subcommands) > 0 {
		docs += levelPrefix + "## Submodules\n\n"
		for _, scmd := range cmd.Subcommands {
			var jump int
			if level == 0 {
				jump = 1
			}
			docs += CommandDocumentation(scmd, cmd, level+jump+1)
		}
	}

	if level == 0 {
		docs += "## Examples\n\n"
		docs += generateExamples(cmd)
	}

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 1, 8, 2, ' ', 0)

	_, err := fmt.Fprint(w, docs)
	if err != nil {
		log.Fatal(err)
	}

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func generateExamples(cmd *Command) string {
	var docs string

	if len(cmd.Examples) > 0 {
		for _, example := range cmd.Examples {
			docs += "### " + example.ShortDescription + "\n\n"
			docs += "```command\n"
			docs += "" + example.Usage + "\n"
			docs += "```\n"
			if example.AsciinemaID != "" {
				docs += `<a id="asciicast-` + example.AsciinemaID + `" data-autoplay="true" data-loop="true"></a>` + "\n"
			}
			if example.GenerateSVG {
				docs += "<img src=\"" + generateSVG(example.Usage) + "\">\n"
			}
			docs += "\n"
		}
	}

	_ = os.MkdirAll("./example_casts", 0600)

	for _, subcmd := range cmd.Subcommands {
		docs += generateExamples(subcmd)
	}

	_ = os.RemoveAll("./example_casts")

	return docs
}

func generateSVG(command string) string {
	castFile := "./example_casts/" + generateCastFile(command)
	svgFile := "./docs/_assets/example_svg/" + generateFileName(command) + ".svg"

	args := []string{"-c", "svg-term --in " + castFile + ".json --out " + svgFile}

	cmd := exec.Command("bash", args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	cmd.Run()

	svgFile = strings.ReplaceAll(svgFile, "./docs", "")

	return svgFile

}

func generateCastFile(command string) string {

	filename := generateFileName(command)

	command = strings.Replace(command, "dops", "go run .", 1)

	args := []string{"-c", "asciinema rec ./example_casts/" + filename + ".json -c '" + command + "'"}

	cmd := exec.Command("bash", args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	cmd.Run()

	if options.CI || options.Debug {
		fmt.Printf("\n\nOutput:\n%v\n\nErrors:\n%v\n", out.String(), stderr.String())
	}

	utils.WriteFile("./example_casts/"+filename+".json", []byte("[5, \"o\", \"\\r\\nrestarting...\\r\\n\"]"), true)

	return filename
}

func generateFileName(usage string) string {
	r := regexp.MustCompile(`(?m)\W`)
	return r.ReplaceAllString(usage, "")
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

	var commands []*Command

	for _, m := range ActiveModules {
		commands = append(commands, m.GetModuleCommands()...)
	}

	t := template.Must(template.New("modules").Funcs(funcMap).Parse(modules))

	err := t.Execute(os.Stdout, Modules{commands})
	if err != nil {
		return err
	}

	return nil
}
