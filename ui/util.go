package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func DrawCenteredText(screen *ebiten.Image, f text.Face, s string, cx, cy int) {
	tw, th := text.Measure(s, f, 6)
	x := float64(cx) - tw/float64(2)
	y := float64(cy) - th/float64(2)

	opt := text.DrawOptions{}
	opt.GeoM.Translate(float64(x), float64(y))
	text.Draw(screen, s, f, &opt)
}
