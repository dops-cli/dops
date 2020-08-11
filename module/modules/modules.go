package modules

import (
	"github.com/urfave/cli/v2"
)

type Module struct{}

func (Module) GetCommands() []*cli.Command {
	panic("implement me")
}
