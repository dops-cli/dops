package interactive

import (
	"os"
	"os/signal"
	"regexp"
	"strings"
	"text/scanner"

	"github.com/c-bata/go-prompt"

	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/module"
)

var (
	modules []prompt.Suggest

	cliCommands []*cli.Command
)

// Start starts the interactive mode of dops
func Start() error {
	for _, m := range cli.ActiveModules {
		for _, command := range m.GetModuleCommands() {
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

		prompt.OptionAddKeyBind(struct {
			Key prompt.Key
			Fn  prompt.KeyBindFunc
		}{Key: prompt.Escape, Fn: func(buffer *prompt.Buffer) {
			os.Exit(0)
		}}),

		prompt.OptionTitle("DOPS"),
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			os.Exit(0)
		}
	}()

	t := prompt.Input(">>> dops ", Completer, options...)
	args := []string{"dops"}
	args = append(args, splitCommand(t)...)
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

func splitCommand(text string) []string {
	var s scanner.Scanner
	s.Init(strings.NewReader(text))
	slice := make([]string, 0, 5)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		slice = append(slice, s.TokenText())
	}
	return slice
}
