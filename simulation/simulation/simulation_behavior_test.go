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
	return player.NewPlayer(ENV, "testplayer")
}

func Test_Update_WithBasicWeapon_HitEachOther(t *testing.T) {
	player := setupPlayer()
	stats := fish.NewWeaponStats(20, 1, 5)
	player.Fish.AddFish(fish.NewFish(ENV, "Whale", "he big", &stats), 0)

	enemyItems := collection.NewCollection(ENV)
	stats2 := fish.NewWeaponStats(20, 1, 5)
	enemyItems.AddFish(fish.NewFish(ENV, "Goldfish", "he little", &stats2), 0)
	sim := NewSimulation(ENV, player)
	sim.Encounter_SetFish(enemyItems)
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
	stats := fish.NewStats(fish.SizeBasedWeapon, fish.SizeHuge, 20, 2, 5)
	player.Fish.AddFish(fish.NewFish(ENV, "Whale", "he big", &stats), 0)

	enemyItems := collection.NewCollection(ENV)
	stats2 := fish.NewWeaponStats(10, 1, 1)
	enemyItems.AddFish(fish.NewFish(ENV, "Goldfish", "he little", &stats2), 0)
	sim := NewSimulation(ENV, player)
	sim.Encounter_SetFish(enemyItems)
	sim.Enable()
	sim.Update(2.0)

	if sim.Encounter_GetFish().GetAllFish()[0].IsAlive() {
		t.Error("Encounter Goldfish should be dead to double damage from whale")
	}
	if sim.Player_GetFish().GetAllFish()[0].Stats.CurrentLife != 20 {
		t.Error("Player whale should be at full life because dead goldfish tell no tales (dont attack)")
	}
}

func Test_Update_WithVenomBasedWeapon_DoesDamageOverTimeToTarget_AndStacks(t *testing.T) {
	player := setupPlayer()
	stats := fish.NewStats(fish.VenomousBasedWeapon, fish.SizeLarge, 20, 1, 5)
	venomFish := fish.NewFish(ENV, "Poisonous", "he ouch", &stats)
	player.Fish.AddFish(venomFish, 0)

	encounterFish := collection.NewCollection(ENV)
	stats2 := fish.NewWeaponStats(100, 999, 999)
	encounterFish.AddFish(fish.NewFish(ENV, "Goldfish", "he little", &stats2), 0)
	sim := NewSimulation(ENV, player)
	sim.Encounter_SetFish(encounterFish)
	sim.Enable()
	sim.Update(1.0)

	targetFish := sim.Encounter_GetFish().GetAllFish()[0]
	if targetFish.Stats.CurrentLife != 95 {
		t.Error("Encounter fish was not hurt for one application of venom")
	}
	if len(targetFish.Debuffs) != 1 {
		t.Error("Encounter fish should only have one debuff")
	}
	sim.Update(1.0)
	// two applications of debuff are applied and one expired + was removed
	if sim.Encounter_GetFish().GetAllFish()[0].Stats.CurrentLife != 85 {
		t.Error("Encounter fish was not hurt for second application of venom")
	}
	if len(targetFish.Debuffs) != 1 {
		t.Error("Encounter fish should only have one debuff")
	}
	venomFish.Stats.MaxDuration = 999 // make venom fish stop
	sim.Update(1)                     // tick
	if sim.Encounter_GetFish().GetAllFish()[0].Stats.CurrentLife != 80 {
		t.Error("Encounter fish was not hurt for third application of venom")
	}
	if len(targetFish.Debuffs) != 0 {
		t.Error("Encounter fish should have zero debuffs")
	}
}
