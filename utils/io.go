package utils

import (
	"fmt"
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

func FileOrStdout(path string, lines []string) {
	if path == "" {
		for _, s := range lines {
			say.Text(s)
		}
	} else {
		var out string
		for _, s := range lines {
			out += fmt.Sprintf("%v", s) + "\n"
		}
		err := ioutil.WriteFile(path, []byte(out), os.ModeAppend)
		if err != nil {
			say.Fatal(err)
		}
	}
}
