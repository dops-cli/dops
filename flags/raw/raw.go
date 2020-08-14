package raw

import (
	"github.com/urfave/cli/v2"
)

// OutputRaw is true if dops was started with the global raw flag.
// If OutputRaw is true, dops outputs an unformatted text.
var OutputRaw bool

type Flag struct{}

func (Flag) GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "raw",
			Aliases:     []string{"r"},
			Usage:       "print minimal unstyled text",
			Destination: &OutputRaw,
		},
	}
}
