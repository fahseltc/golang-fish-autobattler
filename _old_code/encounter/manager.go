package encounter

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/player"
	"fishgame/ui"
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type EncounterInterface interface {
	Update(float64, *player.Player)
	Draw(*ebiten.Image)
	GetItems() *item.Collection

	IsDone() bool
	IsStarted() bool
	IsGameOver() bool

	// GetRewards() []*reward.Reward
	// AddReward(*reward.Reward)
	// SetRewards([]*reward.Reward)
}

type Manager struct {
	Current    EncounterInterface
	Encounters [][]EncounterInterface
	player     *player.Player
	Ended      bool
	ui         *ui.UI

	currentTierIndex      int
	currentEncounterIndex int
}

var ENV *environment.Env

func NewManager(env *environment.Env, player *player.Player, ui *ui.UI) *Manager {
	ENV = env
	manager := &Manager{
		player:                player,
		ui:                    ui,
		Ended:                 false,
		currentTierIndex:      0,
		currentEncounterIndex: 0,
	}
	manager.Encounters = [][]EncounterInterface{
		make([]EncounterInterface, 1), // initial encounters tier with 1 slot
		make([]EncounterInterface, 3), // t1 encounters tier with 3 slots (adjust as needed)
		make([]EncounterInterface, 1), // t2 encounters tier with 3 slots (adjust as needed)
	}
	// Load initial encounters from JSON and assign to the first tier
	initialEnc := LoadEncounters("data/encounters/initial_encounters.json", player, manager)
	copy(manager.Encounters[0], initialEnc)
	manager.Current = manager.Encounters[0][0] // start with the only initial encounter, perhaps add more in the future?

	// Load tier 1 encounters from JSON and assign to the second tier
	t1Enc := LoadEncounters("data/encounters/t1_encounters.json", player, manager)
	copy(manager.Encounters[1], t1Enc)

	t2Enc := LoadEncounters("data/encounters/t2_encounters.json", player, manager)
	copy(manager.Encounters[2], t2Enc)

	return manager
}

func (manager *Manager) SetCurrent(enc EncounterInterface) {
	manager.Current = enc
}

func (mgr *Manager) NextEncounter() EncounterInterface {
	var nextEnc EncounterInterface
	mgr.currentTierIndex += 1
	nextEnc = mgr.getRandomEncounterForTier()

	ENV.Logger.Info("NextEncounter")

	// ENV.Logger.Info("NextEncounter",
	// 	"previous",
	// 	fmt.Sprintf("T%v:%v", mgr.currentTierIndex-1, mgr.Current.GetType().String()),
	// 	"next",
	// 	fmt.Sprintf("T%v:%v", mgr.currentTierIndex, nextEnc.GetType().String()),
	// 	"rewards",
	// 	nextEnc.GetRewards(),
	// )
	mgr.Current = nextEnc
	// for index, it := range mgr.Current.GetItems().ActiveItems {
	// 	mgr.ui.Player2Slots[index].AddItem(it)
	// }

	return nil
}

func (mgr *Manager) getRandomEncounterForTier() EncounterInterface {
	if mgr.currentTierIndex >= len(mgr.Encounters) {
		log.Fatalf("No encounters for this tier: %v", mgr.currentTierIndex)
	}
	encCount := len(mgr.Encounters[mgr.currentTierIndex])
	if encCount == 1 {
		return mgr.Encounters[mgr.currentTierIndex][0]
	}
	rnd := rand.IntN(encCount - 1)
	return mgr.Encounters[mgr.currentTierIndex][rnd]
}
