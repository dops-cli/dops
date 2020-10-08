package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dops-cli/dops/say"
)

// ForEachLineInFile runs cb over all lines in a file
func ForEachLineInFile(path string, cb func(line string) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err = cb(scanner.Text())
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// WriteFile writes content to path. If append is true, the content will be appended to the file at path.
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
		err := ioutil.WriteFile(path, content, 0600)
		if err != nil {
			say.Fatal(err)
		}
	}

}

// Glob walks over a glob pattern and runs a function on the files.
func Glob(glob string, forEach func(path string) error) error {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return err
	}

	for _, match := range matches {
		err := forEach(match)
		if err != nil {
			return err
		}
	}

	return nil
}

// Input is used for flags, which accept input in any form. Input supports HTTP and HTTPS resources, file paths and stdin.
func Input(path string) string {
	if path == "" {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			say.Fatal(err)
		}
		return string(bytes)
	} else if strings.HasPrefix(path, "https://") || strings.HasPrefix(path, "http://") {
		var client http.Client
		resp, err := client.Get(path)
		if err != nil {
			say.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			return bodyString
		}

		return "Error: " + strconv.Itoa(resp.StatusCode)
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		say.Fatal(err)
	}
	return string(file)
}

// Output is used for flags, which accept output paths. If append is true, the output will be appended to the file at path.
func Output(path string, lines []string, append bool) {
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
