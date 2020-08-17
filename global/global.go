package global

import (
	"github.com/urfave/cli/v2"
	"gitlab.com/tslocum/cview"
)

var (
	CviewApp   *cview.Application
	CviewTable *cview.Table

	CliFlags    []cli.Flag
	CliCommands []*cli.Command
)
