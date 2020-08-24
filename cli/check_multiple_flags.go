package cli

import (
	"log"
)

// RequireOneFlag checks that only one provided flag is not nil
func RequireOneFlag(flags ...interface{}) {
	var flagCount int

	for _, flag := range flags {
		if flag != nil {
			flagCount++
		}
	}

	if flagCount > 1 {
		log.Fatalf("the flags are not compatible with each other")
	}
}
