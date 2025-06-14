package ui

import (
	"fishgame/ui/shapes"
	"fishgame/ui/util"
	"image/color"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Slot struct {
	index   int
	rect    shapes.Rectangle
	itemId  *uuid.UUID
	slotImg *ebiten.Image
}

func newSlot(index int, xPos float32) *Slot {
	yPadding := ENV.Config.Get("slot.topPad").(int)
	betweenSlotPadding := ENV.Config.Get("slot.betweenPad").(int)
	spriteSizePx := ENV.Config.Get("sprite.sizeInPx").(int)
	spriteScale := ENV.Config.Get("sprite.scale").(float64)
	slotImg := util.LoadImage("slot.png")
	scaled := util.ScaleImage(slotImg, float32(float64(spriteSizePx)*spriteScale), float32(float64(spriteSizePx)*spriteScale))
	height := float64(spriteSizePx) * spriteScale
	yPos := float32(index)*float32(int(height)+betweenSlotPadding) + float32(yPadding)

	slot := Slot{
		index: index,
		rect: shapes.Rectangle{
			X: xPos,
			Y: yPos,
			W: float32(height), // square so its the same
			H: float32(height),
		},
		slotImg: scaled,
	}
	return &slot
}

func NewPlayerSlot(index int) *Slot {
	playerXPos := float32(ENV.Config.Get("slot.playerColX").(int))
	return newSlot(index, playerXPos)
}

func NewEncounterSlot(index int) *Slot {
	encounterXPos := float32(ENV.Config.Get("slot.encounterColX").(int))
	return newSlot(index, encounterXPos)
}

func (slot *Slot) SetSprite(spr *Sprite) bool {
	if slot.itemId == nil { // only replace the item if its already empty
		slot.itemId = spr.Id
		return true
	}
	return false
}

func (slot *Slot) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(slot.rect.X), float64(slot.rect.Y))
	screen.DrawImage(slot.slotImg, opts)

	if ENV.Config.Get("debugDraw").(bool) {
		ebitenutil.DrawRect(screen, float64(slot.rect.X), float64(slot.rect.Y), float64(slot.rect.W), float64(slot.rect.H), color.RGBA{200, 200, 155, 255})
	}
}
