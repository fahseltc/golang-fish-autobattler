package fish

import (
	"fishgame/shared/environment"
	"testing"
)

// [UnitOfWork_StateUnderTest_ExpectedBehaviour]
// https://osherove.com/blog/2005/4/3/naming-standards-for-unit-tests.html

var ENV *environment.Env

func setupEnv() {
	ENV = environment.NewEnv(nil, nil)
}

func Test_TakeDamage_WithUndamagedFish_TakesDamage(t *testing.T) {
	setupEnv()
	fish := NewFish(ENV, "Goldfish", "he little", NewWeaponStats(20, 1, 5))

	fish.TakeDamage(10)
	if fish.Stats.CurrentLife != 10 {
		t.Error("Fish life should be 10")
	}
	if fish.IsDead() {
		t.Error("Fish should be alive")
	}
}

func Test_TakeDamage_AllItsLife_IsNotAlive(t *testing.T) {
	setupEnv()
	fish := NewFish(ENV, "Shrimp", "he shramp", NewWeaponStats(100, 1, 5))

	fish.TakeDamage(100)

	if fish.IsAlive() {
		t.Error("Fish should be dead")
	}
	if !fish.IsDead() {
		t.Error("Fish should be dead")
	}
}

func TestEvent_TakeDamage_AllItsLife_SendsFishDiedEvent(t *testing.T) {
	setupEnv()
	fish := NewFish(ENV, "Shrimp", "he shramp", NewWeaponStats(100, 1, 5))
	eventReceived := false
	ENV.EventBus.Subscribe("FishDiedEvent", func(event environment.Event) {
		eventReceived = true
	})
	fish.TakeDamage(100)

	if !eventReceived {
		t.Error("Fish dying should send a FishDiedEvent")
	}
}

func TestEvent_Activate_WithTarget_SendsFishAttackedEvent(t *testing.T) {
	setupEnv()
	fish := NewFish(ENV, "Shrimp", "he shramp", NewWeaponStats(100, 1, 5))
	target := NewFish(ENV, "Kelp", "he get hit", NewWeaponStats(5, 999, 0))

	var event *environment.Event
	ENV.EventBus.Subscribe("FishAttackedEvent", func(eventIncoming environment.Event) {
		event = &eventIncoming
	})
	fish.Stats.ActivateFunc(fish, target)
	if event == nil {
		t.Error("Fish attacking should send a FishAttackedEvent")
	}
	if event.Data.(environment.FishAttackedEvent).Damage != 5 {
		t.Error("Fish attacking should have 5 damage")
	}
	if event.Data.(environment.FishAttackedEvent).SourceId != fish.Id {
		t.Error("Fish attacking event should have same ID as fish that attacked")
	}
	if event.Data.(environment.FishAttackedEvent).TargetId != target.Id {
		t.Error("Fish attacking event should have same ID as fish that attacked")
	}
	if event.Data.(environment.FishAttackedEvent).Type != Weapon.String() {
		t.Error("Fish attacking event should have same Type as fish that attacked")
	}

}
