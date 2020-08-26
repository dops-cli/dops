// minimal example CLI used for binary size checking

package main

import (
	"log"

	"github.com/dops-cli/dops/cli"
)

func main() {
	err := (&cli.App{}).Run([]string{})
	if err != nil {
		log.Fatal(err)
	}
}
