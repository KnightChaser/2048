package ui

import (
	"fmt"
	"image/color"

	"2048/engine"

	"github.com/hajimehoshi/ebiten/v2"
	textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func updateGameOver(a *App) {
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		// Reset the game engine and switch to play scene
		a.engine = engine.NewGame()
		a.scene = ScenePlay
	}

	if ebiten.IsKeyPressed(ebiten.KeyM) {
		// Reset the game engine and switch to menu scene
		a.scene = SceneMenu
	}
}

// drawGameOver overlays a semi-transparent backdrop and centered messages.
func drawGameOver(screen *ebiten.Image, score int) {
	// Dark overlay
	overlayCol := color.RGBA{0, 0, 0, 180} // ~70% opacity
	vector.DrawFilledRect(screen,
		0, 0,
		float32(engine.ScreenWidth), float32(engine.ScreenHeight),
		overlayCol, false)

	// "Game Over" title
	title := "Game Over"
	tw, th := textv2.Measure(title, FontFace, 0)
	tx := (engine.ScreenWidth - int(tw)) / 2
	ty := engine.ScreenHeight / 3
	topts := &textv2.DrawOptions{}
	topts.GeoM.Translate(float64(tx), float64(ty))
	textv2.Draw(screen, title, FontFace, topts)

	// Show final score
	scoreMsg := fmt.Sprintf("Score: %d", score)
	sw, sh := textv2.Measure(scoreMsg, FontFace, 0)
	sx := (engine.ScreenWidth - int(sw)) / 2
	sy := ty + int(th) + 20
	sopts := &textv2.DrawOptions{}
	sopts.GeoM.Translate(float64(sx), float64(sy))
	textv2.Draw(screen, scoreMsg, FontFace, sopts)

	// Instructions
	info := "R: Retry    M: Menu"
	iw, _ := textv2.Measure(info, FontFace, 0)
	ix := (engine.ScreenWidth - int(iw)) / 2
	iy := sy + int(sh) + 30
	iopts := &textv2.DrawOptions{}
	iopts.GeoM.Translate(float64(ix), float64(iy))
	textv2.Draw(screen, info, FontFace, iopts)
}
