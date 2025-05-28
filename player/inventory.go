package player

import (
	"fishgame/item"
	"fishgame/util"

	"github.com/hajimehoshi/ebiten/v2"
)

type Inventory struct {
	Player *Player
	Items  map[string]*item.Item
	bg     *ebiten.Image
}

func NewInventory(player *Player) *Inventory {
	return &Inventory{
		Player: player,
		Items:  make(map[string]*item.Item),
		bg:     util.LoadImage(nil, "assets/pond.png"),
	}

}

func (i *Inventory) AddItem(it *item.Item) {
	i.Items[it.Id.String()] = it
}

func (i *Inventory) RemoveItem(it *item.Item) {
	delete(i.Items, it.Id.String())
}

func (i *Inventory) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, 80)
	screen.DrawImage(i.bg, opts)
}
