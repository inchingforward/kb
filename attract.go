package main

import canvas "github.com/oskca/gopherjs-canvas"

const (
	colorTurquoise  = "#1ABC9C"
	colorEmerald    = "#2ECC71"
	colorPeterRiver = "#3498DB"
	colorAmethyst   = "#9B59B6"
	colorWetAsphalt = "#34495E"
	colorSunflower  = "#F1C40F"
	colorCarrot     = "#E67E22"
	colorAlizarin   = "#E74C3C"
	colorClouds     = "#ECF0F1"
	colorConcrete   = "#95A5A6"
)

// AttractScreen is the first screen displayed to the player.
type AttractScreen struct{}

// Update handles object updates for the attract screen.
func (state AttractScreen) Update(ctx *canvas.Context2D) {

}

// Draw handles drawing for the attract screen.
func (state AttractScreen) Draw(ctx *canvas.Context2D) {
	ctx.Font = "36px Noto Sans Japanese"

	height := 36.0

	ctx.FillStyle = colorTurquoise
	ctx.FillText("あえいうお", canvasWidth/2, height*1, canvasWidth)

	ctx.FillStyle = colorEmerald
	ctx.FillText("あえいうお", canvasWidth/2, height*2, canvasWidth)

	ctx.FillStyle = colorPeterRiver
	ctx.FillText("あえいうお", canvasWidth/2, height*3, canvasWidth)

	ctx.FillStyle = colorAmethyst
	ctx.FillText("あえいうお", canvasWidth/2, height*4, canvasWidth)

	ctx.FillStyle = colorWetAsphalt
	ctx.FillText("あえいうお", canvasWidth/2, height*5, canvasWidth)

	ctx.FillStyle = colorSunflower
	ctx.FillText("あえいうお", canvasWidth/2, height*6, canvasWidth)

	ctx.FillStyle = colorCarrot
	ctx.FillText("あえいうお", canvasWidth/2, height*7, canvasWidth)

	ctx.FillStyle = colorAlizarin
	ctx.FillText("あえいうお", canvasWidth/2, height*8, canvasWidth)

	ctx.FillStyle = colorClouds
	ctx.FillText("あえいうお", canvasWidth/2, height*9, canvasWidth)

	ctx.FillStyle = colorConcrete
	ctx.FillText("あえいうお", canvasWidth/2, height*10, canvasWidth)
}

// KeyPressed handles key presses for the attract screen.
func (state AttractScreen) KeyPressed(keyCode int, pressed bool) {
	SetCurrentGameScreen(SettingsScreen{})
}
