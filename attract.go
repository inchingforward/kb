package main

import canvas "github.com/oskca/gopherjs-canvas"

// AttractScreen is the first screen displayed to the player.
type AttractScreen struct{}

// Update handles object updates for the attract screen.
func (state AttractScreen) Update(ctx *canvas.Context2D) {

}

// Draw handles drawing for the attract screen.
func (state AttractScreen) Draw(ctx *canvas.Context2D) {
	ctx.FillStyle = "black"
	ctx.Font = "48px serif"
	ctx.FillText("あえいうお", canvasWidth/2, canvasHeight/2, canvasWidth)
}

// KeyPressed handles key presses for the attract screen.
func (state AttractScreen) KeyPressed(keyCode int, pressed bool) {
	SetCurrentGameScreen(SettingsScreen{})
}
