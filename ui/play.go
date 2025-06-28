package ui

import (
	"image/color"
	"strconv"

	"2048/engine"

	"github.com/hajimehoshi/ebiten/v2"
	textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// processArrows handles arrow-key input once per press,
// applies a move, and returns true if the board changed. (tiles moved or merged)
func processArrows(a *App) bool {
	directions := map[ebiten.Key]engine.Direction{
		ebiten.KeyArrowLeft:  engine.Left,
		ebiten.KeyArrowRight: engine.Right,
		ebiten.KeyArrowUp:    engine.Up,
		ebiten.KeyArrowDown:  engine.Down,
	}

	moved := false
	for key, direction := range directions {
		if ebiten.IsKeyPressed(key) {
			if a.lastKey != key {
				if m, _ := a.engine.Move(direction); m {
					moved = true
				}
				a.lastKey = key // remember last key pressed
			}
			break // only process one key at a time
		}
	}

	// Reset lastkey when no arrow keys are pressed
	if !ebiten.IsKeyPressed(ebiten.KeyArrowLeft) &&
		!ebiten.IsKeyPressed(ebiten.KeyArrowRight) &&
		!ebiten.IsKeyPressed(ebiten.KeyArrowUp) &&
		!ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		a.lastKey = 0
	}

	return moved
}

// updatePlay handles game logic for the play scene.
func updatePlay(a *App) {
	// Press M at any time to abandon the game and return to menu
	if ebiten.IsKeyPressed(ebiten.KeyM) {
		a.engine = nil
		a.lastKey = 0 // reset last key
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
	// Background for the board area
	boardBg := color.RGBA{205, 193, 180, 255}
	vector.DrawFilledRect(screen,
		0, 0,
		float32(engine.ScreenWidth), float32(engine.ScreenHeight),
		boardBg, false)

	// Compute tile dimensions
	tileSize := engine.ScreenWidth / engine.GridN
	margin := 8
	inner := tileSize - 2*margin

	// Draw each tile (row by row)
	for r := 0; r < engine.GridN; r++ {
		for c := 0; c < engine.GridN; c++ {
			x := c*tileSize + margin
			y := r*tileSize + margin
			v := g.Board[r][c]

			// Background color for this tile
			col, ok := TileColors[v]
			if !ok {
				col = TileColors[0]
			}

			vector.DrawFilledRect(screen,
				float32(x), float32(y),
				float32(inner), float32(inner),
				col, false)

			// Draw number if non-zero
			if v != 0 {
				s := strconv.Itoa(v)
				w, h := textv2.Measure(s, FontFace, 0)
				// center text
				px := float64(x + (inner-int(w))/2)
				py := float64(y + (inner-int(h))/2)
				opts := &textv2.DrawOptions{}
				opts.GeoM.Translate(px, py)
				textv2.Draw(screen, s, FontFace, opts)
			}
		}
	}
}
