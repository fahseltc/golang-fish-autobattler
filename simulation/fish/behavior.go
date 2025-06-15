package fish

import (
	"fishgame/shared/environment"
	"fmt"
	"time"
)

func sendFishAttackedEvent(source *Fish, target *Fish, dmg int) {
	source.env.EventBus.Publish(environment.Event{
		Type:      "FishAttackedEvent",
		Timestamp: time.Now(),
		Data: environment.FishAttackedEvent{
			SourceId: source.Id,
			TargetId: target.Id,
			Type:     source.Stats.Type.String(),
			Damage:   dmg,
		},
	})
}

func AttackingBehavior(source *Fish, target *Fish, index int, sourceCollection []*Fish) bool {
	if target.IsAlive() {
		sendFishAttackedEvent(source, target, source.Stats.Damage)
		target.TakeDamage(source.Stats.Damage)
	}

	return target.IsAlive()
}

func VenomousBehavior(source *Fish, target *Fish, index int, sourceCollection []*Fish) bool {
	if target.IsAlive() {
		dbf := NewItemDebuff(target, DebuffTypeVenom, source.Stats.MaxDuration*2, 1, source.Stats.Damage)
		sendFishAttackedEvent(source, target, source.Stats.Damage)
		target.AddDebuff(dbf)
	}
	return target.IsAlive()
}

func LargerSizeAttackingBehavior(source *Fish, target *Fish, index int, sourceCollection []*Fish) bool {
	if target.IsAlive() {
		fmt.Printf("LargerSizeAttackingBehavior: %v, source: %v, target: %v\n", source.Name, source.Stats.Size, target.Stats.Size)
		if source.Stats.Size > target.Stats.Size {
			sendFishAttackedEvent(source, target, source.Stats.Damage*2)
			target.TakeDamage(source.Stats.Damage * 2) // double damage to smaller fish
		} else {
			sendFishAttackedEvent(source, target, source.Stats.Damage)
			target.TakeDamage(source.Stats.Damage)
		}
	}
	return target.IsAlive()
}

func SoloAttackingBehavior(source *Fish, target *Fish, index int, sourceCollection []*Fish) bool {
	if target.IsAlive() {
		aboveFish := true
		belowFish := true
		if index == 0 {
			aboveFish = false
			belowFish = (sourceCollection[index+1] != nil)
		} else if index == 4 {
			aboveFish = (sourceCollection[index-1] != nil)
			belowFish = false
		} else {
			aboveFish = (sourceCollection[index-1] != nil)
			belowFish = (sourceCollection[index+1] != nil)
		}

		if !belowFish && !aboveFish { // fish has no neighbors
			doubleDamage := source.Stats.Damage * 2
			sendFishAttackedEvent(source, target, doubleDamage)
			target.TakeDamage(doubleDamage)
		} else { // fish has at least one neighbor
			sendFishAttackedEvent(source, target, source.Stats.Damage)
			target.TakeDamage(source.Stats.Damage)
		}
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
