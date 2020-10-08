package cli

import (
	"log"
	"strings"
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

type DependingFlag struct {
	Name  string
	Value interface{}
}

func DependingFlags(flags []DependingFlag) {
	var flagCount int
	var names []string

	for _, df := range flags {
		value := df.Value
		name := df.Name
		names = append(names, name)

		if value != nil && value != false && value != "" {
			flagCount++
		}
	}

	if flagCount != len(flags) && flagCount != 0 {
		log.Fatal("these flags depend on each other: " + strings.Join(names, ", "))
	}
}
