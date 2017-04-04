package main

import (
	"github.com/oskca/gopherjs-canvas"
	"github.com/oskca/gopherjs-dom"
)

const (
	stateAttract  = 0
	stateSettings = 1
	statePlaying  = 2
	stateGameOver = 3
)

var (
	ctx          *canvas.Context2D
	canvasWidth  float64
	canvasHeight float64
	gameState    = stateAttract
)

func drawAttract() {
	ctx.FillStyle = "black"
	ctx.Font = "48px serif"
	ctx.FillText("Title", canvasWidth/2, canvasHeight/2, canvasWidth)
}

func drawSettings() {
	ctx.FillStyle = "black"
	ctx.Font = "48px serif"
	ctx.FillText("Settings", canvasWidth/2, canvasHeight/2, canvasWidth)
}

func drawPlaying() {
	ctx.FillStyle = "black"
	ctx.Font = "48px serif"
	ctx.FillText("Playing", canvasWidth/2, canvasHeight/2, canvasWidth)
}

func drawGameOver() {
	ctx.FillStyle = "black"
	ctx.Font = "48px serif"
	ctx.FillText("Game Over", canvasWidth/2, canvasHeight/2, canvasWidth)
}

func update() {

}

func keyPressed(keyCode int, pressed bool) {
	switch gameState {
	case stateAttract:
		gameState = stateSettings
	case stateSettings:
		gameState = statePlaying
	case statePlaying:
		gameState = stateGameOver
	case stateGameOver:
		gameState = stateSettings
	}
}

func draw() {
	ctx.FillStyle = "white"
	ctx.FillRect(0.0, 0.0, canvasWidth, canvasHeight)

	switch gameState {
	case stateAttract:
		drawAttract()
	case stateSettings:
		drawSettings()
	case statePlaying:
		drawPlaying()
	case stateGameOver:
		drawGameOver()
	}
}

func gameLoop() {
	update()
	draw()

	dom.Window().Call("setTimeout", gameLoop, 1000/30)
}

func main() {
	window := dom.Window()
	doc := window.Document
	cnvs := canvas.New(doc.GetElementById("canvas").Object)
	ctx = cnvs.GetContext2D()
	canvasWidth = float64(cnvs.Width)
	canvasHeight = float64(cnvs.Height)

	window.AddEventListener(dom.EvtKeypress, func(event *dom.Event) {
		keyPressed(event.KeyCode, true)
	})

	gameLoop()
}
