package randomgenerator

import (
	"crypto/sha512"
	"encoding/binary"
	"io"
	"math"
	"math/rand"
	"time"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/say"
)

// Module returns the created module
type Module struct{}

// GetModuleCommands returns the commands of the module
func (Module) GetModuleCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "random-generator",
			Aliases: []string{"rg"},
			Usage:   "Generate random values (string, integers, emails, etc..)",
			Description: `This module generates random values of specific types like string, integer, email etc.
You can set the number of generations and the seed.`,
			Warning:  "The generated random values are not cryptographically secure!",
			Category: categories.Generators,
			Action: func(c *cli.Context) error {
				_ = cli.ShowCommandHelp(c, "")
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Aliases:     []string{"s"},
					Name:        "seed",
					Usage:       "Uses `SEED` for the random generation",
					DefaultText: "calculated by current time (nanoseconds)",
				},
			},
			Subcommands: []*cli.Command{
				{
					Name:    "string",
					Aliases: []string{"s"},
					Usage:   "Generate random strings",
					Action: func(context *cli.Context) error {
						setSeed(context.String("seed"))
						charset := context.String("chars")
						say.Text(StringWithCharset(context.Int("length"), charset))
						return nil
					},
					Flags: []cli.Flag{&cli.StringFlag{
						Aliases: []string{"c"},
						Name:    "chars",
						Usage:   "Use `CHARS` to generate a random string",
						Value:   charsetNumbersAndLetters,
					},
						&cli.IntFlag{
							Aliases: []string{"l"},
							Name:    "length",
							Usage:   "Generate a random string of length `LENGTH`",
							Value:   8,
						},
					},
				},
				{
					Name:    "integer",
					Aliases: []string{"i", "n", "number"},
					Usage:   "Generate random integer",
					Action: func(context *cli.Context) error {
						setSeed(context.String("seed"))

						min := context.Int("min")
						max := context.Int("max")

						say.Text(rand.Intn(max+1-min) + min) //nolint:gosec
						return nil
					},
					Flags: []cli.Flag{
						&cli.IntFlag{
							Name:  "min",
							Usage: "Minimum `NUMBER` to be generated",
						},
						&cli.IntFlag{
							Name:  "max",
							Usage: "Maximum `NUMBER` to be generated",
							Value: math.MaxInt64,
						},
					},
				},
			},
		},
	}
}

func setSeed(seedString string) {
	if seedString != "" {
		h := sha512.New() //nolint:gosec
		_, err := io.WriteString(h, seedString)
		if err != nil {
			say.Fatal(err)
		}
		rand.Seed(int64(binary.BigEndian.Uint64(h.Sum(nil))))
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
	}

}

// StringWithCharset returns a string with a specific chatset and length
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
