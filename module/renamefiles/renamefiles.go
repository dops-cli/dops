package renamefiles

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/utils"
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
				recursive := context.Bool("recursive")
				dir := context.String("directory")
				backup := !context.Bool("disablebackup")
				loadBackup := context.Bool("loadbackup")
				if loadBackup {
					backup = false
				}

				var files []string

				if recursive {
					say.Warning("Renaming files recursively!")
					err := filepath.Walk(dir,
						func(path string, info os.FileInfo, err error) error {
							files = append(files, path)
							return nil
						})
					if err != nil {
						return err
					}
				} else {
					f, err := ioutil.ReadDir(dir)
					if err != nil {
						return err
					}
					for _, info := range f {
						files = append(files, info.Name())
					}
				}

				backupFilePath := filepath.Dir(dir) + string(os.PathSeparator) + ".dops-filename-backup"

				if backup {
					if _, err := os.Stat(backupFilePath); err == nil {
						return errors.New("backup file already exists")
					}
				}

				if loadBackup {

					err := utils.ForEachLineInFile(backupFilePath, func(line string) error {
						content := strings.Split(line, "|")
						originalName := content[0]
						renamed := content[1]

						err := os.Rename(renamed, originalName)
						if err != nil {
							say.Error("Could not restore file", renamed+".", "Did you rename/move/delete it?")
						}
						return nil
					})
					if err != nil {
						return err
					}

					return nil
				}

				for _, file := range files {
					info, err := os.Stat(file)
					if err != nil {
						return err
					}
					if info.IsDir() {
						continue
					}

					content, err := os.Open(file)
					if err != nil {
						return err
					}

					hasher := sha1.New()

					switch option {
					case "sha-1":
						hasher = sha1.New()
						break
					case "md5":
						hasher = md5.New()
						break
					}

					if _, err := io.Copy(hasher, content); err != nil {
						return err
					}

					err = content.Close()
					if err != nil {
						return err
					}

					var hasherContent []byte

					fmt.Printf("%s -> %s%s\n", file, hex.EncodeToString(hasher.Sum(hasherContent)), filepath.Ext(file))

					newName := filepath.Dir(file) + string(os.PathSeparator) + hex.EncodeToString(hasher.Sum(hasherContent)) + path.Ext(file)
					err = os.Rename(file, newName)
					if err != nil {
						return err
					}

					if backup {
						// file, _ = filepath.Abs(file)
						// newName, _ = filepath.Abs(newName)
						utils.WriteFile(backupFilePath, []byte(file+"|"+newName+"\n"), true)
					}

				}

				return nil
			},
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:     "directory",
					Aliases:  []string{"dir", "d"},
					Usage:    "`PATH` in which the files should be renamed",
					Required: true,
				},
				&cli.OptionFlag{
					Name:    "pattern",
					Aliases: []string{"p"},
					Usage:   "Rename all files with `OPTION`",
					Options: []string{"sha-1", "md5"},
				},
				&cli.BoolFlag{
					Name:    "recursive",
					Aliases: []string{"r"},
					Usage:   "Rename files in subdirectories too",
				},
				&cli.BoolFlag{
					Name:    "disablebackup",
					Aliases: []string{"db"},
					Usage:   "Disable file name backups",
				},
				&cli.BoolFlag{
					Name:    "loadbackup",
					Aliases: []string{"l", "lb"},
					Usage:   "Reverts the filenames to the original",
				},
			},
		},
	}
}
