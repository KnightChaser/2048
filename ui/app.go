package ui

import (
	"2048/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

type App struct {
	scene     Scene
	engine    *engine.Game
	bestScore int
}

// NewApp initializes a new App instance with the initial scene set to SceneMenu.
func NewApp() *App {
	return &App{
		scene:  SceneMenu,
		engine: nil, // Engine will be initialized lazily (at menu start)
	}
}

// Update processes the current scene and updates the game state accordingly.
func (a *App) Update() error {
	switch a.scene {
	case SceneMenu:
		updateMenu(a)
	case ScenePlay:
		updatePlay(a)
	case SceneGameOver:
		updateGameOver(a)
	}
	return nil
}

// Draw renders the current scene to the provided screen image.
func (a *App) Draw(screen *ebiten.Image) {
	switch a.scene {
	case SceneMenu:
		drawMenu(screen, a.bestScore)
	case ScenePlay:
		drawPlay(screen, a.engine)
		drawHUD(screen, a.engine.Score, a.bestScore)
	case SceneGameOver:
		drawPlay(screen, a.engine)           // show last board
		drawGameOver(screen, a.engine.Score) // overlay + texts
	}
}

// Layout returns the dimensions of the game screen.
func (a *App) Layout(_, _ int) (int, int) {
	return engine.ScreenWidth, engine.ScreenHeight
}
