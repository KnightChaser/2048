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

//go:embed fonts/Roboto-Regular.ttf
var robotoTTF []byte

var (
	fontFace   textv2.Face
	tileColors map[int]color.RGBA
)

func init() {
	// Load font
	tt, err := opentype.Parse(robotoTTF)
	if err != nil {
		log.Fatal(err)
	}

	goFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	fontFace = textv2.NewGoXFace(goFace)

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

// NewGame wraps the engine.Game and initializes the game state.
func NewGame() *Game {
	g := &Game{
		e: engine.NewGame(),
	}

	return g
}

func (g *Game) Update() error {
	directions := map[ebiten.Key]engine.Direction{
		ebiten.KeyArrowLeft:  engine.Left,
		ebiten.KeyArrowRight: engine.Right,
		ebiten.KeyArrowUp:    engine.Up,
		ebiten.KeyArrowDown:  engine.Down,
	}

	moved := false
	for key, dir := range directions {
		if ebiten.IsKeyPressed(key) && g.lastKey != key {
			// Only process the key if it has changed
			movedTmp, _ := g.e.Move(dir)
			moved = moved || movedTmp
			g.lastKey = key
		}
		break
	}

	// Reset lastKey when no arrow key (for 2048 game) is pressed
	if !ebiten.IsKeyPressed(ebiten.KeyArrowLeft) &&
		!ebiten.IsKeyPressed(ebiten.KeyArrowRight) &&
		!ebiten.IsKeyPressed(ebiten.KeyArrowUp) &&
		!ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.lastKey = 0
	}

	if moved {
		// Add a new tile after a successful move
		engine.SpawnTile(&g.e.Board)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{187, 173, 160, 255})
	tileSize := engine.ScreenWidth / engine.GridN
	margin := 8

	// Draw tiles
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
				tw, th := float64(w), float64(h)

				// center within the tile
				posX := float64(x + (tileSize-int(tw))/2)
				posY := float64(y + (tileSize-int(th))/2)

				opts := &textv2.DrawOptions{}
				opts.GeoM.Translate(posX, posY)

				textv2.Draw(screen, s, fontFace, opts)
			}
		}
	}
}

// Layout sets the screen dimensions
type Layout interface {
	Layout(int, int) (int, int)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return engine.ScreenWidth, engine.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(engine.ScreenWidth, engine.ScreenHeight)
	ebiten.SetWindowTitle("2048 Game")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
