package image

import (
	"github.com/dops-cli/dops/cli"
	"github.com/flopp/go-findfont"
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func Watermark() *cli.Command {
	return &cli.Command{
		Name:        "watermark",
		Aliases:     []string{"wm"},
		Usage:       "Adds a watermark to an image",
		Examples:    []cli.Example{},
		Description: "",
		Action: func(context *cli.Context) error {
			color := context.String("color")
			input := context.Path("input")
			//location := context.Option("location")
			//opacity := context.Float64("opacity")
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

			//x, y := convertLocationToXY(img, location)

			ctx := gg.NewContextForImage(img)

			ctx.FontHeight()
			ctx.SetHexColor(color)
			fontPath, err := findfont.Find("arial.ttf")
			err = ctx.LoadFontFace(fontPath, size)
			if err != nil {
				return err
				//ctx.SetFontFace(basicfont.Face7x13)
			}
			ctx.DrawStringAnchored(watermarkText, float64(ctx.Width()-(10+50)), float64(ctx.Height()-(10+50)), 1, 1)
			//ctx.DrawStringWrapped(watermarkText, float64(ctx.Width() - 10), float64(ctx.Height() - 10), 1, 1, 40, 1, gg.AlignCenter)

			err = ctx.SavePNG(output)
			if err != nil {
				return err
			}

			//m := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
			//x, y := convertLocationToXY(img, location)
			//addLabel(m, x, y, watermarkText)
			//outfile, _ := os.Create(output + ".png")
			//err = png.Encode(outfile, m)
			//if err != nil {
			//	return err
			//}

			return nil
		},
		Flags: []cli.Flag{
			&cli.OptionFlag{
				Aliases: []string{"l"},
				Options: []string{"top right", "tr", "top left", "tl", "bottom right", "br", "bottom left", "bl", "center", "c"},
				Name:    "location",
				Usage:   "Watermark location",
				Value: &cli.Option{
					Option: "br",
				},
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
				Value:   "#fff",
			},
			&cli.Float64Flag{
				Aliases: []string{"op"},
				Name:    "opacity",
				Usage:   "Watermark opacity",
				Value:   70,
			},
			&cli.PathFlag{
				Name:      "input",
				Aliases:   []string{"i"},
				Usage:     "use `FILE` as input",
				TakesFile: true,
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

func convertLocationToXY(img image.Image, location string) (x, y int) {

	maxX, maxY := img.Bounds().Dx(), img.Bounds().Dy()

	switch location {
	case "bottom right":
	case "br":
		x = maxX - 10
		y = maxY - 10
		break
	}

	return x, y
}
