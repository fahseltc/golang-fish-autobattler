package encounter

import (
	"fishgame/shared/environment"
	"fishgame/simulation/collection"
	"fishgame/simulation/fish"
	"fmt"
	"log"
	"math/rand/v2"
)

type EncounterInterface interface {
	// Update(float64, *player.Player)
	// Draw(*ebiten.Image)
	GetCollection() *collection.Collection
	GetType() Type

	IsDone() bool
	IsStarted() bool
	IsGameOver() bool

	GetRewards() []*Reward
	GetTitle() string
	// AddReward(*reward.Reward)
	// SetRewards([]*reward.Reward)
}

type Manager struct {
	encounters [][]EncounterInterface

	currentTierIndex      int
	currentEncounterIndex int
}

var ENV *environment.Env

func NewEncounterManager(env *environment.Env, registry *fish.FishStatsRegistry) *Manager {
	ENV = env
	em := &Manager{
		currentTierIndex:      0,
		currentEncounterIndex: 0,
	}

	em.encounters = [][]EncounterInterface{
		make([]EncounterInterface, 0),
		make([]EncounterInterface, 0),
		make([]EncounterInterface, 0),
	}
	// set encounters
	jsonT0File := LoadEncounterFile("encounters/t0_encounters.json")
	for _, enc := range jsonT0File.Encounters {
		e := NewEncounter(enc, registry)
		em.encounters[0] = append(em.encounters[0], e)
	}
	jsonT1File := LoadEncounterFile("encounters/t1_encounters.json")
	for _, enc := range jsonT1File.Encounters {
		e := NewEncounter(enc, registry)
		em.encounters[1] = append(em.encounters[1], e)
	}

	ENV.EventBus.Subscribe("GameOverEvent", em.handleGameOverEvent)
	return em
}

func (mgr *Manager) GetRandomEncounterForTier() (EncounterInterface, error) {
	if mgr.currentTierIndex >= len(mgr.encounters) {
		return nil, fmt.Errorf("no encounters found for this tier: %v", mgr.currentTierIndex)
	}
	encCount := len(mgr.encounters[mgr.currentTierIndex])
	if encCount == 1 {
		return mgr.encounters[mgr.currentTierIndex][0], nil
	}
	if encCount == 0 {
		ENV.Logger.Error(fmt.Sprintf("there are no encounters for tier: %v", mgr.currentTierIndex))
		log.Fatal()
	}
	rnd := rand.IntN(encCount - 1)
	return mgr.encounters[mgr.currentTierIndex][rnd], nil
}

func (mgr *Manager) GetNext() (EncounterInterface, error) {
	mgr.currentTierIndex += 1
	enc, err := mgr.GetRandomEncounterForTier()
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func (em *Manager) GetCurrent() (EncounterInterface, error) {
	tier := em.encounters[em.currentTierIndex]
	if tier != nil {
		enc := tier[em.currentEncounterIndex]
		if enc != nil {
			return enc, nil
		} else {
			return nil, fmt.Errorf("unable to find encounter with tier:%v and encounter:%v", em.currentTierIndex, em.currentEncounterIndex)
		}
	} else {
		return nil, fmt.Errorf("unable to find encounter with tier:%v", em.currentTierIndex)
	}
}

func (em *Manager) handleGameOverEvent(event environment.Event) {
	em.currentEncounterIndex = 0
	em.currentTierIndex = 0
}
