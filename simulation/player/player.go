package player

import (
	"fishgame-sim/collection"
	"fishgame-sim/inventory"
)

type Player struct {
	//env   *environment.Env
	Name      string
	Fish      *collection.Collection
	Inventory *inventory.Inventory
	Currency  int
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:      name,
		Fish:      collection.NewCollection(),
		Inventory: inventory.NewInventory(),
		Currency:  0,
	}
}
