package randomgenerator

import (
	"crypto/sha512"
	"encoding/binary"
	"io"
	"math"
	"math/rand"
	"time"

	"github.com/pterm/pterm"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/cli"
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
					Examples: []cli.Example{
						{
							ShortDescription: "Generate a random string with 15 letters",
							Usage:            "dops random-generator string --length 15",
							GenerateSVG:      true,
						},
						{
							ShortDescription: "Generate a random string with 100 letters and a custom charset",
							Usage:            "dops random-generator string --chars abcde12$# --length 100",
							GenerateSVG:      true,
						},
						{
							ShortDescription: "Generate a random number with a minimum of 0 and a maximum of 1000",
							Usage:            "dops random-generator integer --max 1000",
							GenerateSVG:      true,
						},
					},
					Action: func(context *cli.Context) error {
						setSeed(context.String("seed"))
						charset := context.String("chars")
						pterm.Println(StringWithCharset(context.Int("length"), charset))
						return nil
					},
					Flags: []cli.Flag{&cli.StringFlag{
						Aliases: []string{"c"},
						Name:    "chars",
						Usage:   "Use `CHARS` to generate a random string",
						Value:   CharsetNumbersAndLetters,
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

						pterm.Println(rand.Intn(max+1-min) + min) //nolint:gosec
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
							Value: math.MaxInt32,
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
			pterm.Fatal.Println(err)
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
