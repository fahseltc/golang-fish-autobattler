package fish

import (
	"fishgame/shared/environment"
	"testing"
)

func setup() (source *Fish, target *Fish) {
	ENV = environment.NewEnv(nil, nil)
	source = NewFish(ENV, "sourcefish", "", NewStats(Weapon, SizeSmall, 100, 1, 5))
	target = NewFish(ENV, "targetfish", "", NewStats(Weapon, SizeSmall, 100, 1, 5))
	return source, target
}

func Test_sendFishAttackedEvent_WithValidData_PublishesEvent(t *testing.T) {
	source, target := setup()

	var event *environment.Event
	ENV.EventBus.Subscribe("FishAttackedEvent", func(eventIncoming environment.Event) {
		event = &eventIncoming
	})
	sendFishAttackedEvent(source, target, 10)
	fae := event.Data.(environment.FishAttackedEvent)
	if event.Type != "FishAttackedEvent" {
		t.Error("wrong event")
	}
	if fae.Damage != 10 {
		t.Error("damage should be 10")
	}
	if fae.SourceId != source.Id {
		t.Error("source uuid's should match")
	}
	if fae.TargetId != target.Id {
		t.Error("target uuid's should match")
	}
	if fae.Type != "weapon" {
		t.Error("attack type should be 'weapon'")
	}
}

func Test_AttackingBehavior_WithAliveTarget_DealsDamage(t *testing.T) {
	source, target := setup()

	targetIsAlive := AttackingBehavior(source, target, 0, nil)
	if !targetIsAlive {
		t.Error("target should be alive after one attack")
	}
	if target.Stats.CurrentLife != 95 {
		t.Error("target should have taken 5 damage")
	}
}

func Test_AttackingBehavior_WithAliveTarget_PublishesEvent(t *testing.T) {
	source, target := setup()
	var event *environment.Event
	ENV.EventBus.Subscribe("FishAttackedEvent", func(eventIncoming environment.Event) {
		event = &eventIncoming
	})
	AttackingBehavior(source, target, 0, nil)
	if event.Type != "FishAttackedEvent" {
		t.Error("wrong event")
	}

	fae := event.Data.(environment.FishAttackedEvent)
	if fae.Damage != 5 {
		t.Error("target should have taken 5 damage from event data")
	}
}

func Test_AttackingBehavior_WithDeadTarget_DoesntDealDamage(t *testing.T) {
	source, target := setup()
	target.Stats.CurrentLife = 0

	targetIsAlive := AttackingBehavior(source, target, 0, nil)
	if targetIsAlive {
		t.Error("target fish should be already dead")
	}
	if target.Stats.CurrentLife != 0 {
		t.Error("target fish life should be zero")
	}
}

// func Test_VenomousBehavior_WithAliveTarget_CreatesDebuff(t *testing.T) {
// 	// has double duration of the attacking fish
// }
// func Test_VenomousBehavior_WithAliveTarget_DealsDamageAfterTick(t *testing.T) {}
// func Test_VenomousBehavior_WithAliveTarget_DebuffStacks(t *testing.T)         {}
