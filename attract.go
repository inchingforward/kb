package main

import (
	"math"
	"math/rand"

	canvas "github.com/oskca/gopherjs-canvas"
)

// AttractScreen is the first screen displayed to the player.
type AttractScreen struct{}

type glyph struct {
	x, y, dy float64
	kana     rune
}

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

var (
	glyphs []glyph
)

func randomRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Setup initializes various variables used in the Attract screen.
func (screen AttractScreen) Setup() {
	characters := "あえいうおまめみむもかけきくこ"

	glyphs = make([]glyph, len(characters))

	for i, k := range characters {
		glyphs[i] = glyph{float64(i * 14), randomRange(0, canvasHeight-1), randomRange(0.5, 2.5), k}
	}
}

// Update handles object updates for the attract screen.
func (screen AttractScreen) Update(ctx *canvas.Context2D) {
	for i := range glyphs {
		glyphs[i].y = math.Mod(canvasHeight+glyphs[i].y+glyphs[i].dy, canvasHeight)
	}
}

// Draw handles drawing for the attract screen.
func (screen AttractScreen) Draw(ctx *canvas.Context2D) {
	ctx.Font = "36px Noto Sans Japanese"
	ctx.FillStyle = colorClouds

	for _, g := range glyphs {
		ctx.FillText(string(g.kana), g.x, g.y, canvasWidth)
	}

	ctx.FillStyle = colorPeterRiver
	ctx.TextAlign = "center"
	ctx.FillText("Kana Blast", canvasWidth/2, canvasHeight/2, canvasWidth)

	ctx.Font = "14px sans-serif"
	ctx.FillStyle = colorCarrot
	ctx.FillText("Press Any Key To Start", canvasWidth/2, canvasHeight/2+28, canvasWidth)
}

// KeyPressed handles key presses for the attract screen.
func (screen AttractScreen) KeyPressed(keyCode int, pressed bool) {
	SetCurrentGameScreen(SettingsScreen{})
}
