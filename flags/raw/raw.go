package raw

import (
	"github.com/urfave/cli/v2"
)

var OutputRaw bool

type Flag struct{}

func (Flag) GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "raw",
			Usage:       "print minimal unstyled text",
			Value:       false,
			Destination: &OutputRaw,
		},
	}
}
