package simulation

import (
	"fishgame/shared/environment"
	"fishgame/simulation/collection"
	"fishgame/simulation/fish"
	"fishgame/simulation/player"
	"testing"
)

// [UnitOfWork_StateUnderTest_ExpectedBehaviour]
// https://osherove.com/blog/2005/4/3/naming-standards-for-unit-tests.html
func Test_NewSimulation_Default_SubscribesToAllEvents(t *testing.T) {
	env := environment.NewEnv(nil, nil)
	NewSimulation(env, nil)

	subscriptions := env.EventBus.GetSubscribers("StartSimulationEvent")
	if len(subscriptions) < 1 {
		t.Error("No StartSimulationEvent subscriptions registered")
	}
	subscriptions = env.EventBus.GetSubscribers("StopSimulationEvent")
	if len(subscriptions) < 1 {
		t.Error("No StopSimulationEvent subscriptions registered")
	}
}

func Test_StartSimulationEventHandler_WithEventAndEmptyEnemies_DoesntDisableCollectionChanges(t *testing.T) { // todo test inverse
	env := environment.NewEnv(nil, nil)
	sim := NewSimulation(env, &player.Player{
		Name: "player1",
		Fish: collection.NewCollection(env),
	})
	sim.Encounter_SetFish(collection.NewCollection(env))

	env.EventBus.Publish(environment.Event{
		Type: "StartSimulationEvent",
	})

	if !sim.Player_GetFish().IsChangeable() {
		t.Error("Player collection should be changeable")
	}
}

func Test_StartSimulationEventHandler_WithEventAndEnemies_DisablesCollectionChanges(t *testing.T) { // todo test inverse
	env := environment.NewEnv(nil, nil)
	sim := NewSimulation(env, &player.Player{
		Name: "player1",
		Fish: collection.NewCollection(env),
	})
	enemyFish := collection.NewCollection(env)
	enemyFish.AddFish(fish.NewFish(env, "testfish", "", fish.NewStats(fish.Weapon, fish.SizeSmall, 99, 1, 1)), 1)
	sim.Encounter_SetFish(enemyFish)

	env.EventBus.Publish(environment.Event{
		Type: "StartSimulationEvent",
	})

	if sim.Player_GetFish().IsChangeable() {
		t.Error("Player collection should not be changeable")
	}
}
