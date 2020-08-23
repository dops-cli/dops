package interactive

import (
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
	"gitlab.com/tslocum/cview"

	. "github.com/dops-cli/dops/global"
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/say"
)

func ShowInteractiveModuleList(app *cview.Application) {
	app.SetRoot(CviewTable, true)
}

func ShowModule(app *cview.Application, cmd *cli.Command) error {
	fieldWidth := 0
	flags := make(map[string]string)
	form := cview.NewForm()

	var BoolFlags []*cli.BoolFlag
	var DurationFlags []*cli.DurationFlag
	var Float64Flags []*cli.Float64Flag
	var Float64SliceFlags []*cli.Float64SliceFlag
	var IntFlags []*cli.IntFlag
	var IntSliceFlags []*cli.IntSliceFlag
	var PathFlags []*cli.PathFlag
	var StringFlags []*cli.StringFlag
	var StringSliceFlags []*cli.StringSliceFlag
	var TimestampFlags []*cli.TimestampFlag

	form.SetWrapAround(true)
	form.SetBorder(true)
	form.SetTitle(" " + cmd.Name + " - " + cmd.Usage + " ")
	form.SetTitleAlign(cview.AlignLeft)

	if len(cmd.Flags) > 0 {
		for _, flag := range cmd.Flags {

			boolFlag, ok := flag.(*cli.BoolFlag)
			if ok {
				BoolFlags = append(BoolFlags, boolFlag)
			}

			durationFlag, ok := flag.(*cli.DurationFlag)
			if ok {
				DurationFlags = append(DurationFlags, durationFlag)
			}

			float64Flag, ok := flag.(*cli.Float64Flag)
			if ok {
				Float64Flags = append(Float64Flags, float64Flag)
			}

			float64SliceFlag, ok := flag.(*cli.Float64SliceFlag)
			if ok {
				Float64SliceFlags = append(Float64SliceFlags, float64SliceFlag)
			}

			intFlag, ok := flag.(*cli.IntFlag)
			if ok {
				IntFlags = append(IntFlags, intFlag)
			}

			intSliceFlag, ok := flag.(*cli.IntSliceFlag)
			if ok {
				IntSliceFlags = append(IntSliceFlags, intSliceFlag)
			}

			pathFlag, ok := flag.(*cli.PathFlag)
			if ok {
				PathFlags = append(PathFlags, pathFlag)
			}

			stringFlag, ok := flag.(*cli.StringFlag)
			if ok {
				StringFlags = append(StringFlags, stringFlag)
			}

			stringSliceFlag, ok := flag.(*cli.StringSliceFlag)
			if ok {
				StringSliceFlags = append(StringSliceFlags, stringSliceFlag)
			}

			timestampFlag, ok := flag.(*cli.TimestampFlag)
			if ok {
				TimestampFlags = append(TimestampFlags, timestampFlag)
			}
		}

		for _, flag := range BoolFlags {
			f := *flag
			form.AddCheckBox(flag.Name+" - "+flag.Usage, "", flag.Value, func(text bool) {
				flags[f.Name] = strconv.FormatBool(text)
			})
		}

		for _, flag := range DurationFlags {
			f := *flag
			form.AddInputField(flag.Name+" - "+flag.Usage, flag.Value.String(), fieldWidth, nil, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range Float64Flags {
			f := *flag
			form.AddInputField(flag.Name+" - "+flag.Usage, strconv.FormatFloat(flag.Value, 'G', -1, 64), fieldWidth, cview.InputFieldFloat, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range Float64SliceFlags {
			f := *flag

			var stringSlice []string
			if flag.Value != nil {
				for _, f := range flag.Value.Value() {
					stringSlice = append(stringSlice, strconv.FormatFloat(f, 'G', -1, 64))
				}
			}

			def := strings.Join(stringSlice, ", ")
			form.AddInputField(flag.Name+" - "+flag.Usage, def, fieldWidth, nil, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range IntFlags {
			f := *flag
			form.AddInputField(flag.Name+" - "+flag.Usage, strconv.Itoa(flag.Value), fieldWidth, cview.InputFieldInteger, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range IntSliceFlags {
			f := *flag

			stringSlice := []string{""}
			if flag.Value != nil {
				for _, i := range flag.Value.Value() {
					stringSlice = append(stringSlice, strconv.Itoa(i))
				}
			}

			def := strings.Join(stringSlice, ", ")
			form.AddInputField(flag.Name+" - "+flag.Usage, def, fieldWidth, nil, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range PathFlags {
			f := *flag
			form.AddInputField(flag.Name+" - "+flag.Usage, flag.Value, fieldWidth, nil, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range StringFlags {
			f := *flag
			form.AddInputField(flag.Name+" - "+flag.Usage, flag.Value, fieldWidth, nil, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range StringSliceFlags {
			f := *flag

			var stringSlice []string
			if flag.Value != nil {
				for _, s := range flag.Value.Value() {
					stringSlice = append(stringSlice, s)
				}
			}

			def := strings.Join(stringSlice, ", ")
			form.AddInputField(flag.Name+" - "+flag.Usage, def, fieldWidth, nil, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range TimestampFlags {
			f := *flag

			var ret string
			if flag.Value != nil {
				ret = flag.Value.String()
			}

			form.AddInputField(flag.Name+" - "+flag.Usage, ret, fieldWidth, nil, func(text string) {
				flags[f.Name] = text
			})
		}

		form.AddButton("Run", func() {
			app.Stop()
			err := module.Run(cmd, flags)
			if err != nil {
				panic(err)
			}
		})

		form.AddButton("Cancel", func() {
			ShowInteractiveModuleList(app)
		})

		flex := cview.NewFlex()
		flex.AddItem(cview.NewFlex().SetDirection(cview.FlexRow).
			AddItem(cview.NewTextView().SetText(cmd.Description), 0, 1, false).
			AddItem(form, 0, 4, true), 0, 1, true)

		CviewApp.SetRoot(flex, true)

	} else {
		modal := cview.NewModal()

		modal.SetTitle("Confirm to run " + cmd.Name)
		modal.SetText("This module does not have any options. Run it?")
		modal.AddButtons([]string{"Run", "Cancel"})
		modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Cancel" {
				ShowInteractiveModuleList(app)
			} else if buttonLabel == "Run" {
				app.Stop()
				err := module.Run(cmd, nil)
				if err != nil {
					say.Fatal(err)
				}
			}
		})

		err := CviewApp.SetRoot(modal, true).SetFocus(modal).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
