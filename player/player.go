package player

import (
	"fishgame/environment"
	"fishgame/item"
)

type Player struct {
	Env      *environment.Env
	Name     string
	Items    *item.Collection
	Currency int
}
