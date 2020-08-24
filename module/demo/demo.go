package demo

import (
	"github.com/urfave/cli/v2"

	"github.com/dops-cli/dops/categories"
)

// Module returns the created module
type Module struct{}

// GetCommands returns the commands of the module
func (Module) GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "demo",
			Usage:       "Demo module of dops",
			Description: `NOTICE: This module does nothing, except showing all possible flags for an interactive demo.`,
			Category:    categories.Dops,
			Action: func(c *cli.Context) error {
				return nil
			},
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "Boolean"},
				&cli.DurationFlag{Name: "Duration"},
				&cli.Float64Flag{Name: "Float64"},
				&cli.Float64SliceFlag{Name: "Float64 List"},
				&cli.IntFlag{Name: "Int"},
				&cli.IntSliceFlag{Name: "Int List"},
				&cli.PathFlag{Name: "Path"},
				&cli.StringFlag{Name: "String"},
				&cli.StringSliceFlag{Name: "String List"},
				&cli.TimestampFlag{Name: "Timestamp"},
			},
		},
	}
}
