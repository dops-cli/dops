package ci

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
			Name:        "ci",
			Aliases:     []string{"cd"},
			Usage:       "Runs dops in CI/CD mode - disables fancy styling like progressbars, etc.",
			Destination: &options.Raw,
		},
	}
}
