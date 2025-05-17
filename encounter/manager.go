package encounter

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/player"
	"fishgame/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Manager struct {
	env               *environment.Env
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
	IsGameOver() bool
	GetType() Type
}

func NewManager(env *environment.Env, player *player.Player, ui *ui.UI) *Manager {
	manager := &Manager{
		env:    env,
		player: player,
		ui:     ui,
		Ended:  false,
	}

	initialChoice := LoadEncounters(env, "data/encounters/initial_encounters.json", player, manager)
	manager.Current = initialChoice[0] // only one initial

	t1enc := LoadEncounters(env, "data/encounters/t1_encounters.json", player, manager)
	manager.TierOneEncounters = t1enc
	return manager
}

func (manager *Manager) SetCurrent(enc EncounterInterface) {
	manager.Current = enc
}

func (manager *Manager) NextEncounter() EncounterInterface {
	manager.env.Logger.Info("NextEncounter", "previous", manager.Current.GetType().String(), "next", manager.TierOneEncounters[0].GetType().String())
	manager.setItemSlots()
	manager.Current = manager.TierOneEncounters[0] // just pick first one for now!
	for index, it := range manager.Current.GetItems().ActiveItems {
		manager.ui.Player2Slots[index].AddItem(index, it)
	}

	return nil
}

func (manager *Manager) setItemSlots() {
	for index, it := range manager.player.Items.ActiveItems {
		manager.ui.Player1Slots[index].AddItem(index, it)
		it.SlotIndex = index
		//fmt.Printf("item: %v going into slot: %v successfully: %v\n", it.Name, index, added)
	}
}
