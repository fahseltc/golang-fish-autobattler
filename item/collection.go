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
}

func NewCollection(env *environment.Env, playerNum int) *Collection {
	coll := Collection{
		env: env,
	}

	return &coll
}

func (ic *Collection) Update(dt float64, enemyItems *Collection) {
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

func (coll *Collection) Draw(env environment.Env, screen *ebiten.Image, player int) {

	screenWidth := env.Get("width").(int)
	screenHeight := env.Get("height").(int)

	spriteX := 0.0
	spriteYSpacingFromTop := float64(screenHeight) * float64(0.2)
	if player == 1 {
		spriteX = 0.40 * float64(screenHeight)
	}
	if player == 2 {
		spriteX = 0.60 * float64(screenWidth)
	}
	for index, item := range coll.ActiveItems {
		spriteY := spriteYSpacingFromTop + 64*float64(index)
		if item != nil {
			op := &ebiten.DrawImageOptions{}
			if player == 1 {
				op.GeoM.Scale(-1, 1)                                     // flip the image horizontally for player 1
				op.GeoM.Translate(float64(item.Sprite.Bounds().Dx()), 0) // translate the image to the right
			}
			op.GeoM.Translate(spriteX, spriteY)
			item.X = int(spriteX)
			item.Y = int(spriteY)

			screen.DrawImage(item.Sprite, op)
		}
	}

	// for index, item := range coll.InactiveItems {
	// 	if item != nil {
	// 		op := &ebiten.DrawImageOptions{}
	// 		op.GeoM.Translate(100, 100+(64*float64(index)))
	// 		screen.DrawImage(item.Sprite, op)
	// 	}
	// }
}
