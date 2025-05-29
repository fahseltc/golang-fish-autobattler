package ui

import (
	"fishgame/environment"
	"fishgame/item"

	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Slot struct {
	env   *environment.Env
	index int

	x      int
	y      int
	height int
	width  int
	item   *item.Item
}

func NewPlayerSlot(env *environment.Env, index int) *Slot {
	screenWidth := env.Get("screenWidth").(int)
	screenHeight := env.Get("screenHeight").(int)
	spriteSizePx := float64(env.Get("spriteSizePx").(int))
	spriteScale := env.Get("spriteScale").(float64)

	slotX := int(float64(screenWidth) * 0.4)

	slotYSpacingFromTop := float64(screenHeight) * float64(0.1)
	slotY := int(slotYSpacingFromTop + (spriteSizePx*spriteScale)*float64(index))

	slot := Slot{
		env:    env,
		index:  index,
		x:      slotX,
		y:      slotY,
		height: int(float64(spriteSizePx) * spriteScale),
		width:  int(float64(spriteSizePx) * spriteScale),
	}
	return &slot
}

func NewEncounterSlot(env *environment.Env, playerNum int, index int) *Slot {
	screenWidth := env.Get("screenWidth").(int)
	screenHeight := env.Get("screenHeight").(int)
	spriteSizePx := float64(env.Get("spriteSizePx").(int))
	spriteScale := env.Get("spriteScale").(float64)

	slotX := int(float64(screenWidth) * 0.6)

	slotYSpacingFromTop := float64(screenHeight) * float64(0.1)
	slotY := int(slotYSpacingFromTop + (spriteSizePx*spriteScale)*float64(index))

	slot := Slot{
		env:    env,
		index:  index,
		x:      slotX,
		y:      slotY,
		height: int(float64(spriteSizePx) * spriteScale),
		width:  int(float64(spriteSizePx) * spriteScale),
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

func (slot *Slot) IsEmpty() bool {
	return slot.item == nil
}

// func (slot *Slot) Swap(index int, newItem *item.Item) (*item.Item, error) {
// 	if slot.IsEmpty() {
// 		return nil, fmt.Errorf("item slot with index: %v is empty; unable to swap", index)
// 	}
// 	oldItem := slot.item
// 	slot.item = newItem
// 	return oldItem, nil
// }

func (slot *Slot) Update() {}

func DrawLifeBar(screen *ebiten.Image, healthRatio float64, x, y float64) {
	healthLength := float64((float64(spriteSizePx) * spriteScale) * healthRatio)
	ebitenutil.DrawRect(screen, x, y+4, healthLength, 6, color.RGBA{0, 255, 0, 255})
}

func DrawProgressBar(screen *ebiten.Image, progressRatio float64, x, y float64) {
	progressLength := float64((float64(spriteSizePx) * spriteScale) * progressRatio)
	ebitenutil.DrawRect(screen, x, y+(float64(spriteSizePx)*spriteScale)-8, progressLength, 4, color.White)
}

func (slot *Slot) DrawTooltip(screen *ebiten.Image, ui *UI, mx int, my int, playerNum int) {
	mb := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	coll := slot.Collides(mx, my)
	if !mb && coll.Collides && slot.item != nil {
		var ttx float32
		if playerNum == 1 {
			ttx = float32(slot.x) - float32(slot.height)
		} else {
			ttx = float32(slot.x) + float32(slot.height)
		}

		tty := float32(slot.y)

		vector.DrawFilledRect(screen, ttx, tty, float32(slot.width), float32(slot.width), color.RGBA{128, 128, 128, 255}, true)

		titleX := ttx + float32(slot.height)*0.5
		titleY := tty + float32(slot.width)*0.15
		DrawCenteredText(screen, ui.smallFont, slot.item.Name, int(titleX), int(titleY), nil)

		ttstring := fmt.Sprintf("DPS: %.2f", slot.item.Dps())
		DrawCenteredText(screen, ui.smallFont, ttstring, int(titleX), int(titleY+15), nil)

		hpstring := fmt.Sprintf("HP: %v/%v", slot.item.CurrentLife, slot.item.Life)
		DrawCenteredText(screen, ui.smallFont, hpstring, int(titleX), int(titleY+30), nil)

		DrawCenteredText(screen, ui.smallFont, slot.item.Description, int(titleX), int(titleY+45), nil)
	}
}

func (slot *Slot) Collides(x int, y int) Collision {
	collidesSlot := x > slot.x && x < slot.x+slot.width && y > slot.y && y < slot.y+slot.height
	if collidesSlot {
		//fmt.Printf("point (%v, %v) collides with slot at (%v, %v): %v\n", x, y, slot.x, slot.y, collidesSlot)
		if slot.CollidesBottomHalf(x, y) {
			return Collision{
				Type:     CollisionBottomHalf,
				Collides: true,
			}
		} else if slot.CollidesTopHalf(x, y) {
			return Collision{
				Type:     CollisionTopHalf,
				Collides: true,
			}
		}
	}
	return Collision{
		Type:     CollisionNone,
		Collides: false,
	}
}

func (slot *Slot) CollidesTopHalf(x, y int) bool {
	collides := x > slot.x && x < slot.x+slot.width && y > slot.y && float32(y) <= float32(slot.y)+(float32(0.5)*float32(slot.height))

	// if collides {
	// 	fmt.Printf("point (%v, %v) CollidesTopHalf at (%v, %v)\n", x, y, slot.x, slot.y)
	// }
	return collides
}

func (slot *Slot) CollidesBottomHalf(x, y int) bool {
	collides := x > slot.x && x < slot.x+slot.width && float32(y) > float32(slot.y)+float32(0.5)*float32(slot.height) && y < slot.y+slot.height

	// if collides {
	// 	fmt.Printf("point (%v, %v) CollidesBottomHalf at (%v, %v)\n", x, y, slot.x, slot.y)
	// }
	return collides
}
