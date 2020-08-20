package main

import (
	"io"
	"os"
	"sort"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/urfave/cli/v2"
	"gitlab.com/tslocum/cview"

	. "github.com/dops-cli/dops/global"
	. "github.com/dops-cli/dops/interactive"
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/module/modules"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/say/color"
	"github.com/dops-cli/dops/template"
	"github.com/dops-cli/dops/utils"
)

func init() {
	cli.AppHelpTemplate = template.AppHelp
	cli.CommandHelpTemplate = template.CommandHelp
	cli.SubcommandHelpTemplate = template.SubcommandHelp
	cli.MarkdownDocTemplate = template.MarkdownDoc
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		cli.HelpPrinterCustom(color.Output, templ, data, nil)
	}
	cli.VersionPrinter = func(c *cli.Context) {
		say.Info("dops is currently on version " + color.Primary(c.App.Version) + "!")
	}
}

func main() {

	for _, f := range module.ActiveGlobalFlags {
		CliFlags = append(CliFlags, f.GetFlags()...)
	}

	for _, m := range module.ActiveModules {
		CliCommands = append(CliCommands, m.GetCommands()...)
	}

	CliCommands = append(CliCommands, modules.Module{}.GetCommands()...)

	module.CliApp = &cli.App{
		Name:     "dops",
		HelpName: "dops",
		Usage:    "CLI DevOps Toolkit",
		Version:  "v1.15.1", // <---VERSION---> This comment is used for CI, do NOT modify it!
		Flags:    CliFlags,
		Commands: CliCommands,
		Authors: []*cli.Author{
			{
				Name:  "Marvin Wendt",
				Email: "dops@marvinjwendt.com",
			},
		},
		Copyright:              "(c) 2020 Marvin Wendt",
		Writer:                 color.Output,
		UseShortOptionHandling: true,
		Action: func(ctx *cli.Context) error {
			CviewApp = cview.NewApplication()
			CviewTable = cview.NewTable()

			CviewApp.EnableMouse(true)

			CviewTable.SetTitle("DOPS")
			CviewTable.SetSelectable(true, false)
			CviewTable.SetScrollBarVisibility(cview.ScrollBarAuto)

			var categories []string

			for _, command := range CliCommands {
				if !utils.ContainsString(categories, command.Category) {
					categories = append(categories, command.Category)
				}
			}

			sort.Strings(categories)

			currentRow := 0

			for _, category := range categories {
				categoryCell := cview.NewTableCell(" --- " + category + " --- ")
				categoryCell.Color = tcell.Color87
				CviewTable.SetCell(currentRow, 0, categoryCell)
				currentRow++
				for _, command := range CliCommands {
					if command.Category == category {
						CviewTable.SetCell(currentRow, 0, cview.NewTableCell(command.Name))
						CviewTable.SetCell(currentRow, 1, cview.NewTableCell(command.Usage))
						currentRow++
					}
				}
			}

			CviewTable.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
				if key == tcell.KeyEscape {
					CviewApp.Stop()
				}
			}).SetSelectedFunc(func(row int, column int) {
				cell := CviewTable.GetCell(row, column)
				if strings.Contains(cell.Text, " --- ") {
					return
				}
				cmd, err := module.GetByName(cell.Text)
				if err != nil {
					say.Fatal(err)
				}
				err = ShowModule(CviewApp, cmd)
				if err != nil {
					say.Fatal(err)
				}
			})
			if err := CviewApp.SetRoot(CviewTable, true).SetFocus(CviewTable).Run(); err != nil {
				say.Fatal(err)
			}
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(module.CliApp.Flags))
	sort.Sort(cli.CommandsByName(module.CliApp.Commands))

	err := module.CliApp.Run(os.Args)
	if err != nil {
		say.Fatal(err)
	}
}
