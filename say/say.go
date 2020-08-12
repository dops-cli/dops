package say

import (
	"fmt"
	"github.com/dops-cli/dops/flags/raw"
	"github.com/dops-cli/dops/say/color"
	"log"
)

var (
	DefaultPrefix = color.CyanString("[dops] ")
	InfoPrefix    = color.HiMagentaString("[dops - info] ")
	WarningPrefix = color.YellowString("[dops - warning] ")
	ErrorPrefix   = color.HiRedString("[dops - error] ")
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

// Raw outputs formatted text to the terminal.
func Raw(text ...interface{}) {
	p("", text...)
}

// Text outputs formatted text to the terminal.
func Text(text ...interface{}) {
	p(DefaultPrefix, text...)
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
