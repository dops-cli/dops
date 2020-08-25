package cli

import (
	"log"
)

// IncompatibleFlags checks that only one provided flag is not nil
func IncompatibleFlags(flags ...interface{}) {
	var flagCount int

	for _, flag := range flags {
		if flag != nil && flag != false && flag != "" {
			flagCount++
		}
	}

	if flagCount > 1 {
		log.Fatalf("the flags are not compatible with each other")
	}
}
