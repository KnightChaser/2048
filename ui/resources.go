// ui/resources.go
package ui

import (
	_ "embed"
	"image/color"
	"log"

	textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed fonts/0xProto-Regular.ttf
var protoTTF []byte

var (
	LargeFace  textv2.Face // big numbers, titles, etc.
	MediumFace textv2.Face // medium numbers, small titles, etc.
)

// TileColors maps tile values to their background and text colors.
var TileColors map[int]struct {
	Background color.RGBA
	Foreground color.RGBA
}

func init() {
	// Load the TTF
	tt, err := opentype.Parse(protoTTF)
	if err != nil {
		log.Fatal("parsing proto TTF:", err)
	}

	// Create the base face (used for measuring)
	const dpi = 72

	largeFontBaseFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal("creating base LargeFace:", err)
	}

	LargeFace = textv2.NewGoXFace(largeFontBaseFace)

	MediumFontBaseFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    18,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal("creating MediumFace:", err)
	}

	MediumFace = textv2.NewGoXFace(MediumFontBaseFace)

	// It's good practice to define foreground (text) color along with background.
	// Light numbers (2, 4) have dark text, darker tiles have light text.
	fgDark := color.RGBA{119, 110, 101, 255}
	fgLight := color.RGBA{249, 246, 242, 255}

	TileColors = map[int]struct {
		Background color.RGBA
		Foreground color.RGBA
	}{
		0:    {Background: color.RGBA{205, 193, 180, 255}, Foreground: fgDark},
		2:    {Background: color.RGBA{238, 228, 218, 255}, Foreground: fgDark},
		4:    {Background: color.RGBA{237, 224, 200, 255}, Foreground: fgDark},
		8:    {Background: color.RGBA{242, 177, 121, 255}, Foreground: fgLight},
		16:   {Background: color.RGBA{245, 149, 99, 255}, Foreground: fgLight},
		32:   {Background: color.RGBA{246, 124, 95, 255}, Foreground: fgLight},
		64:   {Background: color.RGBA{246, 94, 59, 255}, Foreground: fgLight},
		128:  {Background: color.RGBA{237, 207, 114, 255}, Foreground: fgLight},
		256:  {Background: color.RGBA{237, 204, 97, 255}, Foreground: fgLight},
		512:  {Background: color.RGBA{237, 200, 80, 255}, Foreground: fgLight},
		1024: {Background: color.RGBA{237, 197, 63, 255}, Foreground: fgLight},
		2048: {Background: color.RGBA{237, 194, 46, 255}, Foreground: fgLight},
	}
}
