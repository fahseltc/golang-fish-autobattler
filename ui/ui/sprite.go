package ui

import (
	"fishgame/simulation/fish"
	"fishgame/ui/shapes"
	"fishgame/ui/util"
	"fmt"
	"image/color"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Id       *uuid.UUID
	Rect     shapes.Rectangle
	Img      *ebiten.Image
	Dragging bool
	toolTip  TooltipInterface

	defaultXpos int

	previousX int
	previousY int

	healthBar   *HealthBar
	progressBar *ProgressBar
}

func NewInventorySprite() *Sprite {
	rect := shapes.Rectangle{
		X: float32(ENV.Config.Get("inventory.x").(int)),
		Y: float32(ENV.Config.Get("inventory.y").(int)),
		W: float32(ENV.Config.Get("inventory.w").(int)),
		H: float32(ENV.Config.Get("inventory.h").(int)),
	}

	img := util.LoadImage("pond.png")
	if img == nil {
		img = util.LoadImage("TEXTURE_MISSING.png")
	}
	scaled := util.ScaleImage(img, rect.W, rect.H)

	return &Sprite{
		Img:  scaled,
		Rect: rect,
	}
}

func NewPlayerFishSprite(fish *fish.Fish, slotIndex int) *Sprite {
	playerSprite := newFishSprite(fish, slotIndex, true)
	playerSprite.toolTip = NewFishToolTip(ENV, playerSprite.Rect, shapes.LeftAlignment, fish)
	playerSprite.healthBar = NewHealthProgressBar(&playerSprite.Rect, fish.Stats)
	playerSprite.progressBar = NewProgressBar(&playerSprite.Rect, fish.Stats)
	return playerSprite
}

func NewEncounterFishSprite(fish *fish.Fish, slotIndex int) *Sprite {
	encounterSprite := newFishSprite(fish, slotIndex, false)
	encounterSprite.toolTip = NewFishToolTip(ENV, encounterSprite.Rect, shapes.LeftAlignment, fish)
	encounterSprite.healthBar = NewHealthProgressBar(&encounterSprite.Rect, fish.Stats)
	encounterSprite.progressBar = NewProgressBar(&encounterSprite.Rect, fish.Stats)
	return encounterSprite
}

func NewInventoryFishSprite(fish *fish.Fish) *Sprite {
	sprite := newFishSprite(fish, 1, true)

	// set the sprite to a position inside the inventory
	xPos, yPos := GetRandomInventoryPosition()
	sprite.Rect.X = float32(xPos)
	sprite.Rect.Y = float32(yPos)

	sprite.toolTip = NewFishToolTip(ENV, sprite.Rect, shapes.LeftAlignment, fish)
	sprite.healthBar = NewHealthProgressBar(&sprite.Rect, fish.Stats)
	sprite.progressBar = NewProgressBar(&sprite.Rect, fish.Stats)
	return sprite
}

func newFishSprite(fish *fish.Fish, slotIndex int, leftSide bool) *Sprite {
	spriteScale := ENV.Config.Get("sprite.scale").(float64)
	img := util.LoadImage(fmt.Sprintf("fish/%v.png", strings.ToLower(fish.Name)))
	if img == nil {
		img = util.LoadImage("TEXTURE_MISSING.png")
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
	rect := shapes.Rectangle{X: xPos, Y: (float32(slotIndex) * (float32(h + betweenSlotPadding))) + float32(yPadding), W: float32(w), H: float32(h)}

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

func GetRandomInventoryPosition() (x, y int) {
	padding := 200 // pixels of padding from the edge
	rect := shapes.Rectangle{
		X: float32(ENV.Config.Get("inventory.x").(int)),
		Y: float32(ENV.Config.Get("inventory.y").(int)),
		W: float32(ENV.Config.Get("inventory.w").(int)),
		H: float32(ENV.Config.Get("inventory.h").(int)),
	}
	rand.Seed(time.Now().UnixNano())
	minX := int(rect.X) + padding
	maxX := int(rect.X+rect.W) - padding
	minY := int(rect.Y) + padding
	maxY := int(rect.Y+rect.H) - padding
	if maxX <= minX {
		maxX = minX + 1
	}
	if maxY <= minY {
		maxY = minY + 1
	}
	x = rand.Intn(maxX-minX) + minX
	y = rand.Intn(maxY-minY) + minY
	return
}
