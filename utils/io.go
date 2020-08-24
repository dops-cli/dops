package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dops-cli/dops/say"
)

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
