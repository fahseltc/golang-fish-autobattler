package simulation

import (
	"fishgame/shared/environment"
	"fishgame/simulation/collection"
	"fishgame/simulation/fish"
	"fishgame/simulation/player"
	"testing"
)

func setupPlayer() *player.Player {
	ENV = environment.NewEnv(nil, nil)
	return &player.Player{
		Name: "player1",
		Fish: collection.NewCollection(ENV),
	}
}

func Test_Update_WithBasicWeapon_HitEachOther(t *testing.T) {
	player := setupPlayer()
	player.Fish.AddFish(fish.NewFish(ENV, "Whale", "he big", fish.NewWeaponStats(20, 1, 5)), 0)

	enemyItems := collection.NewCollection(ENV)
	enemyItems.AddFish(fish.NewFish(ENV, "Goldfish", "he little", fish.NewWeaponStats(20, 1, 5)), 0)
	sim := NewSimulation(ENV, player, enemyItems)
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
	player := setupPlayer()
	player.Fish.AddFish(fish.NewFish(ENV, "Whale", "he big", fish.NewStats(fish.SizeBasedWeapon, fish.SizeHuge, 20, 2, 5)), 0)

	enemyItems := collection.NewCollection(ENV)
	enemyItems.AddFish(fish.NewFish(ENV, "Goldfish", "he little", fish.NewWeaponStats(10, 1, 5)), 0)
	sim := NewSimulation(ENV, player, enemyItems)
	sim.Enable()
	sim.Update(2.0)

	if sim.Encounter_GetFish().GetAllFish()[0].IsAlive() {
		t.Error("Encounter Goldfish should be dead to double damage from whale")
	}
	if sim.Player_GetFish().GetAllFish()[0].Stats.CurrentLife != 20 {
		t.Error("Player whale should be at full life because dead goldfish tell no tales (dont attack)")
	}
}

func Test_Update_WithVenomBasedWeapon_DoesDamageOverTimeToTarget(t *testing.T) {
	env := environment.NewEnv(nil, nil)
	player := &player.Player{
		Name: "player1",
		Fish: collection.NewCollection(env),
	}
	player.Fish.AddFish(fish.NewFish(ENV, "Poisonous", "he ouch", fish.NewStats(fish.VenomousBasedWeapon, fish.SizeLarge, 20, 1, 5)), 0)

	encounterFish := collection.NewCollection(env)
	encounterFish.AddFish(fish.NewFish(ENV, "Goldfish", "he little", fish.NewWeaponStats(100, 999, 999)), 0)
	sim := NewSimulation(env, player, encounterFish)
	sim.Enable()
	sim.Update(1.0)

	if sim.Encounter_GetFish().GetAllFish()[0].Stats.CurrentLife != 95 {
		t.Error("Encounter fish was not hurt for one application of venom")
	}
	sim.Update(1.0)
	if sim.Encounter_GetFish().GetAllFish()[0].Stats.CurrentLife != 90 {
		t.Error("Encounter fish was not hurt for second application of venom")
	}
	sim.Update(1.0)
	if sim.Encounter_GetFish().GetAllFish()[0].Stats.CurrentLife != 85 {
		t.Error("Encounter fish was not hurt for third application of venom")
	}
}
