package main

import (
	"log"

	"2048/engine"
	"2048/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(engine.ScreenWidth, engine.ScreenHeight)
	ebiten.SetWindowTitle("2048 Game")
	if err := ebiten.RunGame(ui.NewApp()); err != nil {
		log.Fatal(err)
	}
}
