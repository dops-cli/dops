package main

import (
	"io"
	"os"
	"sort"

	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/global"
	"github.com/dops-cli/dops/interactive"
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/module/modules"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/say/color"
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
		Version:              "v1.22.0", // <---VERSION---> This comment is used for CI, do NOT modify it!
		Commands:             global.CliCommands,
		Flags:                global.CliFlags,
		EnableBashCompletion: true,
		Action: func(ctx *cli.Context) error {

			err := interactive.Start()
			if err != nil {
				return err
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
