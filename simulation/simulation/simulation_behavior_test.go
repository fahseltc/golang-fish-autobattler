package simulation

import (
	"fishgame-sim/collection"
	"fishgame-sim/environment"
	"fishgame-sim/fish"
	"fishgame-sim/player"
	"testing"
)

func Test_Update_WithBasicWeapon_HitEachOther(t *testing.T) {
	env := environment.NewEnv(nil, nil)
	player := &player.Player{
		Name: "player1",
		Fish: collection.NewCollection(env),
	}
	player.Fish.AddFish(fish.NewFish("Whale", "he big", fish.NewWeaponStats(20, 1, 5)), 0)

	enemyItems := collection.NewCollection(env)
	enemyItems.AddFish(fish.NewFish("Goldfish", "he little", fish.NewWeaponStats(20, 1, 5)), 0)
	sim := NewSimulation(env, player, enemyItems)
	sim.Enable()
	sim.Update(1.0)

	if sim.Player_GetFish().GetAllFish()[0].Stats.CurrentLife != 15 {
		t.Error("Player fish was not hurt for one tick duration")
	}
	if sim.Encounter_GetFish().GetAllFish()[0].Stats.CurrentLife != 15 {
		t.Error("Enemy fish was not hurt for one tick duration")
	}
}

func Test_Update_WithSizeBasedWeapon_DoesDoubleDamageToTarget(t *testing.T) {
	env := environment.NewEnv(nil, nil)
	player := &player.Player{
		Name: "player1",
		Fish: collection.NewCollection(env),
	}
	player.Fish.AddFish(fish.NewFish("Whale", "he big", fish.NewStats(fish.SizeBasedWeapon, fish.SizeHuge, 20, 2, 5)), 0)

	enemyItems := collection.NewCollection(env)
	enemyItems.AddFish(fish.NewFish("Goldfish", "he little", fish.NewWeaponStats(10, 1, 5)), 0)
	sim := NewSimulation(env, player, enemyItems)
	sim.Enable()
	sim.Update(2.0)

	if sim.Encounter_GetFish().GetAllFish()[0].IsAlive() {
		t.Error("Encounter Goldfish should be dead to double damage from whale")
	}
	if sim.Player_GetFish().GetAllFish()[0].Stats.CurrentLife != 20 {
		t.Error("Player whale should be at full life because dead goldfish tell no tales (dont attack)")
	}
}
