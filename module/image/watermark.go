package image

import (
	"fmt"
	"github.com/dops-cli/dops/cli"
	"github.com/dops-cli/dops/global/options"
	"github.com/dops-cli/dops/say"
	"github.com/flopp/go-findfont"
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Watermark() *cli.Command {
	return &cli.Command{
		Name:        "watermark",
		Aliases:     []string{"wm"},
		Usage:       "Adds a watermark to an image",
		Examples:    []cli.Example{},
		Description: "",
		Action: func(context *cli.Context) error {

			input := context.Path("input")
			glob := context.String("glob")

			cli.IncompatibleFlags(input, glob)
			cli.IncompatibleFlags(context.String("output"), glob)

			if glob != "" {
				matches, err := filepath.Glob(glob)
				if err != nil {
					return err
				}
				for _, match := range matches {
					f, err := os.Stat(match)
					if err != nil {
						return err
					}
					if f.IsDir() {
						continue
					}
					err = watermarkInput(match, context)
					if err != nil {
						return err
					}
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

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{R: 200, G: 100, A: 255}
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func convertLocationToXY(img image.Image, location string, size float64) (x, y, xAnchor, yAnchor float64) {
	maxX, maxY := img.Bounds().Dx(), img.Bounds().Dy()

	switch location {
	case "bottom right":
	case "br":
		x = float64(maxX) - (10 + size)
		y = float64(maxY) - (10 + size)
		xAnchor, yAnchor = 1, 1
		break

	case "bottom left":
	case "bl":
		x = 10 + size
		y = float64(maxY) - (10 + size)
		xAnchor, yAnchor = 0, 1
		break

	case "top right":
	case "tr":
		x = float64(maxX) - (10 + size)
		y = 10 + size
		xAnchor, yAnchor = 1, 0
		break

	case "top left":
	case "tl":
		x = 10 + size
		y = 10 + size
		xAnchor, yAnchor = 0, 0
		break

	case "center":
	case "c":
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
