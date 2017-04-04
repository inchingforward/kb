package main

import (
	"github.com/oskca/gopherjs-canvas"
	"github.com/oskca/gopherjs-dom"
)

var (
	ctx           *canvas.Context2D
	canvasWidth   float64
	canvasHeight  float64
	currentScreen GameScreen
)

// GameScreen represents the current screen the player is interacting with.
type GameScreen interface {
	Draw(ctx *canvas.Context2D)
	Update(ctx *canvas.Context2D)
	KeyPressed(keyCode int, pressed bool)
}

// SetCurrentGameScreen switches the screen to the given GameState.
func SetCurrentGameScreen(state GameScreen) {
	currentScreen = state
}

func gameLoop() {
	ctx.FillStyle = "white"
	ctx.FillRect(0.0, 0.0, canvasWidth, canvasHeight)

	currentScreen.Update(ctx)
	currentScreen.Draw(ctx)

	dom.Window().Call("setTimeout", gameLoop, 1000/30)
}

func main() {
	window := dom.Window()
	doc := window.Document
	cnvs := canvas.New(doc.GetElementById("canvas").Object)
	ctx = cnvs.GetContext2D()
	canvasWidth = float64(cnvs.Width)
	canvasHeight = float64(cnvs.Height)

	currentScreen = AttractScreen{}

	window.AddEventListener(dom.EvtKeypress, func(event *dom.Event) {
		currentScreen.KeyPressed(event.KeyCode, true)
	})

	gameLoop()
}
