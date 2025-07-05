package ui

import (
	"fmt"
	"image/color"

	"2048/engine"

	"github.com/hajimehoshi/ebiten/v2"
	textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
)

func drawMenu(screen *ebiten.Image, bestScore int) {
	// Clear the background
	screen.Fill(color.RGBA{187, 173, 160, 255})

	// Title "2048"
	title := "2048"
	tw, th := textv2.Measure(title, FontFace, 0)
	x := (engine.ScreenWidth - int(tw)) / 2
	y := engine.ScreenHeight / 4
	opts := &textv2.DrawOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	textv2.Draw(screen, title, FontFace, opts)

	// Best score display (will default to 0 until load/save)
	bs := fmt.Sprintf("Best Score: %d", bestScore)
	bw, bh := textv2.Measure(bs, FontFace, 0)
	bx := (engine.ScreenWidth - int(bw)) / 2
	by := y + int(th) + 20
	bOpts := &textv2.DrawOptions{}
	bOpts.GeoM.Translate(float64(bx), float64(by))
	textv2.Draw(screen, bs, FontFace, bOpts)

	// Prompt
	prompt := "Press Enter to Play!"
	pw, _ := textv2.Measure(prompt, FontFace, 0)
	px := (engine.ScreenWidth - int(pw)) / 2
	py := by + int(bh) + 40
	pOpts := &textv2.DrawOptions{}
	pOpts.GeoM.Translate(float64(px), float64(py))
	textv2.Draw(screen, prompt, FontFace, pOpts)
}

func updateMenu(a *App) {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		// Rese the engine and clear any leftover key state
		a.engine = engine.NewGame()
		a.scene = ScenePlay
	}
}
