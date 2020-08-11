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

func Raw(text ...interface{}) {
	p("", text...)
}

func Text(text ...interface{}) {
	p(DefaultPrefix, text...)
}

func Info(text ...interface{}) {
	p(InfoPrefix, text...)
}

func Warning(text ...interface{}) {
	p(WarningPrefix, text...)
}

func Error(text ...interface{}) {
	p(ErrorPrefix, text...)
}
