package main

import canvas "github.com/oskca/gopherjs-canvas"

// PlayScreen is displayed while the player is playing the game.
type PlayScreen struct{}

// Update handles object updates for the play screen.
func (state PlayScreen) Update(ctx *canvas.Context2D) {

}

// Draw handles drawing for the play screen.
func (state PlayScreen) Draw(ctx *canvas.Context2D) {
	ctx.FillStyle = "black"
	ctx.Font = "48px serif"
	ctx.FillText("Play", canvasWidth/2, canvasHeight/2, canvasWidth)
}

// KeyPressed handles key presses for the play screen.
func (state PlayScreen) KeyPressed(keyCode int, pressed bool) {
	SetCurrentGameScreen(GameOverScreen{})
}
