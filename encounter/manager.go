package encounter

import (
	"fishgame/environment"
	"fishgame/player"
	"fishgame/util"

	"github.com/hajimehoshi/ebiten/v2"
)

type Manager struct {
	Current *Encounter
	player  *player.Player
	Ended   bool
}

func NewManager(env *environment.Env, player *player.Player) *Manager {
	manager := &Manager{
		player: player,
		Ended:  false,
	}
	// generate levels as linked list of encounters
	startingScene := &Encounter{
		manager:  manager,
		Type:     EncounterTypeBattle,
		bg:       util.LoadImage(*env, "assets/bg/initial.png"),
		Behavior: NewBattleEncounter(env, "first battle"),
		player:   player,
	}
	manager.Current = startingScene

	return manager
}

func (em *Manager) SetCurrent(enc *Encounter) {
	em.Current = enc
}

func (em *Manager) Update(dt float64) {
	em.Current.Update(dt)
}

func (em *Manager) Draw(screen *ebiten.Image) {
	em.Current.Draw(screen)
}
