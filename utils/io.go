package utils

import (
	"io/ioutil"
	"os"

	"github.com/dops-cli/dops/say"
)

func FileOrStdin(path string) string {
	if path == "" {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			say.Fatal(err)
		}
		return string(bytes)
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		say.Fatal(err)
	}
	return string(file)
}
