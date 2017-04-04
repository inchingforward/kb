package main

import canvas "github.com/oskca/gopherjs-canvas"

// SettingsScreen allows the player to choose game settings before the game starts.
type SettingsScreen struct{}

// Update updates objects on the settings screen.
func (state SettingsScreen) Update(ctx *canvas.Context2D) {

}

// Draw draws objects on the settings screen.
func (state SettingsScreen) Draw(ctx *canvas.Context2D) {
	ctx.FillStyle = "black"
	ctx.Font = "48px serif"
	ctx.FillText("Settings", canvasWidth/2, canvasHeight/2, canvasWidth)
}

// KeyPressed handles key presses for the settings screen.
func (state SettingsScreen) KeyPressed(keyCode int, pressed bool) {
	SetCurrentGameScreen(PlayScreen{})
}
