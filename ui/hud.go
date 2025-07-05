package ui

import (
	"fmt"
	"image/color"

	"2048/engine"

	"github.com/hajimehoshi/ebiten/v2"
	textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	HUDHeight = 60
)

// drawHUD draws the heads-up display (HUD) at the top of the screen.
func drawHUD(screen *ebiten.Image, score, best int) {
	// Background bar
	barH := 60
	bg := color.RGBA{143, 122, 102, 255}
	vector.DrawFilledRect(screen,
		0, 0,
		float32(engine.ScreenWidth), float32(barH),
		bg, false)

	// Compute “to-go” (best minus current)
	toGo := best - score
	if toGo < 0 {
		toGo = 0
	}

	// Draw score labels: Score, Best, To go (score)
	margin := 10
	spacing := 40
	x := margin
	labels := []string{
		fmt.Sprintf("Score: %d", score),
		fmt.Sprintf("Best: %d", best),
		fmt.Sprintf("To go: %d", toGo),
	}

	for _, label := range labels {
		w, h := textv2.Measure(label, FontFace, 0)
		y := (barH-int(h))/2 + int(h) // center-vertically
		opts := &textv2.DrawOptions{}
		opts.GeoM.Translate(float64(x), float64(y))
		textv2.Draw(screen, label, FontFace, opts)
		x += int(w) + spacing
	}

	// Menu button on right
	menu := "Menu (M)"
	mw, mh := textv2.Measure(menu, FontFace, 0)
	mx := engine.ScreenWidth - int(mw) - margin
	my := (barH-int(mh))/2 + int(mh)
	mopts := &textv2.DrawOptions{}
	mopts.GeoM.Translate(float64(mx), float64(my))
	textv2.Draw(screen, menu, FontFace, mopts)
}
