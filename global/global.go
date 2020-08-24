package global

import (
	"gitlab.com/tslocum/cview"

	"github.com/dops-cli/dops/cli"
)

var (
	// CviewApp is the interactive cli main component
	CviewApp *cview.Application

	// CviewTable is the table in the interactive cli, which contains the module list
	CviewTable *cview.Table

	// CliFlags contains all global flags for dops
	CliFlags []cli.Flag

	// CliCommands contains all modules
	CliCommands []*cli.Command
)
