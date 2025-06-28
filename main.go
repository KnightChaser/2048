package main

import (
	_ "embed"
	"image/color"
	"log"
	"strconv"

	"2048/engine"

	"github.com/hajimehoshi/ebiten/v2"
	textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed fonts/0xProto-Regular.ttf
var zeroXProtoTTF []byte

var (
	// face is the base font.Face for measuring
	face font.Face
	// fontFace wraps face for text/v2 drawing
	fontFace textv2.Face
	// tileColors maps values to RGBA colors
	tileColors map[int]color.RGBA
)

func init() {
	// Load base font
	tt, err := opentype.Parse(zeroXProtoTTF)
	if err != nil {
		log.Fatal(err)
	}

	face, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Wrap for text/v2
	fontFace = textv2.NewGoXFace(face)

	// Tile color palette
	tileColors = map[int]color.RGBA{
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

type Game struct {
	e       *engine.Game
	lastKey ebiten.Key
}

// NewGame creates a fresh game instance
func NewGame() *Game {
	return &Game{e: engine.NewGame()}
}

// Update processes input and game logic
func (g *Game) Update() error {
	dirs := map[ebiten.Key]engine.Direction{
		ebiten.KeyArrowLeft:  engine.Left,
		ebiten.KeyArrowRight: engine.Right,
		ebiten.KeyArrowUp:    engine.Up,
		ebiten.KeyArrowDown:  engine.Down,
	}

	moved := false
	for key, dir := range dirs {
		if ebiten.IsKeyPressed(key) {
			if g.lastKey != key {
				movedTmp, _ := g.e.Move(dir)
				moved = moved || movedTmp
				g.lastKey = key
			}
			break
		}
	}

	// Reset lastKey when no arrow pressed
	if !ebiten.IsKeyPressed(ebiten.KeyArrowLeft) &&
		!ebiten.IsKeyPressed(ebiten.KeyArrowRight) &&
		!ebiten.IsKeyPressed(ebiten.KeyArrowUp) &&
		!ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.lastKey = 0
	}

	if moved {
		engine.SpawnTile(&g.e.Board)
	}

	return nil
}

// Draw renders the grid and tiles
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{187, 173, 160, 255})
	tileSize := engine.ScreenWidth / engine.GridN
	margin := 8

	for r := 0; r < engine.GridN; r++ {
		for c := 0; c < engine.GridN; c++ {
			x := c*tileSize + margin
			y := r*tileSize + margin
			v := g.e.Board[r][c]
			col, ok := tileColors[v]
			if !ok {
				col = tileColors[0]
			}

			// Draw tile background
			vector.DrawFilledRect(screen,
				float32(x), float32(y),
				float32(tileSize-2*margin), float32(tileSize-2*margin),
				col /* antialias */, false)

			// Draw number
			if v != 0 {
				s := strconv.Itoa(v)
				w, h := textv2.Measure(s, fontFace, 0)

				// Center text in tile
				posX := float64(x + (tileSize-int(w))/2)
				posY := float64(y + (tileSize-int(h))/2)
				opts := &textv2.DrawOptions{}
				opts.GeoM.Translate(posX, posY)
				textv2.Draw(screen, s, fontFace, opts)
			}
		}
	}
}

// Layout sets the screen dimensions
func (g *Game) Layout(_, _ int) (int, int) {
	return engine.ScreenWidth, engine.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(engine.ScreenWidth, engine.ScreenHeight)
	ebiten.SetWindowTitle("2048 Game")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
