package utils

import (
	"github.com/pterm/pterm"
	"os"
)

var OriginalStdout *os.File

func DisableStdout() {
	OriginalStdout = os.Stdout
	_, w, _ := os.Pipe()
	defer w.Close()
	os.Stdout = w
	//os.Stderr = w
	//log.SetOutput(w)
	//color.Output = w
	pterm.SetDefaultOutput(w)
}

func EnableStdout() {
	os.Stdout = OriginalStdout
	pterm.SetDefaultOutput(os.Stdout)
}
