package say

import (
	"fmt"
	"log"

	"github.com/dops-cli/dops/global/options"
	"github.com/dops-cli/dops/progressbar"
	"github.com/dops-cli/dops/progressbar/decor"
	"github.com/dops-cli/dops/say/color"
)

const (
	// FooterPriority is used as priority for footer progressbars
	FooterPriority = 10000
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
	if options.Raw {
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
	log.Fatal(text...)
}

// ProgressBar hows a simple progressbar with a set maximum
func ProgressBar(totalSteps int64) *progressbar.Bar {

	p := progressbar.New()

	bar := p.AddBar(totalSteps,
		progressbar.PrependDecorators(
			decor.Percentage(),
		),
		progressbar.AppendDecorators(
			decor.EwmaETA(decor.ET_STYLE_GO, 90),
		),
	)

	bar.SetPriority(1)

	return bar
}

// ProgressBarFooter shows a simple progressbar with a set maximum at the bottom of the terminal
func ProgressBarFooter(totalSteps int64) *progressbar.Bar {

	p := progressbar.New()

	bar := p.AddBar(totalSteps,
		progressbar.PrependDecorators(
			decor.Percentage(),
		),
		progressbar.AppendDecorators(
			decor.EwmaETA(decor.ET_STYLE_GO, 90),
		),
	)

	bar.SetPriority(FooterPriority)

	return bar
}
