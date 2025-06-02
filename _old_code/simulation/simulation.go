package simulation

import (
	"fishgame/encounter"
	"fishgame/environment"
	"fishgame/item"
	"fishgame/player"
)

type Simulation struct {
	player    *player.Player
	encounter encounter.EncounterInterface

	enabled bool
}

type SimulationInterface interface {
	Update(float64)
	GetPlayerItems() *item.Collection
	GetPlayerInventory() *item.Inventory
	GetEncounterItems() *item.Collection
	// MovePlayerItem(*Slot, *Slot) slot object but thats in UI for now and we dont want that?
	StorePlayerItem(*item.Item)
	// GetItemFromInventory()
}

// this should have all of the game logic in it and nothing external
// ui fires events here saying thing "run game", "move fish", etc...
// remove as many places as possible to remove invalid state.
// top level sim determines adjacency bonus + debuffs

// API
// Get Fish
// Move Fish (slot-to-slot or slot-to-inventory)
// Run Game
// Stop Game
// ...

var ENV *environment.Env

func NewSimulation(env *environment.Env, player *player.Player, encounter encounter.EncounterInterface) SimulationInterface {
	ENV = env

	sim := &Simulation{
		player:    player,
		encounter: encounter,
		enabled:   false,
	}

	ENV.EventBus.Subscribe("StartSimulation", sim.StartSimulationEventHandler)
	ENV.EventBus.Subscribe("StopSimulation", sim.StopSimulationEventHandler)

	return sim
}

func (sim *Simulation) Update(dt float64) {
	sim.encounter.Update(dt, sim.player)
	if sim.enabled {
		// calc DT?
		sim.player.Items.Update(dt, sim.GetEncounterItems())
		sim.encounter.GetItems().Update(dt, sim.GetPlayerItems())
	}
}

func (sim *Simulation) Disable() {
	sim.enabled = false
}
func (sim *Simulation) GetPlayerItems() *item.Collection {
	return sim.player.Items
}
func (sim *Simulation) GetPlayerInventory() *item.Inventory {
	return sim.player.Inventory
}
func (sim *Simulation) GetEncounterItems() *item.Collection {
	return sim.encounter.GetItems()
}
func (sim *Simulation) StorePlayerItem(it *item.Item) {
	sim.player.Inventory.AddItem(it)
}

// Event Handlers
func (sim *Simulation) StartSimulationEventHandler(event environment.Event) {
	sim.enabled = true
}

func (sim *Simulation) StopSimulationEventHandler(event environment.Event) {
	sim.enabled = false
}
