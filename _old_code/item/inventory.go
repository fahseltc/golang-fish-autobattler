package item

import (
	"fishgame/shapes"
	"fishgame/util"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Inventory struct {
	Rect  shapes.Rectangle
	Items map[string]*Item
	bg    *ebiten.Image
}

func NewInventory() *Inventory {
	return &Inventory{
		Rect: shapes.Rectangle{
			X: 0,
			Y: 80,
			W: 333,
			H: 500,
		},
		Items: make(map[string]*Item),
		bg:    util.LoadImage(nil, "assets/pond.png"),
	}
}

func (i *Inventory) AddItem(it *Item) {
	i.Items[it.Id.String()] = it
}

func (i *Inventory) RemoveItem(it *Item) {
	delete(i.Items, it.Id.String())
}

func (i *Inventory) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, 80)
	screen.DrawImage(i.bg, opts)

	for _, item := range i.Items {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(item.X), float64(item.Y)) // get random pos inside inventory?
		screen.DrawImage(item.Sprite, op)
	}
}

func (i *Inventory) Collides(x, y int) bool {
	collides := x > int(i.Rect.X) && x < int(i.Rect.X+i.Rect.W) && y > int(i.Rect.Y) && y < int(i.Rect.Y+i.Rect.H)
	return collides
}
func (i *Inventory) GetRandomPos() (int, int) {
	x := rand.IntN(int(i.Rect.W-60)) + 25
	y := rand.IntN(int(i.Rect.H-60)) + 25
	return x, y
}
