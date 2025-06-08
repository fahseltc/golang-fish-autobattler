package ui

import (
	"fishgame/simulation/fish"
	"fishgame/ui/shapes"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ProgressBar struct {
	rect    *shapes.Rectangle
	stats   *fish.Stats
	offsetY float32
}

func NewProgressBar(rect *shapes.Rectangle, stats *fish.Stats) *ProgressBar {
	spriteSizePx := float64(ENV.Get("sprite.sizeInPx").(int))
	spriteScale := ENV.Get("sprite.scale").(float64)
	offset := float32((spriteScale * spriteSizePx)) - 10
	return &ProgressBar{
		rect:    rect,
		stats:   stats,
		offsetY: offset,
	}
}

func (pb *ProgressBar) Draw(screen *ebiten.Image) {
	ratio := float32(pb.stats.CurrentDuration) / float32(pb.stats.MaxDuration)
	ebitenutil.DrawRect(screen, float64(pb.rect.X), float64(pb.rect.Y+pb.offsetY), float64(ratio*pb.rect.W), 6, color.RGBA{122, 122, 122, 255})
}
