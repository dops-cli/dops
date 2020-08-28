package interactive

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
	"gitlab.com/tslocum/cview"

	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/global"
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/say/color"
)

// ShowInteractiveModuleList shows the interactive table of all modules to the user
func ShowInteractiveModuleList(app *cview.Application) {
	app.SetRoot(global.CviewTable, true)
}

// ShowModule shows the form dialoge for a module to the user
func ShowModule(app *cview.Application, cmd *cli.Command, data ...interface{}) error {
	const fieldWidth = 0

	cli.SetupFlags(cmd)

	var flags []string
	flagMap := make(map[string]string)

	if data != nil {
		v, ok := data[0].([]string)
		if ok {
			flags = v
		}
	}

	flags = append(flags, cmd.Name)

	var subCommand *cli.Command

	if len(cmd.Flags) > 0 || len(cmd.Subcommands) > 0 {
		form := cview.NewForm()

		if len(cmd.Subcommands) > 0 {
			var subcommands []string

			for _, subcommand := range cmd.Subcommands {
				subcommands = append(subcommands, subcommand.Name)
			}

			form.AddDropDown("Command", subcommands, 0, func(option string, _ int) {
				for _, subcmd := range cmd.Subcommands {
					if subcmd.Name == option {
						subCommand = subcmd
					}
				}
			})
		}

		for _, flag := range cli.BoolFlags {
			f := *flag
			form.AddCheckBox(flag.Name+" - "+flag.Usage, "", flag.Value, func(text bool) {
				flagMap[f.Name] = strconv.FormatBool(text)
			})
		}

		for _, flag := range cli.DurationFlags {
			f := *flag
			form.AddInputField(flag.Name+" - "+flag.Usage, flag.Value.String(), fieldWidth, nil, func(text string) {
				flagMap[f.Name] = text
			})
		}

		for _, flag := range cli.Float64Flags {
			f := *flag
			form.AddInputField(flag.Name+" - "+flag.Usage, strconv.FormatFloat(flag.Value, 'G', -1, 64), fieldWidth, cview.InputFieldFloat, func(text string) {
				flagMap[f.Name] = text
			})
		}

		for _, flag := range cli.Float64SliceFlags {
			f := *flag

			var stringSlice []string
			if flag.Value != nil {
				for _, f := range flag.Value.Value() {
					stringSlice = append(stringSlice, strconv.FormatFloat(f, 'G', -1, 64))
				}
			}

			def := strings.Join(stringSlice, ", ")
			form.AddInputField(flag.Name+" - "+flag.Usage, def, fieldWidth, nil, func(text string) {
				flagMap[f.Name] = text
			})
		}

		for _, flag := range cli.IntFlags {
			f := *flag
			form.AddInputField(flag.Name+" - "+flag.Usage, strconv.Itoa(flag.Value), fieldWidth, cview.InputFieldInteger, func(text string) {
				flagMap[f.Name] = text
			})
		}

		for _, flag := range cli.IntSliceFlags {
			f := *flag

			stringSlice := []string{""}
			if flag.Value != nil {
				for _, i := range flag.Value.Value() {
					stringSlice = append(stringSlice, strconv.Itoa(i))
				}
			}

			def := strings.Join(stringSlice, ", ")
			form.AddInputField(flag.Name+" - "+flag.Usage, def, fieldWidth, nil, func(text string) {
				flagMap[f.Name] = text
			})
		}

		for _, flag := range cli.PathFlags {
			f := *flag
			form.AddInputField(flag.Name+" - "+flag.Usage, flag.Value, fieldWidth, nil, func(text string) {
				flagMap[f.Name] = text
			})
		}

		for _, flag := range cli.StringFlags {
			f := *flag
			form.AddInputField(flag.Name+" - "+flag.Usage, flag.Value, fieldWidth, nil, func(text string) {
				flagMap[f.Name] = text
			})
		}

		for _, flag := range cli.StringSliceFlags {
			f := *flag

			var stringSlice []string
			if flag.Value != nil {
				stringSlice = append(stringSlice, flag.Value.Value()...)
			}

			def := strings.Join(stringSlice, ", ")
			form.AddInputField(flag.Name+" - "+flag.Usage, def, fieldWidth, nil, func(text string) {
				flagMap[f.Name] = text
			})
		}

		for _, flag := range cli.TimestampFlags {
			f := *flag

			var ret string
			if flag.Value != nil {
				ret = flag.Value.String()
			}

			form.AddInputField(flag.Name+" - "+flag.Usage, ret, fieldWidth, nil, func(text string) {
				flagMap[f.Name] = text
			})
		}

		for _, flag := range cli.OptionFlags {
			f := *flag
			form.AddDropDown(flag.Name+" - "+flag.Usage, flag.Options, 0, func(text string, _ int) {
				flagMap[f.Name] = text
			})
		}

		form.SetWrapAround(true)
		form.SetBorder(true)
		form.SetTitle(" " + cmd.Name + " - " + cmd.Usage + " ")
		form.SetTitleAlign(cview.AlignLeft)

		form.SetTitleColor(color.TitleColor)
		form.SetButtonBackgroundColor(color.ButtonBackgroundColor)
		form.SetFieldBackgroundColor(color.FieldBackgroundColor)

		if subCommand == nil {
			form.AddButton("Run", func() {
				for f, v := range flagMap {
					flags = append(flags, flagToString(f, v))
				}
				global.CviewApp.Stop()
				err := module.Run(flags)
				if err != nil {
					say.Fatal(err)
				}
			})
		} else {
			form.AddButton("Next", func() {
				for f, v := range flagMap {
					flags = append(flags, flagToString(f, v))
				}
				err := ShowModule(global.CviewApp, subCommand, flags)
				if err != nil {
					say.Fatal(err)
				}
			})
		}

		form.AddButton("Cancel", func() {
			ShowInteractiveModuleList(app)
		})

		flex := cview.NewFlex()
		flex.AddItem(cview.NewFlex().SetDirection(cview.FlexRow).
			AddItem(cview.NewTextView().SetText(cmd.Description).SetTextColor(tcell.ColorGreenYellow), 0, 1, false).
			AddItem(form, 0, 4, true), 0, 1, true)

		global.CviewApp.SetRoot(flex, true)
	} else if len(cmd.Flags) == 0 && len(cmd.Subcommands) == 0 {
		modal := cview.NewModal()

		modal.SetButtonBackgroundColor(color.ButtonBackgroundColor)
		modal.SetTitleColor(color.TitleColor)
		modal.SetBackgroundColor(color.ModalBackgroundColor)

		modal.SetTitle("Confirm to run " + cmd.Name)
		modal.SetText("This module does not have any options. Run it?")
		modal.AddButtons([]string{"Run", "Cancel"})
		modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Cancel" {
				ShowInteractiveModuleList(app)
			} else if buttonLabel == "Run" {
				app.Stop()
				err := module.Run(flags)
				if err != nil {
					say.Fatal(err)
				}
			}
		})

		err := global.CviewApp.SetRoot(modal, true).SetFocus(modal).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func flagToString(flag string, value string) string {
	return "-" + flag + "=" + value
}
