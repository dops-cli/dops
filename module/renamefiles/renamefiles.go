package renamefiles

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/dops-cli/dops/cli"

	"github.com/dops-cli/dops/categories"
)

// Module returns the created module
type Module struct{}

// GetModuleCommands returns the commands of the module
func (Module) GetModuleCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "rename-files",
			Aliases: []string{"rf"},
			Usage:   "Renames all selected files to a specific pattern",
			Description: `This module can be used to rename multiple files according to a specified pattern.
The pattern could be a timestamp, or the hashcode of the file, among others.`,
			Category: categories.IO,
			Action: func(context *cli.Context) error {
				option := context.Option("pattern")

				files, err := ioutil.ReadDir(".")
				if err != nil {
					fmt.Println(err.Error())
				}

				hasher := sha1.New()

				for _, file := range files {
					if file.IsDir() {
						continue
					}

					switch option {
					case "sha-1":
						hasher = sha1.New()
						break
					case "md5":
						hasher = md5.New()
						break
					}

					content, _ := os.Open(file.Name())

					if _, err := io.Copy(hasher, content); err != nil {
						log.Fatal(err)
					}

					_ = content.Close()

					fmt.Printf("%s -> %s%s\n", file.Name(), hex.EncodeToString(hasher.Sum(nil)), filepath.Ext(file.Name()))
					err := os.Rename(file.Name(), hex.EncodeToString(hasher.Sum(nil))+path.Ext(file.Name()))

					if err != nil {
						fmt.Println(err.Error())
					}

				}

				return nil
			},
			Flags: []cli.Flag{
				&cli.OptionFlag{
					Name:    "pattern",
					Aliases: []string{"p"},
					Usage:   "Rename all files with `OPTION`",
					Options: []string{"sha-1", "md5"},
				},
			},
		},
	}
}
