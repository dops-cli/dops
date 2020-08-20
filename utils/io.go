package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dops-cli/dops/say"
)

func WriteFile(path string, content []byte, append bool) {
	if append {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			say.Fatal(err)
		}
		defer f.Close()
		if _, err := f.Write(content); err != nil {
			say.Fatal(err)
		}
	} else {
		err := ioutil.WriteFile(path, content, 0644)
		if err != nil {
			say.Fatal(err)
		}
	}

}

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

func FileOrStdout(path string, lines []string, append bool) {
	if path == "" {
		for _, s := range lines {
			say.Text(s)
		}
	} else {
		var out string
		for _, s := range lines {
			out += fmt.Sprintf("%v", s) + "\n"
		}
		WriteFile(path, []byte(out), append)
	}
}
