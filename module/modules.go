package module

import (
	"github.com/dops-cli/dops/flags/debug"
	"github.com/dops-cli/dops/module/bulkdownload"
	"github.com/dops-cli/dops/module/update"
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
