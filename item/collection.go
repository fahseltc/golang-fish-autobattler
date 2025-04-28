package item

import (
	"fishgame/environment"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Collection struct {
	ActiveItems   []*Item
	InactiveItems []*Item
}

func (ic *Collection) Update(dt float32, enemyItems *Collection) {
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
	spriteX := 0.0
	spriteY := 100.0
	if player == 1 {
		spriteX = 100.0
	}
	if player == 2 {
		spriteX = 500.0
	}
	for index, item := range coll.ActiveItems {
		spriteY = 100 + 64*float64(index)
		if item != nil {
			op := &ebiten.DrawImageOptions{}
			if player == 1 {
				op.GeoM.Scale(-1, 1)                                     // flip the image horizontally for player 1
				op.GeoM.Translate(float64(item.Sprite.Bounds().Dx()), 0) // translate the image to the right
			}
			op.GeoM.Translate(spriteX, spriteY)

			screen.DrawImage(item.Sprite, op)

			// draw life bar above each sprite
			healthRatio := float64(item.CurrentLife) / float64(item.Life)
			healthLength := float64(64 * healthRatio)
			ebitenutil.DrawRect(screen, spriteX, spriteY+64, healthLength, 4, color.White)
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
