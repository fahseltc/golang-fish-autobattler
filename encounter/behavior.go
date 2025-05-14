package encounter

import (
	"fishgame/item"
	"fishgame/player"

	"github.com/hajimehoshi/ebiten/v2"
)

type BehaviorInterface interface {
	Update(float64, *player.Player)
	Draw(*ebiten.Image)
	GetItems() *item.Collection
	IsDone() bool
}
