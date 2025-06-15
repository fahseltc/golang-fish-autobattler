package fish_test

import (
	"fishgame/shared/environment"
	"fishgame/simulation/fish"
	"slices"
	"testing"
)

var ENV *environment.Env

func setup() (source *fish.Fish, target *fish.Fish) {
	ENV = environment.NewEnv(nil, nil)
	stats := fish.NewStats(fish.Weapon, fish.SizeSmall, 100, 1, 5)
	source = fish.NewFish(ENV, "sourcefish", "", &stats)
	stats2 := fish.NewStats(fish.Weapon, fish.SizeSmall, 100, 1, 5)
	target = fish.NewFish(ENV, "targetfish", "", &stats2)
	return source, target
}

func Test_SoloAttackingBehavior_WithNoNeighbors_DealsDoubleDamage(t *testing.T) {
	source, target := setup()

	source.Stats.Type = fish.SoloBasedWeapon
	sourceIndex := 2
	fishSlice := make([]*fish.Fish, 4)
	fishSlice = slices.Insert(fishSlice, sourceIndex, source)
	fish.SoloAttackingBehavior(source, target, sourceIndex, fishSlice)

	if target.Stats.CurrentLife != 90 {
		t.Error("should have taken 10 double damage from soloattack")
	}
}
func Test_SoloAttackingBehavior_WithNoNeighbors_TopSlot_DealsDoubleDamage(t *testing.T) {
	source, target := setup()

	source.Stats.Type = fish.SoloBasedWeapon
	sourceIndex := 0
	fishSlice := make([]*fish.Fish, 4)
	fishSlice = slices.Insert(fishSlice, sourceIndex, source)
	fish.SoloAttackingBehavior(source, target, sourceIndex, fishSlice)

	if target.Stats.CurrentLife != 90 {
		t.Error("should have taken 10 double damage from soloattack")
	}
}
func Test_SoloAttackingBehavior_WithNoNeighbors_BottomSlot_DealsDoubleDamage(t *testing.T) {
	source, target := setup()

	source.Stats.Type = fish.SoloBasedWeapon
	sourceIndex := 4
	fishSlice := make([]*fish.Fish, 4)
	fishSlice = slices.Insert(fishSlice, sourceIndex, source)
	fish.SoloAttackingBehavior(source, target, sourceIndex, fishSlice)

	if target.Stats.CurrentLife != 90 {
		t.Error("should have taken 10 double damage from soloattack")
	}
}
func Test_SoloAttackingBehavior_WithOneNeighbor_DealsNormalDamage(t *testing.T) {
	source, target := setup()
	source.Stats.Type = fish.SoloBasedWeapon

	otherFishStats := fish.NewStats(fish.Weapon, fish.SizeSmall, 100, 1, 5)
	otherFish := fish.NewFish(ENV, "blocking-solo", "", &otherFishStats)

	sourceIndex := 2
	fishSlice := make([]*fish.Fish, 4)
	fishSlice = slices.Insert(fishSlice, sourceIndex, source)
	fishSlice = slices.Insert(fishSlice, sourceIndex+1, otherFish)
	fish.SoloAttackingBehavior(source, target, sourceIndex, fishSlice)

	if target.Stats.CurrentLife != 95 {
		t.Error("should have taken undoubled 5 damage from soloattack")
	}
}
