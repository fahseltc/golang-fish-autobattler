package ui

import (
	"fishgame/simulation/fish"
	images "fishgame/ui/images"
	myshapes "fishgame/ui/shapes"
	"fishgame/ui/util"
	"fmt"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Id       *uuid.UUID
	Rect     myshapes.Rectangle
	Img      *ebiten.Image
	Dragging bool

	defaultXpos int

	previousX int
	previousY int
}

func NewInventorySprite(imageRegistry *images.Registry) *Sprite {
	w := 400
	h := 520
	rect := myshapes.Rectangle{
		X: 0,
		Y: 50,
		W: float32(w),
		H: float32(h),
	}

	img := imageRegistry.Images["pond.png"]
	if img == nil {
		img = imageRegistry.Images["TEXTURE_MISSING.png"]
	}
	scaled := util.ScaleImage(img, float32(w), float32(h))

	return &Sprite{
		Img:  scaled,
		Rect: rect,
	}
}

func NewPlayerFishSprite(imageRegistry *images.Registry, fish *fish.Fish, slotIndex int) *Sprite {
	return newFishSprite(imageRegistry, fish, slotIndex, true)
}

func NewEncounterFishSprite(imageRegistry *images.Registry, fish *fish.Fish, slotIndex int) *Sprite {
	return newFishSprite(imageRegistry, fish, slotIndex, false)
}

func newFishSprite(imageRegistry *images.Registry, fish *fish.Fish, slotIndex int, leftSide bool) *Sprite {
	spriteScale := ENV.Config.Get("sprite.scale").(float64)
	img := imageRegistry.Images[fmt.Sprintf("fish/%v.png", fish.Name)]
	if img == nil {
		img = imageRegistry.Images["TEXTURE_MISSING.png"]
	}
	w, h := img.Size()
	scaled := ebiten.NewImage(int(float64(w)*spriteScale), int(float64(h)*spriteScale))
	op := &ebiten.DrawImageOptions{} // Draw original onto the new image with scaling
	if leftSide {
		op.GeoM.Scale(spriteScale*-1, spriteScale)
		op.GeoM.Translate(float64(w)*spriteScale, 0)
	} else {
		op.GeoM.Scale(spriteScale, spriteScale)
	}

	scaled.DrawImage(img, op)

	yPadding := ENV.Config.Get("slot.topPad").(int)
	var xPos float32
	if leftSide {
		xPos = float32(ENV.Config.Get("slot.playerColX").(int))
	} else {
		xPos = float32(ENV.Config.Get("slot.encounterColX").(int))
	}
	betweenSlotPadding := ENV.Config.Get("slot.betweenPad").(int)
	w, h = scaled.Size()
	rect := myshapes.Rectangle{X: xPos, Y: (float32(slotIndex) * (float32(h + betweenSlotPadding))) + float32(yPadding), W: float32(w), H: float32(h)}

	sprite := &Sprite{
		Id:          &fish.Id,
		Rect:        rect,
		Img:         scaled,
		Dragging:    false,
		defaultXpos: int(xPos),
	}
	return sprite
}

func (spr *Sprite) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(spr.Rect.X), float64(spr.Rect.Y))
	screen.DrawImage(spr.Img, opts)

	// Draw a plain rectangle with the dimensions of spr.Rect
	//ebitenutil.DrawRect(screen, float64(spr.Rect.X), float64(spr.Rect.Y), float64(spr.Rect.W), float64(spr.Rect.H), color.RGBA{155, 155, 155, 155})
}

// Todo: some funny fish flopping/rotation animation while dragging?
func (spr *Sprite) MoveCentered(mx, my int) {
	spr.Rect.X = float32(mx) - spr.Rect.H/2
	spr.Rect.Y = float32(my) - spr.Rect.W/2
}

func (spr *Sprite) SetPosition(slotIndex int) {
	betweenSlotPadding := ENV.Config.Get("slot.betweenPad").(int)
	yPadding := ENV.Config.Get("slot.topPad").(int)
	_, h := spr.Img.Size()
	spr.Rect.Y = (float32(slotIndex)*(float32(float32(h)+float32(betweenSlotPadding))) + float32(yPadding))
	spr.Rect.X = float32(spr.defaultXpos)
}

func (spr *Sprite) SavePositionBeforeDrag() {
	spr.previousX = int(spr.Rect.X)
	spr.previousY = int(spr.Rect.Y)
}

func (spr *Sprite) ResetToPositionBeforeDrag() {
	spr.Rect.X = float32(spr.previousX)
	spr.Rect.Y = float32(spr.previousY)
}
