package ui

import (
	"fishgame/simulation/fish"
	"fishgame/ui/shapes"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ProgressBar struct {
	maxVal     int
	currentVal int
	rect       shapes.Rectangle
	countUp    bool
	stats      *fish.Stats
}

func NewCountDownProgressBar(max int, rect shapes.Rectangle, stats *fish.Stats) *ProgressBar {
	return &ProgressBar{
		maxVal:     max,
		currentVal: max,
		rect:       rect,
		countUp:    false,
		stats:      stats,
	}
}

func NewCountUpProgressBar(max int, rect shapes.Rectangle) *ProgressBar {
	return &ProgressBar{
		maxVal:     max,
		currentVal: 0,
		rect:       rect,
		countUp:    true,
	}
}

func (pb *ProgressBar) Draw(screen *ebiten.Image) {
	var ratio float64

	if pb.countUp {
		ratio = float64(pb.currentVal) / float64(pb.maxVal)
	} else {
		ratio = 1 - float64(pb.currentVal)/float64(pb.maxVal)
	}
	healthLength := float64(pb.rect.W) * ratio
	ebitenutil.DrawRect(screen, float64(pb.rect.X), float64(pb.rect.X), healthLength, 6, color.RGBA{0, 255, 0, 255})
}

// healthLength := float64((float64(spriteSizePx) * spriteScale) * healthRatio)
// 	ebitenutil.DrawRect(screen, x, y+4, healthLength, 6, color.RGBA{0, 255, 0, 255})
