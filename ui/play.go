package ui

import (
	"image/color"
	"strconv"

	"2048/engine"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// processArrows handles arrow-key input once per press,
// applies a move, and returns true if the board changed. (tiles moved or merged)
func processArrows(a *App) bool {
	var direction engine.Direction
	var keyPressed bool

	// Did the user press an arrow key?
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		direction, keyPressed = engine.Left, true
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		direction, keyPressed = engine.Right, true
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		direction, keyPressed = engine.Up, true
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		direction, keyPressed = engine.Down, true
	}

	if keyPressed {
		if moved, _ := a.engine.Move(direction); moved {
			return true
		}
	}
	return false
}

// updatePlay handles game logic for the play scene.
func updatePlay(a *App) {
	// Press M at any time to abandon the game and return to menu
	if ebiten.IsKeyPressed(ebiten.KeyM) {
		a.engine = nil
		a.scene = SceneMenu
		return
	}

	if moved := processArrows(a); moved {
		// If the board changed, spawn a new tile to keep the game going
		engine.SpawnTile(&a.engine.Board)
	}

	if !a.engine.CanMove() {
		// end of game
		if a.engine.Score > a.bestScore {
			a.bestScore = a.engine.Score
		}
		a.scene = SceneGameOver
	}
}

// drawPlay renders the game board and HUD.
func drawPlay(screen *ebiten.Image, g *engine.Game) {
	// Background for the board area (starts below the HUD)
	boardBg := color.RGBA{187, 173, 160, 255}
	vector.DrawFilledRect(screen,
		0,
		float32(HUDHeight), // Start Y at HUDHeight
		float32(engine.ScreenWidth),
		float32(engine.ScreenHeight-HUDHeight),
		boardBg,
		false)

	// Compute tile dimensions relative to the new board area
	boardSize := float64(engine.ScreenHeight - HUDHeight)
	tileSize := boardSize / float64(engine.GridN)
	margin := 8.0
	innerSize := tileSize - 2*margin

	for r := 0; r < engine.GridN; r++ {
		for c := 0; c < engine.GridN; c++ {
			// Calculate position of the empty cell background
			cellX := float64(c)*tileSize + margin
			cellY := float64(r)*tileSize + margin + HUDHeight

			// Draw the background for an empty cell first
			vector.DrawFilledRect(screen,
				float32(cellX), float32(cellY),
				float32(innerSize), float32(innerSize),
				TileColors[0].Background, false)

			v := g.Board[r][c]
			if v == 0 {
				continue // Skip drawing number for empty tiles
			}

			// A tile with a value exists, so draw its specific background over the empty cell
			colors := TileColors[v]
			vector.DrawFilledRect(screen,
				float32(cellX), float32(cellY),
				float32(innerSize), float32(innerSize),
				colors.Background, false)

			// Draw the number on the tile with perfect centering
			s := strconv.Itoa(v)
			// Use the font's metrics to get accurate dimensions for centering
			boundsX, boundsY := textv2.Measure(s, LargeFace, LargeFace.Metrics().CapHeight)

			// Calculate position to center the text inside the tile
			px := cellX + (innerSize-boundsX)/2
			py := cellY + (innerSize+boundsY)/2 // This formula correctly centers vertically

			opts := &textv2.DrawOptions{}
			opts.GeoM.Translate(px, py)
			opts.ColorScale.SetR(float32(colors.Foreground.R) / 255.0)
			opts.ColorScale.SetG(float32(colors.Foreground.G) / 255.0)
			opts.ColorScale.SetB(float32(colors.Foreground.B) / 255.0)
			opts.ColorScale.SetA(float32(colors.Foreground.A) / 255.0)
			textv2.Draw(screen, s, LargeFace, opts)
		}
	}
}
