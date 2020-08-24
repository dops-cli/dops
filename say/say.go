package say

import (
	"fmt"
	"log"

	"github.com/dops-cli/dops/flags/raw"
	"github.com/dops-cli/dops/global"
	"github.com/dops-cli/dops/say/color"
)

var (
	// InfoPrefix should be used, when any kind of information is shown to the user, which is not part of the modules result
	InfoPrefix = color.HiMagentaString("[info] ")

	// WarningPrefix should be used, when a warning is displayed to the user
	WarningPrefix = color.YellowString("[warning] ")

	// ErrorPrefix should be used, when the module fails at something, but keeps running.
	// If the error is fatal, use say.Fatal
	ErrorPrefix = color.HiRedString("[error] ")
)

func p(prefix string, text ...interface{}) {
	if raw.OutputRaw {
		prefix = ""
	}
	_, err := fmt.Fprint(color.Output, prefix)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(color.Output, text...)
	if err != nil {
		log.Fatal(err)
	}
}

// Text outputs formatted text to the terminal.
func Text(text ...interface{}) {
	p("", text...)
}

// Info outputs formatted text to the terminal.
func Info(text ...interface{}) {
	p(InfoPrefix, text...)
}

// Warning outputs formatted text to the terminal.
func Warning(text ...interface{}) {
	p(WarningPrefix, text...)
}

// Error outputs formatted text to the terminal.
func Error(text ...interface{}) {
	p(ErrorPrefix, text...)
}

// Fatal outputs formatted text to the terminal.
func Fatal(text ...interface{}) {
	if global.CviewApp != nil {
		global.CviewApp.Stop()
	}
	log.Fatal(text...)
}
