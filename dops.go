package main

import (
	"io"
	"os"
	"sort"
	"strings"

	"github.com/gdamore/tcell"
	"gitlab.com/tslocum/cview"

	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/global"
	"github.com/dops-cli/dops/interactive"
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/module/modules"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/say/color"
	"github.com/dops-cli/dops/utils"
)

func init() {
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		cli.HelpPrinterCustom(color.Output, templ, data, nil)
	}
	cli.VersionPrinter = func(c *cli.Context) {
		say.Info("dops is currently on version " + color.Primary(c.App.Version) + "!")
	}
}

func main() {

	for _, f := range cli.ActiveGlobalFlags {
		global.CliFlags = append(global.CliFlags, f.GetFlags()...)
	}

	for _, m := range cli.ActiveModules {
		global.CliCommands = append(global.CliCommands, m.GetModuleCommands()...)
	}

	global.CliCommands = append(global.CliCommands, modules.Module{}.GetModuleCommands()...)

	module.CliApp = &cli.App{
		Name:                 "dops",
		HelpName:             "dops",
		Usage:                "CLI DevOps Toolkit",
		Version:              "v1.19.0", // <---VERSION---> This comment is used for CI, do NOT modify it!
		Commands:             global.CliCommands,
		Flags:                global.CliFlags,
		EnableBashCompletion: true,
		Action: func(ctx *cli.Context) error {
			global.CviewApp = cview.NewApplication()
			global.CviewTable = cview.NewTable()

			global.CviewApp.EnableMouse(true)

			global.CviewTable.SetTitle("DOPS")
			global.CviewTable.SetSelectable(true, false)
			global.CviewTable.SetScrollBarVisibility(cview.ScrollBarAuto)

			var categories []string

			for _, command := range global.CliCommands {
				if !utils.SliceContainsString(categories, command.Category) {
					categories = append(categories, command.Category)
				}
			}

			sort.Strings(categories)

			currentRow := 0

			for _, category := range categories {
				categoryCell := cview.NewTableCell(" --- " + category + " --- ")
				categoryCell.Color = tcell.Color87
				global.CviewTable.SetCell(currentRow, 0, categoryCell)
				currentRow++
				for _, command := range global.CliCommands {
					if command.Category == category {
						global.CviewTable.SetCell(currentRow, 0, cview.NewTableCell(command.Name))
						global.CviewTable.SetCell(currentRow, 1, cview.NewTableCell(command.Usage))
						currentRow++
					}
				}
			}

			global.CviewTable.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
				if key == tcell.KeyEscape {
					global.CviewApp.Stop()
				}
			}).SetSelectedFunc(func(row int, column int) {
				cell := global.CviewTable.GetCell(row, column)
				if strings.Contains(cell.Text, " --- ") {
					return
				}
				cmd, err := module.GetByName(cell.Text)
				if err != nil {
					say.Fatal(err)
				}
				err = interactive.ShowModule(global.CviewApp, cmd)
				if err != nil {
					say.Fatal(err)
				}
			})
			if err := global.CviewApp.SetRoot(global.CviewTable, true).SetFocus(global.CviewTable).Run(); err != nil {
				say.Fatal(err)
			}
			return nil
		},
		Authors: []*cli.Author{
			{
				Name:  "Marvin Wendt",
				Email: "dops@marvinjwendt.com",
			},
		},
		Copyright:              "(c) 2020 Marvin Wendt",
		Writer:                 color.Output,
		UseShortOptionHandling: true,
	}

	sort.Sort(cli.FlagsByName(module.CliApp.Flags))
	sort.Sort(cli.CommandsByName(module.CliApp.Commands))

	err := module.CliApp.Run(os.Args)
	if err != nil {
		say.Fatal(err)
	}
}
