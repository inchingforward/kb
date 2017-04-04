package main

import canvas "github.com/oskca/gopherjs-canvas"

// GameOverScreen displays when the user has finished a game.
type GameOverScreen struct{}

// Update handles updating objects for the game over screen.
func (state GameOverScreen) Update(ctx *canvas.Context2D) {

}

// Draw handles drawing for the game over screen.
func (state GameOverScreen) Draw(ctx *canvas.Context2D) {
	ctx.FillStyle = "black"
	ctx.Font = "48px serif"
	ctx.FillText("Game Over", canvasWidth/2, canvasHeight/2, canvasWidth)
}

// KeyPressed handles key presses for the game over screen.
func (state GameOverScreen) KeyPressed(keyCode int, pressed bool) {
	SetCurrentGameScreen(SettingsScreen{})
}
