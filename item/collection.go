package item

import (
	"fishgame/environment"
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
		if !item.Alive {
			// remove item from active items and add to inactive items
			ic.ActiveItems = append(ic.ActiveItems[:index], ic.ActiveItems[index+1:]...)
			ic.InactiveItems = append(ic.InactiveItems, item)
		}
		item.Update(dt, enemyItems)
	}
}

func (coll *Collection) Reset() {
	coll.ActiveItems = append(coll.ActiveItems, coll.InactiveItems...)
	coll.InactiveItems = []*Item{}
}

func (coll *Collection) GetRandomActive() (int, *Item) {
	if len(coll.ActiveItems) == 0 {
		return 0, nil
	}
	randomIndex := rand.Intn(len(coll.ActiveItems))
	randomItem := coll.ActiveItems[randomIndex]
	if randomItem == nil || !randomItem.Alive {
		return 0, nil
	}

	return randomIndex, coll.ActiveItems[randomIndex]
}

func (coll *Collection) Draw(env *environment.Env, screen *ebiten.Image, player int) {
	for _, item := range coll.ActiveItems {
		if item != nil {
			op := &ebiten.DrawImageOptions{}
			if player == 1 {
				op.GeoM.Scale(-1, 1)                                     // flip the image horizontally for player 1
				op.GeoM.Translate(float64(item.Sprite.Bounds().Dx()), 0) // translate the image to the right
			}
			op.GeoM.Translate(float64(item.X), float64(item.Y))
			screen.DrawImage(item.Sprite, op)
		}
	}

}
func (coll *Collection) AddItem(it *Item) bool {
	if len(coll.ActiveItems) >= SlotCount {
		return false
	}
	coll.setItemSprite(it)
	coll.ActiveItems = append(coll.ActiveItems, it)
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

	itemCount := len(coll.ActiveItems)
	spriteY := int(spriteYSpacingFromTop + (spriteSizePx*spriteScale)*float64(itemCount))
	item.X = int(spriteX)
	item.Y = int(spriteY)
}
