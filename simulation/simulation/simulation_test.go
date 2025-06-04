package simulation

import (
	"fishgame/shared/environment"
	"fishgame/simulation/collection"
	"fishgame/simulation/player"
	"testing"
)

// [UnitOfWork_StateUnderTest_ExpectedBehaviour]
// https://osherove.com/blog/2005/4/3/naming-standards-for-unit-tests.html
func Test_NewSimulation_Default_SubscribesToAllEvents(t *testing.T) {
	env := environment.NewEnv(nil, nil)
	NewSimulation(env, nil, nil)

	subscriptions := env.EventBus.GetSubscribers("StartSimulationEvent")
	if len(subscriptions) < 1 {
		t.Error("No StartSimulationEvent subscriptions registered")
	}
	subscriptions = env.EventBus.GetSubscribers("StopSimulationEvent")
	if len(subscriptions) < 1 {
		t.Error("No StopSimulationEvent subscriptions registered")
	}
}

func Test_StartSimulationEventHandler_WithEvent_DisablesCollectionChanges(t *testing.T) { // todo test inverse
	env := environment.NewEnv(nil, nil)
	sim := NewSimulation(env, &player.Player{
		Name: "player1",
		Fish: collection.NewCollection(env),
	}, collection.NewCollection(env))

	env.EventBus.Publish(environment.Event{
		Type: "StartSimulationEvent",
	})

	if sim.Player_GetFish().IsChangeable() {
		t.Error("Player collection should not be changeable")
	}
	if sim.Encounter_GetFish().IsChangeable() {
		t.Error("Encounter collection should not be changeable")
	}
}
