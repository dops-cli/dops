package main

import (
	. "github.com/dops-cli/dops/interactive"
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/module/modules"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/say/color"
	"github.com/dops-cli/dops/template"
	"github.com/gdamore/tcell"
	"github.com/urfave/cli/v2"
	"gitlab.com/tslocum/cview"
	"io"
	"os"
	"sort"
)

var (
	CliFlags    []cli.Flag
	CliCommands []*cli.Command
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
		Version:  "v1.12.0", // <---VERSION---> This comment is used for CI, do NOT modify it!
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
		Action: func(context *cli.Context) error {
			CviewApp = cview.NewApplication()
			CviewTable = cview.NewTable()

			CviewApp.EnableMouse(true)
			CviewTable.SetTitle("DOPS")

			for i, m := range module.ActiveModules {
				for j, command := range m.GetCommands() {
					CviewTable.SetCell(i+j, 0, cview.NewTableCell(command.Name))
					CviewTable.SetCell(i+j, 1, cview.NewTableCell(command.Usage))
				}
			}

			CviewTable.SetSelectable(true, false)
			CviewTable.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
				if key == tcell.KeyEscape {
					CviewApp.Stop()
				}
			}).SetSelectedFunc(func(row int, column int) {
				cell := CviewTable.GetCell(row, column)
				cmd := module.GetByName(cell.Text)
				err := ShowModule(CviewApp, cmd)
				if err != nil {
					say.Error(err)
				}
			})
			if err := CviewApp.SetRoot(CviewTable, true).SetFocus(CviewTable).Run(); err != nil {
				say.Error(err)
			}
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(module.CliApp.Flags))
	sort.Sort(cli.CommandsByName(module.CliApp.Commands))

	err := module.CliApp.Run(os.Args)
	if err != nil {
		say.Error(err)
	}
}
