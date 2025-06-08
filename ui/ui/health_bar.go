package ui

import (
	"fishgame/simulation/fish"
	"fishgame/ui/shapes"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type HealthBar struct {
	rect  *shapes.Rectangle
	stats *fish.Stats
}

func NewHealthProgressBar(rect *shapes.Rectangle, stats *fish.Stats) *HealthBar {
	return &HealthBar{
		rect:  rect,
		stats: stats,
	}
}

// func NewCountUpProgressBar(max int, rect *shapes.Rectangle) *ProgressBar {
// 	return &ProgressBar{
// 		maxVal:     max,
// 		currentVal: 0,
// 		rect:       rect,
// 		countUp:    true,
// 	}
// }

func (pb *HealthBar) Draw(screen *ebiten.Image) {
	ratio := float32(pb.stats.CurrentLife) / float32(pb.stats.MaxLife)
	ebitenutil.DrawRect(screen, float64(pb.rect.X), float64(pb.rect.Y+4), float64(ratio*pb.rect.W), 6, color.RGBA{0, 255, 0, 255})
}
