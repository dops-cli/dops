package modules

import (
	"github.com/dops-cli/dops/modules/bulkdownload"
	"github.com/dops-cli/dops/modules/debug"
	"github.com/urfave/cli/v2"
)

var RegisteredModules = []Module{debug.Module{}, bulkdownload.Module{}}

type Module interface {
	GetFlags() []cli.Flag
	GetCommands() []*cli.Command
}
