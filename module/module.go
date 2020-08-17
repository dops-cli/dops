package module

import (
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/dops-cli/dops/flags/debug"
	"github.com/dops-cli/dops/flags/raw"
	"github.com/dops-cli/dops/module/bulkdownload"
	"github.com/dops-cli/dops/module/demo"
	"github.com/dops-cli/dops/module/extract"
	"github.com/dops-cli/dops/module/update"
	"github.com/dops-cli/dops/say"
)

// * <<< Add modules and global flags here! >>> *
func init() {
	// Add the global flags
	addGlobalFlag(debug.Flag{})
	addGlobalFlag(raw.Flag{})

	// Add modules
	addModule(bulkdownload.Module{})
	addModule(extract.Module{})
	addModule(update.Module{})
	addModule(demo.Module{})
}

// ActiveGlobalFlags contains all global flags.
// If a global flag is not in this slice, it won't be activated.
var ActiveGlobalFlags []GlobalFlag

// ActiveModules contains all available modules.
// If a module is not in this slice, it won't be activated.
// Except for the module `modules`, which is registered in the main package.
var ActiveModules []Module

var CliApp *cli.App

// Module is the interface of each module available in dops.
// Each module must return at least one command.
type Module interface {
	GetCommands() []*cli.Command
}

// GlobalFlag is the interface of each global flag in dops.
// Each flag module must return at least one flag.
type GlobalFlag interface {
	GetFlags() []cli.Flag
}

func Run(cmd *cli.Command, flags map[string]string) error {

	say.Text("\033[2J")
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()

	args := []string{"dops"}
	args = append(args, cmd.Name)

	for name, value := range flags {
		args = append(args, "-"+name+"="+value)
	}

	say.Info("Running: ", strings.Join(args, " "))

	err := CliApp.Run(args)
	if err != nil {
		return err
	}

	return nil
}

func GetByName(name string) *cli.Command {
	for _, m := range ActiveModules {
		for _, c := range m.GetCommands() {
			if c.Name == name {
				return c
			}
		}
	}
	return nil
}

func addModule(module Module) {
	ActiveModules = append(ActiveModules, module)
}

func addGlobalFlag(flag GlobalFlag) {
	ActiveGlobalFlags = append(ActiveGlobalFlags, flag)
}
