package modules

import (
	"github.com/dops-cli/dops/modules/bulkdownload"
	"github.com/dops-cli/dops/modules/debug"
	"github.com/dops-cli/dops/modules/update"
	"github.com/urfave/cli/v2"
)

var RegisteredModules = []Module{debug.Module{}, bulkdownload.Module{}, update.Module{}}

type Module interface {
	GetFlags() []cli.Flag
	GetCommands() []*cli.Command
}
