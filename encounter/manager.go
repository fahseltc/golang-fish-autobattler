package encounter

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/player"
	"fishgame/ui"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Manager struct {
	Current           EncounterInterface
	TierOneEncounters []EncounterInterface
	player            *player.Player
	Ended             bool
	ui                *ui.UI
}

type EncounterInterface interface {
	Update(float64, *player.Player)
	Draw(*ebiten.Image)
	GetItems() *item.Collection
	IsDone() bool
	GetType() Type
}

func NewManager(env *environment.Env, player *player.Player, ui *ui.UI) *Manager {
	manager := &Manager{
		player: player,
		ui:     ui,
		Ended:  false,
	}

	initialChoice := LoadEncounters(env, "data/encounters/initial.json", player, manager)
	manager.Current = initialChoice[0] // only one initial

	t1enc := LoadEncounters(env, "data/encounters/t1.json", player, manager)
	manager.TierOneEncounters = t1enc
	return manager
}

func (manager *Manager) SetCurrent(enc EncounterInterface) {
	manager.Current = enc
}

func (manager *Manager) NextEncounter() EncounterInterface {
	fmt.Printf("NEXT ENCOUNTER\n")
	manager.setItemSlots()
	manager.Current = manager.TierOneEncounters[0]
	return nil
}

func (manager *Manager) setItemSlots() {
	for index, it := range manager.player.Items.ActiveItems {
		manager.ui.Player1Slots[index].AddItem(index, it)
		it.SlotIndex = index
	}
}
