// minimal example CLI used for binary size checking

package main

import (
	"github.com/dops-cli/dops/cli"
)

func main() {
	(&cli.App{}).Run([]string{})
}
