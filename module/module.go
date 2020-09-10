package module

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/module/ci"
	"github.com/dops-cli/dops/module/open"
	"github.com/dops-cli/dops/module/ping"
	"github.com/dops-cli/dops/module/randomgenerator"

	ciflag "github.com/dops-cli/dops/flags/ci"
	"github.com/dops-cli/dops/flags/debug"
	"github.com/dops-cli/dops/flags/raw"
	"github.com/dops-cli/dops/global"
	"github.com/dops-cli/dops/module/bulkdownload"
	"github.com/dops-cli/dops/module/extract"
	"github.com/dops-cli/dops/module/renamefiles"
	"github.com/dops-cli/dops/module/update"
	"github.com/dops-cli/dops/say"
)

// * <<< Add modules and global flags here! >>> *
func init() {
	// Add the global flags
	addGlobalFlag(debug.Flag{})
	addGlobalFlag(raw.Flag{})
	addGlobalFlag(ciflag.Flag{})

	// Add modules
	addModule(bulkdownload.Module{})
	addModule(extract.Module{})
	addModule(update.Module{})
	// addModule(demo.Module{})
	addModule(renamefiles.Module{})
	addModule(ping.Module{})
	addModule(randomgenerator.Module{})
	addModule(open.Module{})

	addModule(ci.Module{})
}

// CliApp is the main component of dops, which contains all modules and flags
var CliApp *cli.App

// Run runs a specific module with specific flags
func Run(flags []string) error {

	say.Text("\033[2J")
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	_ = clear.Run()

	args := []string{"dops"}
	args = append(args, flags...)

	say.Info("Running: ", strings.Join(args, " "))

	err := CliApp.Run(args)
	if err != nil {
		return err
	}

	return nil
}

// GetByName searches for a module by name and returns it.
// If no module is found, it will return an error.
func GetByName(name string) (*cli.Command, error) {
	for _, c := range global.CliCommands {
		if c.Name == name {
			return c, nil
		}
	}
	return nil, errors.New("could not find module by name " + name)
}

func addModule(module cli.Module) {
	cli.ActiveModules = append(cli.ActiveModules, module)
}

func addGlobalFlag(flag cli.GlobalFlag) {
	cli.ActiveGlobalFlags = append(cli.ActiveGlobalFlags, flag)
}
