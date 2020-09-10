package open

import (
	"github.com/pkg/browser"

	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/utils"
)

// URL command
func URL() *cli.Command {
	return &cli.Command{
		Name:    "url",
		Aliases: []string{"browser", "u"},
		Usage:   "Open `URL` with the standard browser",
		Examples: []cli.Example{
			{
				ShortDescription: "Opens marvinjwendt.com in your standard browser",
				Usage:            "dops open url https://marvinjwendt.com",
			},
			{
				ShortDescription: "Opens URL from stdin in your standard browser",
				Usage:            `echo "https://marvinjwendt.com" | dops open url`,
			},
			{
				GenerateSVG:      true,
				ShortDescription: "Returns an error if no standard browser is set",
				Usage:            "dops open url https://marvinjwendt.com",
			},
		},
		Description: "This modules locates the standard browser of the system and opens a specific URL.",
		Action: func(context *cli.Context) error {

			if context.Args().Len() > 0 {
				err := browser.OpenURL(context.Args().First())
				if err != nil {
					return err
				}
			} else {
				input := utils.Input(context.String("input"))
				err := browser.OpenURL(input)
				if err != nil {
					return err
				}
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases: []string{"i"},
				Name:    "input",
				Usage:   "input takes a URL from a path, URL or stdin if it's not set",
			},
		},
	}
}
