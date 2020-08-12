package main

import (
	"github.com/dops-cli/dops/constants"
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/module/modules"
	"github.com/dops-cli/dops/say"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

var (
	CliFlags    []cli.Flag
	CliCommands []*cli.Command
)

func init() {
	cli.AppHelpTemplate = constants.AppHelpTemplate
	cli.VersionPrinter = func(c *cli.Context) {
		say.Info("dops is currently on version " + c.App.Version + "!")
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

	app := &cli.App{
		Name:    "dops",
		Version: "v1.3.0",
		Authors: []*cli.Author{
			{
				Name:  "Marvin Wendt",
				Email: "dops@marvinjwendt.com",
			},
		},
		Copyright: "(c) 2020 Marvin Wendt",
		HelpName:  "dops",
		Usage:     "DevOps",
		Flags:     CliFlags,
		Commands:  CliCommands,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
