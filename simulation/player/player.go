package player

import (
	"fishgame/shared/environment"
	"fishgame/simulation/collection"
	"fishgame/simulation/inventory"
)

type Player struct {
	env       *environment.Env
	Name      string
	Fish      *collection.Collection
	Inventory *inventory.Inventory
	currency  int
}

func NewPlayer(env *environment.Env, name string) *Player {
	return &Player{
		env:       env,
		Name:      name,
		Fish:      collection.NewCollection(env),
		Inventory: inventory.NewInventory(),
		currency:  0,
	}
}
func (p *Player) GetCurrencyAmount() int {
	return p.currency
}
func (p *Player) SpendCurrency(amt int) bool {
	valid := (p.currency - amt) >= 0

	if valid {
		p.currency = p.currency - amt
		return true
	} else {
		return false
	}
}

func (p *Player) AddCurrency(amt int) {
	p.currency = p.currency + amt
}
