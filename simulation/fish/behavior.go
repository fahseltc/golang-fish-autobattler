package fish

import (
	"fishgame/shared/environment"
	"time"
)

// type BehaviorProps struct {
// 	data map[string]any
// }

func sendFishAttackedEvent(source *Fish, target *Fish) {
	source.env.EventBus.Publish(environment.Event{
		Type:      "FishAttackedEvent",
		Timestamp: time.Now(),
		Data: environment.FishAttackedEvent{
			SourceId: source.Id,
			TargetId: target.Id,
			Type:     source.Stats.Type.String(),
			Damage:   source.Stats.Damage,
		},
	})
}

func AttackingBehavior(source *Fish, target *Fish) bool {
	if target.IsAlive() {
		sendFishAttackedEvent(source, target)
		target.TakeDamage(source.Stats.Damage)
	}

	return target.IsAlive()
}

// // In a reactingBehavior, the source fish has already hit the target fish and this method handles the damage done back to source
// func ReactingBehavior(source *Item, target *Item, props *BehaviorProps) bool {
// 	// print into and args
// 	//Env.Logger.Info("ItemReacted", "source", source.Name, "target", target.Name, "damage", source.Damage)
// 	//fmt.Printf("SourceItem: '%v' deals '%v' Damage to Target: '%s'\n", source.Name, source.Damage, target.Name)
// 	if target.Alive {
// 		// the source will take damage from the target
// 		if source.Type == Reactive {
// 			source.TakeDamage(target.Damage, false)
// 			if !source.Alive {
// 				Env.Logger.Info("ItemDied",
// 					slog.Group(
// 						"source", source.ToSlogGroup()...,
// 					),
// 					slog.Group(
// 						"target", target.ToSlogGroup()...,
// 					))
// 			}
// 		}
// 	}

// 	return target.Alive
// }

func VenomousBehavior(source *Fish, target *Fish) bool {
	if target.IsAlive() {
		// TODO: Should venom to stack, or for only one instance to be on a target?
		dbf := NewItemDebuff(target, DebuffTypeVenom, source.Stats.MaxDuration, 1, source.Stats.Damage)
		sendFishAttackedEvent(source, target)
		target.AddDebuff(dbf)
	}
	return target.IsAlive()
}

func LargerSizeAttackingBehavior(source *Fish, target *Fish) bool {
	if target.IsAlive() {
		//fmt.Printf("LargerSizeAttackingBehavior, source: %v, target: %v\n", source.Size, target.Size)
		if source.Stats.Size > target.Stats.Size {
			sendFishAttackedEvent(source, target)
			target.TakeDamage(source.Stats.Damage * 2) // double damage to smaller fish
		} else {
			sendFishAttackedEvent(source, target)
			target.TakeDamage(source.Stats.Damage)
		}

		// if !target.Alive {
		// 	Env.Logger.Info("ItemDied",
		// 		slog.Group(
		// 			"source", source.ToSlogGroup()...,
		// 		),
		// 		slog.Group(
		// 			"target", target.ToSlogGroup()...,
		// 		))
		// }
	}
	return target.IsAlive()
}

// func AdjacentAttackingBehavior(source *Item, target *Item, props *BehaviorProps) bool {
// 	if target.Alive {
// 		adjacentCount := 0
// 		if props.data["itemAbove"] != nil &&
// 			props.data["itemAbove"].(*Item) != nil &&
// 			props.data["itemAbove"].(*Item).Name == source.Name {
// 			adjacentCount += 1
// 		}
// 		if props.data["itemBelow"] != nil &&
// 			props.data["itemBelow"].(*Item) != nil &&
// 			props.data["itemBelow"].(*Item).Name == source.Name {
// 			adjacentCount += 1
// 		}
// 		fmt.Printf("adjacentFishCount: %v\n", adjacentCount)
// 		target.TakeDamage(source.Damage+adjacentCount, false) // add adjacent count to damage done TODO: make this a variable number

// 		if !target.Alive {
// 			Env.Logger.Info("ItemDied",
// 				slog.Group(
// 					"source", source.ToSlogGroup()...,
// 				),
// 				slog.Group(
// 					"target", target.ToSlogGroup()...,
// 				))
// 		}
// 	}
// 	return target.Alive
// }
