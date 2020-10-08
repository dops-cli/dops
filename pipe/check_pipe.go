package pipe

import (
	"os"
)

func IsPiped() bool {
	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		return true
	} else {
		return false
	}
}
