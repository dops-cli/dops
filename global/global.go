package global

import (
	"github.com/dops-cli/dops/cli"
)

var (
	// CliFlags contains all global flags for dops
	CliFlags []cli.Flag

	// CliCommands contains all modules
	CliCommands []*cli.Command
)
