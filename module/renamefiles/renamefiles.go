package renamefiles

import (
	"crypto/md5"  //nolint:gosec
	"crypto/sha1" //nolint:gosec
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pterm/pterm"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/cli"
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
					pterm.Warning.Println("Renaming files recursively!")
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

				if loadBackup {

					err := utils.ForEachLineInFile(backupFilePath, func(line string) error {
						content := strings.Split(line, "|")
						if line == "\n" || line == "" {
							return nil
						}
						originalName := content[0]
						renamed := content[1]

						err := os.Rename(renamed, originalName)
						if err != nil {
							pterm.Error.Println("Could not restore file", renamed+".", "Did you rename/move/delete it? If not, then it was a duplicate.")
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

					if info == nil {
						continue
					}

					if info.Name() == filepath.Base(backupFilePath) {
						continue
					}

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

					hasher := sha1.New() //nolint:gosec

					switch option {
					case "sha-1":
						hasher = sha1.New() //nolint:gosec
					case "md5":
						hasher = md5.New() //nolint:gosec
					}

					if _, err := io.Copy(hasher, content); err != nil {
						return err
					}

					err = content.Close()
					if err != nil {
						return err
					}

					var hasherContent []byte

					pterm.Printf("%s -> %s%s\n", file, hex.EncodeToString(hasher.Sum(hasherContent)), filepath.Ext(file))

					newName := filepath.Dir(file) + string(os.PathSeparator) + hex.EncodeToString(hasher.Sum(hasherContent)) + path.Ext(file)
					err = os.Rename(file, newName)
					if err != nil {
						pterm.Error.Println(err)
					}

					if backup {
						utils.WriteFile(backupFilePath, []byte(file+"|"+newName+"\n"), true)
					}

				}

				if backup {
					file, err := ioutil.ReadFile(backupFilePath)
					if err != nil {
						return err
					}

					lines := strings.Split(string(file), "\n")
					lines = utils.UniqueStringSlice(lines)
					sort.Strings(lines)
					lines = append(lines, "")

					for i, v := range lines {
						if v == "\n" {
							lines = append(lines[:i], lines[i+1:]...)
						}
					}

					utils.WriteFile(backupFilePath, []byte(strings.Join(lines, "\n")), false)
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
