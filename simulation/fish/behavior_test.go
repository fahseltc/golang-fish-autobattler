package fish

import (
	"fishgame/shared/environment"
	"testing"
)

func setup() (source *Fish, target *Fish) {
	ENV = environment.NewEnv(nil, nil)
	stats := NewStats(Weapon, SizeSmall, 100, 1, 5)
	source = NewFish(ENV, "sourcefish", "", &stats)
	stats2 := NewStats(Weapon, SizeSmall, 100, 1, 5)
	target = NewFish(ENV, "targetfish", "", &stats2)
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

func Test_VenomousBehavior_WithAliveTarget_CreatesDebuffWithDoubleDuration(t *testing.T) {
	source, target := setup()
	source.Stats.Type = VenomousBasedWeapon
	source.Stats.MaxDuration = 2
	VenomousBehavior(source, target, 0, nil)
	if len(target.Debuffs) != 1 {
		t.Error("target should have one debuff")
	}
	if target.Debuffs[0].RemainingDuration != 4 {
		t.Error("targets debuff should have 4s duration ")
	}
}

func Test_VenomousBehavior_WithAliveTarget_DealsDamageAfterTicks(t *testing.T) {
	source, target := setup()
	source.Stats.Type = VenomousBasedWeapon
	source.Stats.MaxDuration = 1
	VenomousBehavior(source, target, 0, nil)
	if len(target.Debuffs) != 1 {
		t.Error("target should have one debuff")
	}
	target.UpdateDebuffs(1)
	if target.Stats.CurrentLife != 95 {
		t.Error("target should have taken one tick of 5 damage")
	}
	if len(target.Debuffs) != 1 {
		t.Error("target should have one debuff still")
	}
	target.UpdateDebuffs(1)
	if target.Stats.CurrentLife != 90 {
		t.Error()
	}
	if len(target.Debuffs) != 0 {
		t.Error("target debuff should have expired")
	}
}

func Test_VenomousBehavior_WithAliveTarget_DebuffStacks(t *testing.T) {
	source, target := setup()
	source.Stats.Type = VenomousBasedWeapon
	source.Stats.MaxDuration = 1
	VenomousBehavior(source, target, 0, nil)
	VenomousBehavior(source, target, 0, nil)
	if len(target.Debuffs) != 2 {
		t.Error("target should have two debuffs")
	}
	target.UpdateDebuffs(1)
	if target.Stats.CurrentLife != 90 {
		t.Error("target should have taken 5 damage from each debuff")
	}
	if len(target.Debuffs) != 2 {
		t.Error("target should have one debuff still which hasnt expired yet")
	}
}

// func Test_VenomousBehavior_WithDeadTarget_DoesntDealDamage(t *testing.T) {

func Test_LargerSizeAttackingBehavior_WithEqualSizeTarget_DealsNormalDamage(t *testing.T) {
	source, target := setup()
	//source.Stats.Size = SizeMedium
	source.Stats.Type = SizeBasedWeapon
	LargerSizeAttackingBehavior(source, target, 0, nil)
	if target.Stats.CurrentLife != 95 {
		t.Error("5 damage should have been dealt to same size fish")
	}
}
func Test_LargerSizeAttackingBehavior_WithSmallerSizeTarget_DealsDoubleDamage(t *testing.T) {
	source, target := setup()
	source.Stats.Size = SizeMedium
	source.Stats.Type = SizeBasedWeapon
	LargerSizeAttackingBehavior(source, target, 0, nil)
	if target.Stats.CurrentLife != 90 {
		t.Error("10 damage should have been dealt to smaller size fish")
	}
}
func Test_LargerSizeAttackingBehavior_WithLargerSizeTarget_DealsNormalDamage(t *testing.T) {
	source, target := setup()
	source.Stats.Type = SizeBasedWeapon
	target.Stats.Size = SizeLarge
	LargerSizeAttackingBehavior(source, target, 0, nil)
	if target.Stats.CurrentLife != 95 {
		t.Error("5 damage should have been dealt to larger fish")
	}
}
func Test_LargerSizeAttackingBehavior_WithDeadTarget_DealsNoDamage(t *testing.T) {
	source, target := setup()
	source.Stats.Size = SizeMedium
	source.Stats.Type = SizeBasedWeapon
	target.Stats.CurrentLife = 0
	targetAlive := LargerSizeAttackingBehavior(source, target, 0, nil)
	if targetAlive {
		t.Error("fish should be dead")
	}
	if target.Stats.CurrentLife != 0 {
		t.Error("fish was already dead")
	}
}

// ... LargerSizeAttackingBehavior dead target tests
