package encounter

import (
	"fishgame/player"

	"github.com/hajimehoshi/ebiten/v2"
)

type Encounter struct {
	manager        *Manager
	player         *player.Player
	Type           Type
	bg             *ebiten.Image
	Behavior       BehaviorInterface
	NextEncounters []*Encounter
}

func (enc *Encounter) Update(dt float64) {
	if enc.Behavior != nil {
		enc.Behavior.Update(dt, enc.player)
	}
}

func (enc *Encounter) Draw(screen *ebiten.Image) {
	screen.DrawImage(enc.bg, nil)
	if enc.Behavior != nil {
		enc.Behavior.Draw(screen)
	}
}
