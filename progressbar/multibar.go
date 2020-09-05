package progressbar

import (
	"fmt"
	"io"

	"github.com/dops-cli/dops/progressbar/decor"
)

type LogFiller struct {
	Message string
}

func (l LogFiller) Fill(w io.Writer, reqWidth int, stat decor.Statistics) {
	fmt.Fprint(w, l.Message)
}

func MultiBarLog(message string) LogFiller {
	return LogFiller{
		Message: message,
	}
}
