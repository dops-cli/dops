package debug

import (
	"github.com/urfave/cli/v2"
)

var IsDebug bool

type Module struct{}

func (Module) GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Value:       false,
			Destination: &IsDebug,
		},
	}
}

func (Module) GetCommands() []*cli.Command {
	return []*cli.Command{}
}
