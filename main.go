package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 400
	screenHeight = 400
	gridN        = 4
)

type Game struct {
	board [gridN][gridN]int
}

func NewGame() *Game {
	g := &Game{}

	return g
}

func (g *Game) Update() error {
	// TODO:
	// 1) read arrow keys
	// 2) shift/merge g.board rows or columns
	// 3) spawn a new tile if the board changed

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// TODO: render g.board as squares + numbers
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("2048 Game")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
