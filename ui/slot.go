package ui

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/util"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Slot struct {
	env   *environment.Env
	index int

	sprite *ebiten.Image

	x      int
	y      int
	height int
	width  int
	item   *item.Item
}

func NewSlot(env *environment.Env, playerNum int, index int) *Slot {
	screenWidth := env.Get("width").(int)
	screenHeight := env.Get("height").(int)

	slotX := 0

	if playerNum == 1 {
		slotX = int(0.40 * float64(screenHeight))
	}
	if playerNum == 2 {
		slotX = int(0.60 * float64(screenWidth))
	}
	slotYSpacingFromTop := float64(screenHeight) * float64(0.2)
	slotY := int(slotYSpacingFromTop + 64*float64(index))

	sprite := util.LoadImage(*env, "assets/slot.png")

	slot := Slot{
		env:    env,
		index:  index,
		sprite: sprite,
		x:      slotX,
		y:      slotY,
		height: 64,
		width:  64,
	}
	return &slot
}

func (slot *Slot) AddItem(index int, it *item.Item) bool {
	if slot.item == nil { // only replace the item if its already empty
		slot.item = it
		return true
	}
	return false
}

func (slot *Slot) IsEmpty(index int) bool {
	return slot.item == nil
}

func (slot *Slot) Swap(index int, newItem *item.Item) (*item.Item, error) {
	if slot.IsEmpty(index) {
		return nil, fmt.Errorf("item slot with index: %v is empty; unable to swap", index)
	}
	oldItem := slot.item
	slot.item = newItem
	return oldItem, nil
}

func (slot *Slot) Update() {

	// if canDrop {
	// 	slot.Text.Label = "* Can Drop *"
	// 	if targetWidget != nil {
	// 		targetWidget.(*widget.Container).BackgroundImage = image.NewNineSliceColor(color.NRGBA{100, 100, 255, 255})
	// 		slot.TargetedObject = targetWidget
	// 	}
	// } else {
	// 	slot.Text.Label = "Cannot Drop"
	// 	if slot.TargetedObject != nil {
	// 		slot.TargetedObject.(*widget.Container).BackgroundImage = image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255})
	// 		slot.TargetedObject = nil
	// 	}
	// }
}

func DrawLifeBar(screen *ebiten.Image, healthRatio float64, x, y float64) {
	healthLength := float64(64 * healthRatio)
	ebitenutil.DrawRect(screen, x, y+4, healthLength, 6, color.RGBA{0, 255, 0, 255})
}

func DrawProgressBar(screen *ebiten.Image, progressRatio float64, x, y float64) {
	progressLength := float64(64 * progressRatio)
	ebitenutil.DrawRect(screen, x, y+58, progressLength, 4, color.White)
}

func (slot *Slot) Collides(x, y int) CollisionType {
	collidesSlot := x > slot.x && x < slot.x+slot.width && y > slot.y && y < slot.y+slot.height
	if collidesSlot {
		fmt.Printf("point (%v, %v) collides with slot at (%v, %v): %v\n", x, y, slot.x, slot.y, collidesSlot)
		if slot.CollidesBottomHalf(x, y) {
			return CollisionBottomHalf
		} else if slot.CollidesTopHalf(x, y) {
			return CollisionTopHalf
		} else {
			return CollisionNone
		}
	}
	return CollisionNone
}

func (slot *Slot) CollidesTopHalf(x, y int) bool {
	collides := x > slot.x && x < slot.x+slot.width && y > slot.y && float32(y) < float32(slot.y)+(float32(0.5)*float32(slot.height))

	if collides {
		fmt.Printf("point (%v, %v) CollidesTopHalf at (%v, %v)\n", x, y, slot.x, slot.y)
	}
	return collides
}

func (slot *Slot) CollidesBottomHalf(x, y int) bool {
	collides := x > slot.x && x < slot.x+slot.width && float32(y) > float32(slot.y)+float32(0.5)*float32(slot.height) && y < slot.y+slot.height

	if collides {
		fmt.Printf("point (%v, %v) CollidesBottomHalf at (%v, %v)\n", x, y, slot.x, slot.y)
	}
	return collides
}
