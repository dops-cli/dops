package modules

import (
	"github.com/dops-cli/dops/flags/debug"
	"github.com/dops-cli/dops/modules/bulkdownload"
	"github.com/dops-cli/dops/modules/update"
	"github.com/urfave/cli/v2"
)

var RegisteredGlobalFlags = []GFlag{debug.Flag{}}
var RegisteredModules = []Module{bulkdownload.Module{}, update.Module{}}

type Module interface {
	GetCommands() []*cli.Command
}

type GFlag interface {
	GetFlags() []cli.Flag
}
