package debug

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
			Aliases:     []string{"d"},
			Name:        "debug",
			Usage:       "Enables debugging mode - only useful if you are working on dops",
			Destination: &options.IsDebug,
		},
	}
}
