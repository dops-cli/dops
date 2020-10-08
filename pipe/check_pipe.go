package pipe

import (
	"github.com/mattn/go-isatty"
	"os"
)

func IsPiped() bool {
	return !isatty.IsTerminal(os.Stdout.Fd())
}
