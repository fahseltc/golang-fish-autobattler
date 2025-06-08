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
	player    *player.Player
	enemyFish *collection.Collection

	enabled bool
}

type SimulationInterface interface {
	Update(float64)
	Enable()
	Disable()
	IsEnabled() bool
	Player_GetFish() *collection.Collection
	Encounter_GetFish() *collection.Collection
	Player_GetInventory() *inventory.Inventory
	Player_GetAllStoredFish() []*fish.Fish
	Player_StoreNewFish(it *fish.Fish)
	Player_StoreExistingFish(id string) error
	Player_GetStoredFish(id string)
	// MovePlayerItem(*Slot, *Slot) slot object but thats in UI for now and we dont want that?
	GetFishByID(id string) (int, *fish.Fish)
	IsPlayerFish(id string) bool
	IsEncounterFish(id string) bool
	IsGameOver() bool
	IsDone() bool
}

// this should have all of the game logic in it and nothing external
// ui fires events here saying thing "run game", "move fish", etc...
// remove as many places as possible to remove invalid state.
// top level sim determines adjacency bonus + debuffs

var ENV *environment.Env

func NewSimulation(env *environment.Env, player *player.Player, enemyFish *collection.Collection) SimulationInterface {

	ENV = env

	sim := &Simulation{
		player:    player,
		enemyFish: enemyFish,
		enabled:   false,
	}
	fmt.Printf("SIM Constructor--ENV UUID:%v \n", env.UUID.String())

	ENV.EventBus.Subscribe("StartSimulationEvent", sim.startSimulationEventHandler)
	ENV.EventBus.Subscribe("StopSimulationEvent", sim.stopSimulationEventHandler)

	return sim
}

func (sim *Simulation) Update(dt float64) {
	//sim.encounter.Update(dt, sim.player)
	if sim.enabled {
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

func (sim *Simulation) Player_GetFish() *collection.Collection {
	return sim.player.Fish
}
func (sim *Simulation) Encounter_GetFish() *collection.Collection {
	return sim.enemyFish
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
					return fmt.Errorf("Unable to remove fish with ID:%v from player collection", id)
				}
			}
		}
	}
	if fishToMove == nil {
		return fmt.Errorf("No fish found with ID:%v", id)
	}
	sim.Player_StoreNewFish(fishToMove)
	return nil // success
}

func (sim *Simulation) Player_GetStoredFish(id string) {
	sim.player.Inventory.Get(id)
}
func (sim *Simulation) IsGameOver() bool {
	gameOver := sim.Player_GetFish().AllFishDead()
	if gameOver {
		sim.Disable()
		ENV.EventBus.Publish(environment.Event{
			Type:      "GameOverEvent",
			Timestamp: time.Now(),
			// do we need data in there? who killed you maybe?
		})
	}
	return gameOver
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

// Event Handlers
func (sim *Simulation) startSimulationEventHandler(event environment.Event) {
	sim.enabled = true
	sim.Player_GetFish().DisableChanges()
	sim.Encounter_GetFish().DisableChanges()
}

func (sim *Simulation) stopSimulationEventHandler(event environment.Event) {
	sim.enabled = false
	sim.Player_GetFish().EnableChanges()
	sim.Encounter_GetFish().DisableChanges()
}
