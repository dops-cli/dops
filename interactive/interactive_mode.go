package interactive

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/c-bata/go-prompt"

	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/say"
)

var (
	modules []prompt.Suggest

	cliCommands []*cli.Command
)

// Start starts the interactive mode of dops
func Start() error {
	for _, module := range cli.ActiveModules {
		for _, command := range module.GetModuleCommands() {
			modules = append(modules, prompt.Suggest{Text: command.Name, Description: command.Usage})
			cliCommands = append(cliCommands, command)
		}
	}

	options := []prompt.Option{
		prompt.OptionSuggestionBGColor(prompt.Turquoise),
		prompt.OptionSuggestionTextColor(prompt.DarkGray),

		prompt.OptionDescriptionBGColor(prompt.DefaultColor),
		prompt.OptionDescriptionTextColor(prompt.White),

		prompt.OptionSelectedDescriptionBGColor(prompt.DefaultColor),
		prompt.OptionSelectedDescriptionTextColor(prompt.Green),

		prompt.OptionTitle("DOPS"),
	}

	t := prompt.Input(">>> dops ", Completer, options...)
	args := []string{"dops"}
	args = append(args, strings.Split(t, " ")...)
	say.Info("Running: dops ", t)
	err := module.CliApp.Run(args)
	if err != nil {
		return err
	}

	return nil
}

// Completer handles the suggestions for autocomplete in interactive mode
func Completer(d prompt.Document) []prompt.Suggest {
	beforeCursor := d.TextBeforeCursor()

	if !strings.Contains(beforeCursor, " ") {
		return prompt.FilterContains(modules, d.GetWordBeforeCursor(), true)
	}

	args := strings.Split(beforeCursor, " ")
	moduleName := args[0]

	for _, command := range cliCommands {
		if command.Name == moduleName {
			lastCommand := FindLastSubCommand(command, beforeCursor)
			if lastCommand != nil {
				command = lastCommand
			}
			return suggestionsForCommand(d, d.Text, d.TextBeforeCursor(), *command)
		}
	}

	return nil
}

func suggestionsForCommand(d prompt.Document, currentText string, beforeCursorText string, command cli.Command) []prompt.Suggest {
	var flags []prompt.Suggest
	args := strings.Split(strings.TrimSpace(currentText), " ")
	var lastArgIsSubcommand bool

	lastSubcommand := FindLastSubCommand(&command, beforeCursorText)
	if lastSubcommand != nil {
		if args[len(args)-1] == lastSubcommand.Name {
			lastArgIsSubcommand = true
		}
	}

	if !lastArgIsSubcommand {
		for _, flag := range command.Flags {
			regexFlag := regexp.MustCompile(`(?m)--` + flag.Names()[0] + `\b`)

			if !regexFlag.MatchString(currentText) {
				flags = append(flags, prompt.Suggest{
					Text:        "--" + flag.Names()[0],
					Description: strings.TrimSpace(strings.Split(flag.String(), "|")[1]),
				})
			}
		}

	}

	regexSubcommandIsBeforeCursor := regexp.MustCompile(`(?m)\b` + args[len(args)-1] + `\b`)
	if lastSubcommand != nil && regexSubcommandIsBeforeCursor.MatchString(beforeCursorText) {

		newBeforeCursor := beforeCursorText
		newBeforeCursor = strings.TrimSpace(lastSubcommand.Name + " " + strings.TrimSpace(strings.Split(beforeCursorText, lastSubcommand.Name)[1]))
		newCurrentText := strings.TrimSpace(lastSubcommand.Name + " " + strings.TrimSpace(strings.Split(currentText, lastSubcommand.Name)[1]))
		fmt.Println("")
		fmt.Println(newBeforeCursor)
		fmt.Println(newCurrentText)
		fmt.Println("Name", command.Name)
		fmt.Println("")
		return suggestionsForCommand(d, newCurrentText, newBeforeCursor, *lastSubcommand)
	}

	for _, subcommand := range command.Subcommands {
		regexSubcommand := regexp.MustCompile(`(?m)\b` + subcommand.Name + `\b`)

		if !regexSubcommand.MatchString(currentText) {
			flags = append(flags, prompt.Suggest{
				Text:        subcommand.Names()[0],
				Description: strings.TrimSpace(subcommand.Usage),
			})
		}
	}

	return prompt.FilterContains(flags, d.GetWordBeforeCursor(), true)
}

// FindLastSubCommand finds the last subcommand in a string
func FindLastSubCommand(parent *cli.Command, text string) *cli.Command {
	var cmd *cli.Command
	args := strings.Split(text, " ")
	for _, command := range parent.Subcommands {
		for _, arg := range args {
			if command.Name == arg {
				cmd = command
			}
		}
	}

	return cmd
}
