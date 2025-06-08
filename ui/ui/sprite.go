package ui

import (
	"fishgame/simulation/fish"
	images "fishgame/ui/images"
	myshapes "fishgame/ui/shapes"
	"fishgame/ui/util"
	"fmt"
	"image/color"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Id       *uuid.UUID
	Rect     myshapes.Rectangle
	Img      *ebiten.Image
	Dragging bool
	toolTip  TooltipInterface

	defaultXpos int

	previousX int
	previousY int

	healthBar   *HealthBar
	progressBar *ProgressBar
}

func NewInventorySprite(imageRegistry *images.Registry) *Sprite {
	w := 340
	h := 520
	rect := myshapes.Rectangle{
		X: 30,
		Y: 35,
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
	playerSprite := newFishSprite(imageRegistry, fish, slotIndex, true)
	playerSprite.toolTip = NewFishToolTip(ENV, playerSprite.Rect, LeftAlignment, fish)
	playerSprite.healthBar = NewHealthProgressBar(&playerSprite.Rect, fish.Stats)
	playerSprite.progressBar = NewProgressBar(&playerSprite.Rect, fish.Stats)
	return playerSprite
}

func NewEncounterFishSprite(imageRegistry *images.Registry, fish *fish.Fish, slotIndex int) *Sprite {
	encounterSprite := newFishSprite(imageRegistry, fish, slotIndex, false)
	encounterSprite.toolTip = NewFishToolTip(ENV, encounterSprite.Rect, LeftAlignment, fish)
	encounterSprite.healthBar = NewHealthProgressBar(&encounterSprite.Rect, fish.Stats)
	encounterSprite.progressBar = NewProgressBar(&encounterSprite.Rect, fish.Stats)
	return encounterSprite
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
		//healthBar: *NewCountDownProgressBar()
	}
	return sprite
}

func (spr *Sprite) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(spr.Rect.X), float64(spr.Rect.Y))
	screen.DrawImage(spr.Img, opts)
	if spr.healthBar != nil {
		spr.healthBar.Draw(screen)
	}
	if spr.progressBar != nil {
		spr.progressBar.Draw(screen)
	}

	if ENV.Config.Get("debugDraw").(bool) {
		ebitenutil.DrawRect(screen, float64(spr.Rect.X), float64(spr.Rect.Y), float64(spr.Rect.W), float64(spr.Rect.H), color.RGBA{155, 155, 155, 155})
	}
}

// func DrawLifeBar(screen *ebiten.Image, healthRatio float64, x, y float64) {
// 	healthLength := float64((float64(spriteSizePx) * spriteScale) * healthRatio)
// 	ebitenutil.DrawRect(screen, x, y+4, healthLength, 6, color.RGBA{0, 255, 0, 255})
// }

// func DrawProgressBar(screen *ebiten.Image, progressRatio float64, x, y float64) {
// 	progressLength := float64((float64(spriteSizePx) * spriteScale) * progressRatio)
// 	ebitenutil.DrawRect(screen, x, y+(float64(spriteSizePx)*spriteScale)-8, progressLength, 4, color.White)
// }

func (spr *Sprite) DrawToolTip(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()
	if spr.toolTip != nil && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && spr.Rect.Collides(float32(mx), float32(my)) {
		spr.toolTip.ReAlign(spr)
		spr.toolTip.OnHover(screen)
	}
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
