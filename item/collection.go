package item

import (
	"fishgame/environment"
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var SlotCount int = 5

type Collection struct {
	env           *environment.Env
	ActiveItems   []*Item
	InactiveItems []*Item
	spriteScale   float64
}

func NewEmptyPlayerCollection(env *environment.Env) *Collection {
	spriteScale := env.Get("spriteScale").(float64)

	coll := &Collection{
		env:         env,
		spriteScale: spriteScale,
	}
	for i := 0; i < SlotCount; i++ {
		coll.ActiveItems = append(coll.ActiveItems, nil)
		coll.InactiveItems = append(coll.InactiveItems, nil)
	}

	return coll
}

func NewEncounterCollection(env *environment.Env, items []*Item) *Collection {
	coll := &Collection{
		env:         env,
		ActiveItems: items,
		spriteScale: env.Get("spriteScale").(float64),
	}
	coll.SetItemLocations()
	return coll
}
func (coll *Collection) SetItemLocations() {
	screenWidth := coll.env.Get("screenWidth").(int)
	screenHeight := coll.env.Get("screenHeight").(int)
	spriteSizePx := float64(coll.env.Get("spriteSizePx").(int))
	spriteScale := coll.env.Get("spriteScale").(float64)

	spriteX := 0.60 * float64(screenWidth)
	spriteYSpacingFromTop := float64(screenHeight) * float64(0.1)
	for index, item := range coll.ActiveItems {
		spriteY := int(spriteYSpacingFromTop + (spriteSizePx*spriteScale)*float64(index))
		if item != nil {
			item.X = int(spriteX)
			item.Y = int(spriteY)
		}
	}

}

func (ic *Collection) Update(dt float64, enemyItems *Collection) {
	// assume enemyItems will be not-nil
	for index, item := range ic.ActiveItems {
		if item != nil && !item.Alive {
			// remove item from active items and add to inactive items
			ic.ActiveItems = append(ic.ActiveItems[:index], ic.ActiveItems[index+1:]...)
			ic.InactiveItems = append(ic.InactiveItems, item)
		}
		item.Update(dt, enemyItems, ic, index)
	}
}

func (coll *Collection) Reset() {
	coll.ActiveItems = append(coll.ActiveItems, coll.InactiveItems...)
	coll.InactiveItems = []*Item{}
}

func (coll *Collection) GetRandomActive() (int, *Item) {
	var validIndexes []int

	for idx, it := range coll.ActiveItems {
		if it != nil {
			validIndexes = append(validIndexes, idx)
		}
	}
	if len(validIndexes) == 0 {
		return 999, nil
	}

	randomIndex := rand.Intn(len(validIndexes))
	randomItem := coll.ActiveItems[validIndexes[randomIndex]]
	if randomItem == nil || !randomItem.Alive {
		return 999, nil
	}

	return randomIndex, coll.ActiveItems[randomIndex]
}

func (coll *Collection) Draw(env *environment.Env, screen *ebiten.Image, player int) {
	for _, item := range coll.ActiveItems {
		if item != nil && item.Alive {
			op := &ebiten.DrawImageOptions{}
			if player == 1 {
				op.GeoM.Scale(-1, 1)                                     // flip the image horizontally for player 1
				op.GeoM.Translate(float64(item.Sprite.Bounds().Dx()), 0) // translate the image to the right
			}
			op.GeoM.Translate(float64(item.X), float64(item.Y))
			screen.DrawImage(item.Sprite, op)

			// Draw Debuffs
			for _, dbf := range item.debuffs {
				dbf.Draw(screen)
			}
		}

	}

}
func (coll *Collection) AddItem(it *Item) bool {
	if coll.SlotsFull() {
		return false
	}
	// if len(coll.ActiveItems) >= SlotCount {
	// 	return false
	// }

	emptyIndex := coll.FirstEmptyIndex()
	fmt.Printf("AddItem - emptyIndex: %v, itemName: %v\n", emptyIndex, it.Name)
	coll.ActiveItems[emptyIndex] = it
	it.SlotIndex = emptyIndex
	coll.setItemSprite(it)
	return true
}

func (coll *Collection) AddItems(items []*Item) bool {
	for _, it := range items {
		res := coll.AddItem(it)
		if !res {
			return false
		}
	}
	return true
}

func (coll *Collection) RemoveItem(it *Item) bool { // todo implement
	return true
}

func (coll *Collection) setItemSprite(item *Item) {
	screenWidth := coll.env.Get("screenWidth").(int)
	screenHeight := coll.env.Get("screenHeight").(int)
	spriteSizePx := float64(coll.env.Get("spriteSizePx").(int))
	spriteScale := coll.env.Get("spriteScale").(float64)

	spriteX := int(float64(screenWidth) * 0.4)
	spriteYSpacingFromTop := float64(screenHeight) * float64(0.1)

	itemIndex := item.SlotIndex
	spriteY := int(spriteYSpacingFromTop + (spriteSizePx*spriteScale)*float64(itemIndex))
	item.X = int(spriteX)
	item.Y = int(spriteY)
}

func (coll *Collection) GetItemAbove(index int) *Item {
	fmt.Printf("GetItemAbove - itemIndex: %v\n", index)
	if index == 0 || coll.ActiveItems[index-1] == nil {
		return nil
	}
	fmt.Printf("GetItemAbove-notnil - itemName: %v\n", coll.ActiveItems[index-1].Name)
	return coll.ActiveItems[index-1]
}
func (coll *Collection) GetItemBelow(index int) *Item {
	fmt.Printf("GetItemBelow - itemIndex: %v\n", index)
	if index+1 >= SlotCount || coll.ActiveItems[index+1] == nil {
		return nil
	}
	fmt.Printf("GetItemAbove-notnil - itemName: %v\n", coll.ActiveItems[index+1].Name)
	return coll.ActiveItems[index+1]
}

func (coll *Collection) SlotsFull() bool {
	for _, it := range coll.ActiveItems {
		if it == nil {
			return false
		}
	}
	return true
}

func (coll *Collection) FirstEmptyIndex() int {
	for index, it := range coll.ActiveItems {
		if it == nil {
			fmt.Printf("EmptyIndex? - index: %v, item NIL RETURNING IT\n", index)
			return index
		} else {
			fmt.Printf("EmptyIndex? - index: %v, item: %v\n", index, it)
		}
	}
	return 999
}
