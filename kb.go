package main

import (
	"github.com/oskca/gopherjs-canvas"
	"github.com/oskca/gopherjs-dom"
)

var (
	ctx          *canvas.Context2D
	canvasWidth  float64
	canvasHeight float64
)

func drawTriangle(ctx *canvas.Context2D) {
	ctx.Save()
	ctx.LineWidth = 1
	ctx.StrokeStyle = "black"
	ctx.FillStyle = "black"
	ctx.Translate(canvasWidth/2, canvasHeight/2)
	ctx.BeginPath()
	ctx.MoveTo(0, 0)
	ctx.LineTo(50, 50)
	ctx.LineTo(-50, 50)
	ctx.LineTo(0, 0)
	ctx.Fill()
	ctx.Stroke()
	ctx.Restore()
}

func main() {
	window := dom.Window()
	doc := window.Document
	cnvs := canvas.New(doc.GetElementById("canvas").Object)
	ctx = cnvs.GetContext2D()
	canvasWidth = float64(cnvs.Width)
	canvasHeight = float64(cnvs.Height)

	drawTriangle(ctx)
}
