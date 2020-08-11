package main

import (
	"github.com/dops-cli/dops/modules"
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
	cli.AppHelpTemplate += "\nContribute to this tool here: https://github.com/dops-cli <3\n"
}

func main() {

	for _, flag := range modules.RegisteredGlobalFlags {
		CliFlags = append(CliFlags, flag.GetFlags()...)
	}

	for _, module := range modules.RegisteredModules {
		CliCommands = append(CliCommands, module.GetCommands()...)
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
