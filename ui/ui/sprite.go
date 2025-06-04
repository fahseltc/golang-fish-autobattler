package ui

import (
	"fishgame/ui/shapes"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Id   *uuid.UUID
	rect shapes.Rectangle
	Img  *ebiten.Image
}

func NewSprite(id *uuid.UUID, rect shapes.Rectangle, img *ebiten.Image) *Sprite {
	// resize img to match sprite scale!
	sprite := &Sprite{
		Id:   id,
		rect: rect,
		Img:  img,
	}
	return sprite
}

func (spr *Sprite) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(spr.rect.X), float64(spr.rect.Y))
	screen.DrawImage(spr.Img, opts)
}
