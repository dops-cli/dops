package module

import (
	"github.com/dops-cli/dops/flags/debug"
	"github.com/dops-cli/dops/module/bulkdownload"
	"github.com/dops-cli/dops/module/update"
	"github.com/urfave/cli/v2"
)

var RegisteredGlobalFlags = []GlobalFlag{debug.Flag{}}
var RegisteredModules = []Module{bulkdownload.Module{}, update.Module{}}

type Module interface {
	GetCommands() []*cli.Command
}

type GlobalFlag interface {
	GetFlags() []cli.Flag
}
