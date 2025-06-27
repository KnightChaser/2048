package main

import (
	"log"

	"2048/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	board [engine.GridN][engine.GridN]int // The game board
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
	return engine.ScreenWidth, engine.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(engine.ScreenWidth, engine.ScreenHeight)
	ebiten.SetWindowTitle("2048 Game")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
