package debug

import (
	"github.com/urfave/cli/v2"
)

// IsDebug returns true if dops is run in debugging mode.
var IsDebug bool

// Flag returns the created flag
type Flag struct{}

// GetFlags returns the global flags
func (Flag) GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Usage:       "Enables debugging mode - only useful if you are working on dops",
			Destination: &IsDebug,
		},
	}
}
