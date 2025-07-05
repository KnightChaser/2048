// ui/hud.go
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
	// Increase HUD height to comfortably fit two lines of text.
	HUDHeight = 90
)

// drawHUD draws the heads-up display (HUD) at the top of the screen.
func drawHUD(screen *ebiten.Image, score, best int) {
	// Background bar
	barBg := color.RGBA{187, 173, 160, 255}
	vector.DrawFilledRect(screen,
		0, 0,
		float32(engine.ScreenWidth), float32(HUDHeight),
		barBg, false)

	// Define layout for the score widgets
	const widgetWidth = 120
	const widgetPadding = 20

	// Draw Score, Best, and Menu widgets
	drawScoreWidget(screen, "SCORE", score, widgetPadding)
	drawScoreWidget(screen, "BEST", best, widgetPadding*2+widgetWidth)
	drawMenuWidget(screen, "MENU (M)", engine.ScreenWidth-widgetWidth-widgetPadding)
}

// drawScoreWidget draws a single box with a title and a right-aligned value.
func drawScoreWidget(screen *ebiten.Image, title string, value int, xPos float64) {
	// Widget background
	widgetBg := color.RGBA{143, 122, 102, 255}
	const widgetHeight = 60
	const widgetWidth = 120
	widgetY := (HUDHeight - widgetHeight) / 2 // Center vertically in HUD
	vector.DrawFilledRect(screen, float32(xPos), float32(widgetY), widgetWidth, widgetHeight, widgetBg, false)

	// Draw Title (e.g., "SCORE")
	titleColor := color.RGBA{238, 228, 218, 255}
	titleBoundsX, _ := textv2.Measure(title, MediumFace, 0)
	titleX := xPos + (widgetWidth-titleBoundsX)/2
	titleY := float64(widgetY + 25) // Position in top half

	opts := &textv2.DrawOptions{}
	opts.GeoM.Translate(titleX, titleY)
	opts.ColorScale.SetR(float32(titleColor.R / 255.0))
	opts.ColorScale.SetG(float32(titleColor.G / 255.0))
	opts.ColorScale.SetB(float32(titleColor.B / 255.0))
	opts.ColorScale.SetA(float32(titleColor.A / 255.0))
	textv2.Draw(screen, title, MediumFace, opts)

	// Draw Value (e.g., "4096")
	valueStr := fmt.Sprintf("%d", value)
	_, valueBoundsY := textv2.Measure(valueStr, LargeFace, 0)
	valueX := xPos + (widgetWidth-valueBoundsY)/2
	valueY := float64(widgetY + 58) // Position in bottom half

	opts.GeoM.Reset()
	opts.GeoM.Translate(valueX, valueY)
	opts.ColorScale.SetR(255) // White text for values
	opts.ColorScale.SetG(255)
	opts.ColorScale.SetB(255)
	opts.ColorScale.SetA(255)
	textv2.Draw(screen, valueStr, LargeFace, opts)
}

// drawMenuWidget is a simpler widget for the menu button.
func drawMenuWidget(screen *ebiten.Image, text string, xPos float64) {
	// Widget background (same as score)
	widgetBg := color.RGBA{143, 122, 102, 255}
	const widgetHeight = 60
	const widgetWidth = 120
	widgetY := (HUDHeight - widgetHeight) / 2
	vector.DrawFilledRect(screen, float32(xPos), float32(widgetY), widgetWidth, widgetHeight, widgetBg, false)

	// Draw Text
	textColor := color.RGBA{238, 228, 218, 255}
	boundsX, boundsY := textv2.Measure(text, MediumFace, 0)
	textX := xPos + (widgetWidth-boundsX)/2
	textY := float64(widgetY) + (widgetHeight+boundsY)/2 // Vertically center text in the box

	opts := &textv2.DrawOptions{}
	opts.GeoM.Translate(textX, textY)
	opts.ColorScale.SetR(float32(textColor.R / 255.0))
	opts.ColorScale.SetG(float32(textColor.G / 255.0))
	opts.ColorScale.SetB(float32(textColor.B / 255.0))
	opts.ColorScale.SetA(float32(textColor.A / 255.0))
	textv2.Draw(screen, text, MediumFace, opts)
}

