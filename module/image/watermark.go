package image

import (
	"fmt"
	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/global/options"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/utils"
	"github.com/flopp/go-findfont"
	"github.com/fogleman/gg"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Watermark contains the watermark logic
func Watermark() *cli.Command {
	return &cli.Command{
		Name:    "watermark",
		Aliases: []string{"wm"},
		Usage:   "Adds a watermark to an image",
		Examples: []cli.Example{
			{
				ShortDescription: "Adds a watermark to the example.jpg image and saves it as example_watermarked.png",
				Usage:            `dops image watermark --input example.jpg --text "example watermark text" --location "top left" --opacity 50 --output "example_watermarked.png"`,
			},
			{
				ShortDescription: "Adds a watermark to every image with .png ending in this path",
				Usage:            `dops image watermark --glob c/images/*.png --text "example watermark text" --location "top left" --opacity 50`,
			},
		},
		Description: "This module watermark adds a watermark to one or more images from the input with a custom text on one of the corners.",
		Action: func(context *cli.Context) error {

			input := context.Path("input")
			glob := context.String("glob")
			output := context.String("output")

			cli.IncompatibleFlags(input, glob)
			cli.IncompatibleFlags(output, glob)
			cli.DependingFlags([]cli.DependingFlag{
				{
					Name:  "input",
					Value: input,
				},
				{
					Name:  "output",
					Value: output,
				},
			})

			if glob != "" {
				err := utils.Glob(glob, func(path string) error {
					var match = path
					f, err := os.Stat(match)
					if err != nil {
						return err
					}
					if f.IsDir() {
						return nil
					}
					err = watermarkInput(match, context)
					if err != nil {
						return err
					}
					return nil
				})
				if err != nil {
					return err
				}
			} else if input != "" {

				err := watermarkInput(input, context)
				if err != nil {
					return err
				}
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.OptionFlag{
				Aliases:  []string{"l"},
				Options:  []string{"top right", "tr", "top left", "tl", "bottom right", "br", "bottom left", "bl", "center", "c"},
				Name:     "location",
				Usage:    "Watermark location",
				Required: true,
			},
			&cli.StringFlag{
				Aliases:  []string{"t"},
				Name:     "text",
				Usage:    "Watermark text",
				Required: true,
			},
			&cli.Float64Flag{
				Aliases: []string{"s"},
				Name:    "size",
				Usage:   "Watermark size",
				Value:   12,
			},
			&cli.StringFlag{
				Aliases: []string{"c"},
				Name:    "color",
				Usage:   "Watermark color",
				Value:   "#ffffff",
			},
			&cli.IntFlag{
				Aliases: []string{"op"},
				Name:    "opacity",
				Usage:   "Watermark opacity - range 0-100",
				Value:   100,
			},
			&cli.PathFlag{
				Name:      "input",
				Aliases:   []string{"i"},
				Usage:     "use `FILE` as input",
				TakesFile: true,
			},
			&cli.StringFlag{
				Name:    "glob",
				Aliases: []string{"g"},
				Usage:   "uses a `GLOB` pattern to input multiple files",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "outputs to directory `DIR`",
			},
		},
	}
}

func convertLocationToXY(img image.Image, location string, size float64) (x, y, xAnchor, yAnchor float64) {
	maxX, maxY := img.Bounds().Dx(), img.Bounds().Dy()

	switch location {
	case "bottom right", "br":
		x = float64(maxX) - (10 + size)
		y = float64(maxY) - (10 + size)
		xAnchor, yAnchor = 1, 1
		break

	case "bottom left", "bl":
		x = 10 + size
		y = float64(maxY) - (10 + size)
		xAnchor, yAnchor = 0, 1
		break

	case "top right", "tr":
		x = float64(maxX) - (10 + size)
		y = 10 + size
		xAnchor, yAnchor = 1, 0
		break

	case "top left", "tl":
		x = 10 + size
		y = 10 + size
		xAnchor, yAnchor = 0, 0
		break

	case "center", "c":
		x = (float64(maxX) - (size)) / 2
		y = (float64(maxY) - (size)) / 2
		xAnchor, yAnchor = 0.5, 0.5
		break
	}

	return x, y, xAnchor, yAnchor
}

func watermarkInput(input string, context *cli.Context) error {
	watermarkColor := context.String("color")
	location := context.Option("location")
	opacity := context.Int("opacity")
	output := context.String("output")
	size := context.Float64("size")
	watermarkText := context.String("text")

	f, err := os.Open(input)
	if err != nil {
		return err
	}

	var img image.Image

	if filepath.Ext(input) == ".png" {
		img, err = png.Decode(f)
	} else if filepath.Ext(input) == ".jpg" {
		img, err = jpeg.Decode(f)
	}

	if err != nil {
		return err
	}

	img, err = gg.LoadImage(input)
	if err != nil {
		return err
	}

	watermarkColor = strings.ReplaceAll(watermarkColor, "#", "")

	if len(watermarkColor) == 3 {
		say.Fatal("color code must have 6 digits")
	}

	opacity = opacity * 255 / 100

	opacityHex := fmt.Sprintf("%x", opacity)
	if len(opacityHex) == 1 {
		opacityHex = "0" + opacityHex
	}
	if len(watermarkColor) == 6 {
		watermarkColor = watermarkColor + opacityHex
	}

	ctx := gg.NewContextForImage(img)

	ctx.SetHexColor(watermarkColor)
	fontPath, err := findfont.Find("arial.ttf")
	if err != nil {
		log.Fatal(err)
	}
	err = ctx.LoadFontFace(fontPath, size)
	if err != nil {
		return err
		//ctx.SetFontFace(basicfont.Face7x13)
	}

	x, y, xAnchor, yAnchor := convertLocationToXY(img, location, size)

	ctx.DrawStringAnchored(watermarkText, x, y, xAnchor, yAnchor)

	var watermarkedFilename string

	if context.String("glob") != "" {
		watermarkedFilename = strings.TrimSuffix(input, filepath.Ext(input)) + "_watermarked" + filepath.Ext(input)
	} else {
		watermarkedFilename = output
	}

	if !options.Debug {
		err = ctx.SavePNG(watermarkedFilename)
		if err != nil {
			return err
		}
	} else {
		fmt.Println(input + " -> " + watermarkedFilename)
	}

	return nil
}
