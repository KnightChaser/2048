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

// FontFace is the text/v2.Face used for all UI text
var FontFace textv2.Face

// TileColors maps tile values to RGBA background colors
var TileColors map[int]color.RGBA

func init() {
	// Load the TTF
	tt, err := opentype.Parse(protoTTF)
	if err != nil {
		log.Fatal("parsing proto TTF:", err)
	}

	// Create the base face (used for measuring)
	baseFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal("creating base face:", err)
	}
	FontFace = textv2.NewGoXFace(baseFace)

	// Tile color palette
	TileColors = map[int]color.RGBA{
		0:    {205, 193, 180, 255},
		2:    {238, 228, 218, 255},
		4:    {237, 224, 200, 255},
		8:    {242, 177, 121, 255},
		16:   {245, 149, 99, 255},
		32:   {246, 124, 95, 255},
		64:   {246, 94, 59, 255},
		128:  {237, 207, 114, 255},
		256:  {237, 204, 97, 255},
		512:  {237, 200, 80, 255},
		1024: {237, 197, 63, 255},
		2048: {237, 194, 46, 255},
	}
}
