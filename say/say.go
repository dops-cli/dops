package say

import (
	"fmt"
	"log"
	"os"

	"github.com/dops-cli/dops/global"
	"github.com/dops-cli/dops/global/options"
	"github.com/dops-cli/dops/progressbar"
	"github.com/dops-cli/dops/say/color"
)

var (
	// InfoPrefix should be used, when any kind of information is shown to the user, which is not part of the modules result
	InfoPrefix = color.SHiMagenta("[info] ")

	// WarningPrefix should be used, when a warning is displayed to the user
	WarningPrefix = color.SYellow("[warning] ")

	// ErrorPrefix should be used, when the module fails at something, but keeps running.
	// If the error is fatal, use say.Fatal
	ErrorPrefix = color.SHiRed("[error] ")

	// SuccessPrefix should be used, when something succeeded.
	SuccessPrefix = color.SGreen("[success] ")
)

func p(prefix string, text ...interface{}) {
	if options.OutputRaw {
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

// Success outputs formatted text to the terminal.
func Success(text ...interface{}) {
	p(SuccessPrefix, text...)
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
	global.CviewApp.Stop()
	log.Fatal(text...)
}

func ProgressBar(totalSteps int) *progressbar.ProgressBar {
	bar := progressbar.NewOptions(
		totalSteps,
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionFullWidth(),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionSpinnerType(14),
	)
	bar.RenderBlank()

	return bar
}
