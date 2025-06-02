package simulation

import (
	"fishgame-sim/collection"
	"fishgame-sim/environment"
	"fishgame-sim/fish"
	"fishgame-sim/inventory"
	"fishgame-sim/player"
	"fmt"
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
	Player_GetFish() *collection.Collection
	Encounter_GetFish() *collection.Collection
	Player_GetInventory() *inventory.Inventory
	Player_GetAllStoredFish() []*fish.Fish
	Player_StoreNewFish(it *fish.Fish)
	Player_StoreExistingFish(id string) bool
	Player_GetStoredFish(id string)
	// MovePlayerItem(*Slot, *Slot) slot object but thats in UI for now and we dont want that?
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

	ENV.EventBus.Subscribe("StartSimulation", sim.StartSimulationEventHandler)
	ENV.EventBus.Subscribe("StopSimulation", sim.StopSimulationEventHandler)

	return sim
}

func (sim *Simulation) Update(dt float64) {
	//sim.encounter.Update(dt, sim.player)
	if sim.enabled {
		// calc DT?
		sim.player.Fish.Update(dt, sim.Encounter_GetFish())
		sim.enemyFish.Update(dt, sim.Player_GetFish())
	}
}

func (sim *Simulation) Enable() {
	sim.enabled = true
}
func (sim *Simulation) Disable() {
	sim.enabled = false
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

// Event Handlers
func (sim *Simulation) StartSimulationEventHandler(event environment.Event) {
	sim.enabled = true
}

func (sim *Simulation) StopSimulationEventHandler(event environment.Event) {
	sim.enabled = false
}
