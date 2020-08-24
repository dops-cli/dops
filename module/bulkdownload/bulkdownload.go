package bulkdownload

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/urfave/cli/v2"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/say/color"
)

var wg sync.WaitGroup

// Module returns the created module
type Module struct{}

// GetCommands returns the commands of the module
func (Module) GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "bulkdownload",
			Usage: "Download multiple files from a list",
			Description: `Bulkdownload downloads all files from a list. 
You can set how many files should be downloaded concurrently..`,
			Category: categories.Web,
			Aliases:  []string{"bd"},
			Action: func(c *cli.Context) error {
				inputFile := c.String("input")
				outputDir := c.String("output")
				concurrentDownloads := c.Int("concurrent")

				urls, err := readLines(inputFile)
				if err != nil {
					return err
				}

				downloadMultipleFiles(urls, outputDir, concurrentDownloads)
				wg.Wait()
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "input",
					Aliases: []string{"i"},
					Usage:   "load URLs from `FILE`",
					Value:   "urls.txt",
				},
				&cli.StringFlag{
					Name:        "output",
					Aliases:     []string{"o"},
					Usage:       "save the downloaded files to `DIR`",
					DefaultText: "current directory",
				},
				&cli.IntFlag{
					Name:    "concurrent",
					Aliases: []string{"c"},
					Usage:   "downloads `NUMBER` files concurrently",
					Value:   3,
				},
			},
		},
	}
}

func downloadMultipleFiles(urls []string, outputDir string, concurrentDownloads int) {
	say.Text("Downloading files...")

	guard := make(chan struct{}, concurrentDownloads)

	for index, URL := range urls {
		guard <- struct{}{}
		go func(URL string, outputDir string, index int) {
			err := downloadFile(URL, outputDir, index, len(urls))
			if err != nil {
				fmt.Println(err)
			}
			<-guard
		}(URL, outputDir, index)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func downloadFile(URL string, outputDir string, index int, total int) error {
	wg.Add(1)
	say.Text(fmt.Sprintf("Downloading %s [%s/%s]", URL, color.HiGreenString(strconv.Itoa(index+1)), color.GreenString(strconv.Itoa(total))))
	response, err := http.Get(URL) //nolint:gosec
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println(response.Status)
	}

	file := filepath.Base(URL)

	if outputDir != "" {

		err = os.MkdirAll(outputDir, 0770)
		if err != nil {
			return err
		}

		outputDir += string(os.PathSeparator)
	}

	out, err := os.Create(filepath.FromSlash(outputDir) + file)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		say.Fatal(err)
	}
	wg.Done()

	return nil
}
