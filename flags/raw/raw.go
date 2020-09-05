package raw

import (
	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/global/options"
)

// Flag returns the created flag
type Flag struct{}

// GetFlags returns the global flags
func (Flag) GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "raw",
			Aliases:     []string{"r"},
			Usage:       "Print minimal unstyled text - good for writing to files",
			Destination: &options.Raw,
		},
	}
}
