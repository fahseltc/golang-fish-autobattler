package simulation

import (
	"fishgame/shared/environment"
	"fishgame/simulation/collection"
	"fishgame/simulation/fish"
	"fishgame/simulation/inventory"
	"fishgame/simulation/player"
	"fmt"
	"time"
)

type Simulation struct {
	player       *player.Player
	enemyFish    *collection.Collection
	fishRegistry *fish.FishStatsRegistry

	enabled bool
}

type SimulationInterface interface {
	Update(float64)
	Enable()
	Disable()
	IsEnabled() bool
	IsInitialized() bool
	Player_Get() *player.Player
	Player_GetFish() *collection.Collection
	Player_GetInventory() *inventory.Inventory
	Player_GetAllStoredFish() []*fish.Fish
	Player_StoreNewFish(it *fish.Fish)
	Player_StoreExistingFish(id string) error
	Player_GetStoredFish(id string)
	Encounter_GetFish() *collection.Collection
	Encounter_SetFish(*collection.Collection)
	// MovePlayerItem(*Slot, *Slot) slot object but thats in UI for now and we dont want that?
	GetFishByID(id string) (int, *fish.Fish)
	IsPlayerFish(id string) bool
	IsEncounterFish(id string) bool
	IsGameOver() bool
	IsDone() bool
	GetFishRegistry() *fish.FishStatsRegistry
}

// this should have all of the game logic in it and nothing external
// ui fires events here saying thing "run game", "move fish", etc...
// remove as many places as possible to remove invalid state.
// top level sim determines adjacency bonus + debuffs

var ENV *environment.Env

func NewSimulation(env *environment.Env, player *player.Player) SimulationInterface {

	ENV = env

	sim := &Simulation{
		player:       player,
		enemyFish:    collection.NewCollection(ENV),
		fishRegistry: fish.NewFishStatsRegistry(ENV),
		enabled:      false,
	}

	ENV.EventBus.Subscribe("StartSimulationEvent", sim.startSimulationEventHandler)
	ENV.EventBus.Subscribe("StopSimulationEvent", sim.stopSimulationEventHandler)

	return sim
}

func (sim *Simulation) Update(dt float64) {
	if sim.enabled && sim.IsInitialized() {
		// calc DT?
		sim.player.Fish.Update(dt, sim.Encounter_GetFish())
		sim.enemyFish.Update(dt, sim.Player_GetFish())

		// Check if the simulation should stop
		sim.IsDone()
		sim.IsGameOver()
	}
}

func (sim *Simulation) Enable() {
	sim.enabled = true
	sim.player.Fish.DisableChanges()
	sim.enemyFish.DisableChanges()
}
func (sim *Simulation) Disable() {
	sim.enabled = false
	sim.player.Fish.EnableChanges()
	sim.enemyFish.EnableChanges()
}
func (sim *Simulation) IsEnabled() bool {
	return sim.enabled
}
func (sim *Simulation) IsInitialized() bool {
	// enemyFishPresent := false
	// for _, fish := range sim.enemyFish.GetAllFish() {
	// 	if fish != nil {
	// 		enemyFishPresent = true
	// 	}
	// }
	// playerFishPresent := false
	// for _, fish := range sim.player.Fish.GetAllFish() {
	// 	if fish != nil {
	// 		playerFishPresent = true
	// 	}
	// }
	return true
}

func (sim *Simulation) Player_Get() *player.Player {
	return sim.player
}

func (sim *Simulation) Player_GetFish() *collection.Collection {
	return sim.player.Fish
}

func (sim *Simulation) Player_GetInventory() *inventory.Inventory {
	return sim.player.Inventory
}
func (sim *Simulation) Player_GetAllStoredFish() []*fish.Fish {
	return sim.player.Inventory.GetAll()
}
func (sim *Simulation) Player_StoreNewFish(it *fish.Fish) {
	sim.player.Inventory.Add(it)
}
func (sim *Simulation) Player_StoreExistingFish(id string) error {
	var fishToMove *fish.Fish
	for _, fish := range sim.Player_GetFish().GetAllFish() {
		if fish != nil {
			if fish.Id.String() == id {
				success := sim.Player_GetFish().RemoveFish(id)
				if success {
					fishToMove = fish
				} else {
					return fmt.Errorf("unable to remove fish with ID:%v from player collection", id)
				}
			}
		}
	}
	if fishToMove == nil {
		return fmt.Errorf("no fish found with ID:%v", id)
	}
	sim.Player_StoreNewFish(fishToMove)
	return nil // success
}
func (sim *Simulation) Player_GetStoredFish(id string) {
	sim.player.Inventory.Get(id)
}

func (sim *Simulation) Encounter_GetFish() *collection.Collection {
	return sim.enemyFish
}
func (sim *Simulation) Encounter_SetFish(encounterFish *collection.Collection) {
	sim.enemyFish = encounterFish
}

func (sim *Simulation) IsGameOver() bool {
	playerFishDead := sim.Player_GetFish().AllFishDead()
	playerInventoryFish := sim.Player_GetInventory().GetCount()
	isGameOver := playerFishDead && playerInventoryFish <= 0
	if isGameOver {
		sim.Disable()
		ENV.EventBus.Publish(environment.Event{
			Type:      "GameOverEvent",
			Timestamp: time.Now(),
			// do we need data in there? who killed you maybe?
		})
	}
	return isGameOver
}
func (sim *Simulation) IsDone() bool {
	encounterDone := sim.Encounter_GetFish().AllFishDead()
	if encounterDone {
		sim.Disable()
		ENV.EventBus.Publish(environment.Event{
			Type:      "EncounterDoneEvent",
			Timestamp: time.Now(),
			// do we need data in there
		})
	}
	return encounterDone
}
func (sim *Simulation) GetFishByID(id string) (int, *fish.Fish) {
	index, fish := sim.Player_GetFish().ById(id)
	if fish == nil {
		index, fish = sim.Encounter_GetFish().ById(id)
	}
	if fish == nil {
		index = 999
		fish = sim.Player_GetInventory().Get((id))
	}
	return index, fish
}
func (sim *Simulation) IsPlayerFish(id string) bool {
	_, fish := sim.Player_GetFish().ById(id)
	return fish != nil
}
func (sim *Simulation) IsEncounterFish(id string) bool {
	_, fish := sim.Encounter_GetFish().ById(id)
	return fish != nil
}
func (sim *Simulation) GetFishRegistry() *fish.FishStatsRegistry {
	return sim.fishRegistry
}

// Event Handlers
func (sim *Simulation) startSimulationEventHandler(event environment.Event) {
	encounterFish := sim.Encounter_GetFish()
	anyPlayerFishPresent := sim.Player_GetFish().AnyFishPresent() // todo: check this for encounters as well?
	if encounterFish != nil && !encounterFish.AllFishDead() && anyPlayerFishPresent {
		sim.enabled = true
		sim.Player_GetFish().DisableChanges()
		sim.Encounter_GetFish().DisableChanges()
	} else {
		ENV.Logger.Warn("simulation failed to start, all encounter fish are dead or no player fish present in collection")
	}
}

func (sim *Simulation) stopSimulationEventHandler(event environment.Event) {
	sim.enabled = false
	sim.Player_GetFish().EnableChanges()
	sim.Encounter_GetFish().EnableChanges()
}
