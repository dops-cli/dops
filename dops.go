package main

import (
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/module/modules"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/say/color"
	. "github.com/dops-cli/dops/screens"
	"github.com/dops-cli/dops/template"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/urfave/cli/v2"
	"io"
	"log"
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
			TviewApp = tview.NewApplication()
			TviewTable = tview.NewTable()

			TviewApp.EnableMouse(true)

			TviewTable.SetTitle("DOPS")

			for i, m := range module.ActiveModules {
				for j, command := range m.GetCommands() {
					TviewTable.SetCell(i+j, 0, tview.NewTableCell(command.Name))
					TviewTable.SetCell(i+j, 1, tview.NewTableCell(command.Usage))
				}
			}

			TviewTable.SetSelectable(true, false)
			TviewTable.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
				if key == tcell.KeyEscape {
					TviewApp.Stop()
				}
			}).SetSelectedFunc(func(row int, column int) {
				cell := TviewTable.GetCell(row, column)
				cmd := module.GetByName(cell.Text)
				ShowModule(TviewApp, cmd)
			})
			if err := TviewApp.SetRoot(TviewTable, true).SetFocus(TviewTable).Run(); err != nil {
				say.Error(err)
			}
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(module.CliApp.Flags))
	sort.Sort(cli.CommandsByName(module.CliApp.Commands))

	err := module.CliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
