package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func DrawCenteredText(screen *ebiten.Image, f text.Face, s string, cx, cy int, clr color.Color) {
	tw, th := text.Measure(s, f, 6)
	x := float64(cx) - tw/float64(2)
	y := float64(cy) - th/float64(2)

	var textColor color.Color
	if clr == nil {
		textColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	} else {
		textColor = clr
	}

	opt := text.DrawOptions{}
	opt.ColorScale.ScaleWithColor(textColor)
	opt.GeoM.Translate(float64(x), float64(y))
	text.Draw(screen, s, f, &opt)
}
