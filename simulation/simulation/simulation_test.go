package simulation

import (
	"fishgame-sim/collection"
	"fishgame-sim/environment"
	"fishgame-sim/fish"
	"fishgame-sim/player"
	"testing"
)

func TestEventSubscriptions(t *testing.T) {
	env := environment.NewEnv(nil, nil)
	NewSimulation(env, nil, nil)

	subscriptions := env.EventBus.GetSubscribers("StartSimulation")
	if len(subscriptions) < 1 {
		t.Error("No StartSimulationEvent subscriptions registered")
	}
	subscriptions = env.EventBus.GetSubscribers("StopSimulation")
	if len(subscriptions) < 1 {
		t.Error("No StopSimulationEvent subscriptions registered")
	}
}

func TestSimpleBattle(t *testing.T) {
	env := environment.NewEnv(nil, nil)
	player := &player.Player{
		Name: "player1",
		Fish: collection.NewCollection(),
	}
	player.Fish.AddFish(fish.NewFish("Whale", "he big", fish.NewStats(20, 1, 5)), 0)

	enemyItems := collection.NewCollection()
	enemyItems.AddFish(fish.NewFish("Goldfish", "he little", fish.NewStats(20, 1, 5)), 0)
	sim := NewSimulation(env, player, enemyItems)
	sim.Enable()
	sim.Update(1.0)

	if sim.GetPlayerFish().FishSlotMap[0].Stats.CurrentLife != 15 {
		t.Error("Player fish was not hurt for one tick duration")
	}
	if sim.GetEncounterFish().FishSlotMap[0].Stats.CurrentLife != 15 {
		t.Error("Enemy fish was not hurt for one tick duration")
	}
}
