package verbose

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
			Aliases:     []string{"v"},
			Name:        "verbose",
			Usage:       "Enables verbose mode - print more information when running a command",
			Destination: &options.Verbose,
		},
	}
}
