package fish

import "testing"

// [UnitOfWork_StateUnderTest_ExpectedBehaviour]
// https://osherove.com/blog/2005/4/3/naming-standards-for-unit-tests.html

func Test_TakeDamage_WithUndamagedFish_TakesDamage(t *testing.T) {
	fish := NewFish("Goldfish", "he little", NewWeaponStats(20, 1, 5))

	fish.TakeDamage(10)
	if fish.Stats.CurrentLife != 10 {
		t.Error("Fish life should be 10")
	}
	if fish.IsDead() {
		t.Error("Fish should be alive")
	}
}

func Test_TakeDamage_AllItsLife_IsNotAlive(t *testing.T) {
	fish := NewFish("Shrimp", "he shramp", NewWeaponStats(100, 1, 5))

	fish.TakeDamage(100)

	if fish.IsAlive() {
		t.Error("Fish should be dead")
	}
	if !fish.IsDead() {
		t.Error("Fish should be dead")
	}
}
