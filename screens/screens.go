package screens

import (
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/say"
	"github.com/urfave/cli/v2"
	"gitlab.com/tslocum/cview"
	"strconv"
	"strings"
)

var CviewApp *cview.Application
var CviewTable *cview.Table

func ShowTable(app *cview.Application) {
	app.SetRoot(CviewTable, true)
}

func ShowEmpty(app *cview.Application) {
	app.SetRoot(nil, false)
}

func ShowModule(app *cview.Application, cmd *cli.Command) error {
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

	flags := make(map[string]string)

	fieldWidth := 10

	if len(cmd.Flags) > 0 {
		for _, flag := range cmd.Flags {

			bf, ok := flag.(*cli.BoolFlag)
			if ok {
				BoolFlags = append(BoolFlags, bf)
			}

			df, ok := flag.(*cli.DurationFlag)
			if ok {
				DurationFlags = append(DurationFlags, df)
			}

			ff, ok := flag.(*cli.Float64Flag)
			if ok {
				Float64Flags = append(Float64Flags, ff)
			}

			fsf, ok := flag.(*cli.Float64SliceFlag)
			if ok {
				Float64SliceFlags = append(Float64SliceFlags, fsf)
			}

			iflag, ok := flag.(*cli.IntFlag)
			if ok {
				IntFlags = append(IntFlags, iflag)
			}

			isf, ok := flag.(*cli.IntSliceFlag)
			if ok {
				IntSliceFlags = append(IntSliceFlags, isf)
			}

			pf, ok := flag.(*cli.PathFlag)
			if ok {
				PathFlags = append(PathFlags, pf)
			}

			sf, ok := flag.(*cli.StringFlag)
			if ok {
				StringFlags = append(StringFlags, sf)
			}

			ssf, ok := flag.(*cli.StringSliceFlag)
			if ok {
				StringSliceFlags = append(StringSliceFlags, ssf)
			}

			tf, ok := flag.(*cli.TimestampFlag)
			if ok {
				TimestampFlags = append(TimestampFlags, tf)
			}

		}

		for _, flag := range BoolFlags {
			f := *flag
			form.AddCheckBox(flag.Name, "Message", flag.Value, func(text bool) {
				flags[f.Name] = strconv.FormatBool(text)
			})
		}

		for _, flag := range DurationFlags {
			f := *flag
			form.AddInputField(flag.Name, flag.Value.String(), fieldWidth, nil, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range Float64Flags {
			f := *flag
			form.AddInputField(flag.Name, strconv.FormatFloat(flag.Value, 'G', -1, 64), fieldWidth, cview.InputFieldFloat, func(text string) {
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
			form.AddInputField(flag.Name, def, fieldWidth, nil, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range IntFlags {
			f := *flag
			form.AddInputField(flag.Name, strconv.Itoa(flag.Value), fieldWidth, cview.InputFieldInteger, func(text string) {
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
			form.AddInputField(flag.Name, def, fieldWidth, nil, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range PathFlags {
			f := *flag
			form.AddInputField(flag.Name, flag.Value, fieldWidth, nil, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range StringFlags {
			f := *flag
			form.AddInputField(flag.Name, flag.Value, fieldWidth, nil, func(text string) {
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
			form.AddInputField(flag.Name, def, fieldWidth, nil, func(text string) {
				flags[f.Name] = text
			})
		}

		for _, flag := range TimestampFlags {
			f := *flag

			var ret string
			if flag.Value != nil {
				ret = flag.Value.String()
			}

			form.AddInputField(flag.Name, ret, fieldWidth, nil, func(text string) {
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
			ShowTable(app)
		})

		form.SetWrapAround(true)

		form.SetBorder(true).SetTitle(" " + cmd.Name + " - " + cmd.Usage + " ").SetTitleAlign(cview.AlignLeft)
		CviewApp.SetRoot(form, true)

	} else {
		modal := cview.NewModal()

		modal.SetTitle("Confirm to run " + cmd.Name)
		modal.SetText("This module does not have any options. Run it?")
		modal.AddButtons([]string{"Run", "Cancel"})
		modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Cancel" {
				ShowTable(app)
			} else if buttonLabel == "Run" {
				app.Stop()
				err := module.Run(cmd, nil)
				if err != nil {
					say.Error(err)
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
