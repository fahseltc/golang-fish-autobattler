package player

import (
	"fishgame-sim/collection"
	"fishgame-sim/environment"
	"fishgame-sim/inventory"
)

type Player struct {
	env       *environment.Env
	Name      string
	Fish      *collection.Collection
	Inventory *inventory.Inventory
	Currency  int
}

func NewPlayer(env *environment.Env, name string) *Player {
	return &Player{
		env:       env,
		Name:      name,
		Fish:      collection.NewCollection(env),
		Inventory: inventory.NewInventory(),
		Currency:  0,
	}
}
