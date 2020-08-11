package main

import (
	"github.com/dops-cli/dops/constants"
	"github.com/dops-cli/dops/module"
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
}

func main() {

	for _, f := range module.RegisteredGlobalFlags {
		CliFlags = append(CliFlags, f.GetFlags()...)
	}

	for _, m := range module.RegisteredModules {
		CliCommands = append(CliCommands, m.GetCommands()...)
	}

	app := &cli.App{
		Name:    "dops",
		Version: "v1.0.0",
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
