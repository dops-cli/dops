package extract

import (
	"regexp"
	"testing"
)

func TestPredefinedRegexps(t *testing.T) {
	for _, cmd := range RegexList {
		t.Run(cmd.Name, func(t *testing.T) {
			r, err := regexp.Compile(cmd.Regex)
			if err != nil {
				t.Error(err)
			}

			for _, match := range cmd.Matches {
				if !r.MatchString(match) {
					t.Error("regex does not match, but should")
				}
			}

			for _, fail := range cmd.Fails {
				if r.MatchString(fail) {
					t.Error("regex does match, but should not")
				}
			}
		})
	}
}
