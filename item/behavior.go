package item

import (
	"fmt"
	"log/slog"
)

func AttackingBehavior(source *Item, target *Item) bool {
	// print into and args
	//Env.Logger.Info("ItemAttacked", "source", source.Name, "target", target.Name, "damage", source.Damage)
	//fmt.Printf("SourceItem: '%v' deals '%v' Damage to Target: '%s'\n", source.Name, source.Damage, target.Name)
	if target.Alive {
		target.TakeDamage(source.Damage, false)
		if !target.Alive {
			Env.Logger.Info("ItemDied",
				slog.Group(
					"source", source.ToSlogGroup()...,
				),
				slog.Group(
					"target", target.ToSlogGroup()...,
				))
		}
	}

	return target.Alive
}

// In a reactingBehavior, the source fish has already hit the target fish and this method handles the damage done back to source
func ReactingBehavior(source *Item, target *Item) bool {
	// print into and args
	//Env.Logger.Info("ItemReacted", "source", source.Name, "target", target.Name, "damage", source.Damage)
	//fmt.Printf("SourceItem: '%v' deals '%v' Damage to Target: '%s'\n", source.Name, source.Damage, target.Name)
	if target.Alive {
		// the source will take damage from the target
		if source.Type == Reactive {
			source.TakeDamage(target.Damage, false)
			if !source.Alive {
				Env.Logger.Info("ItemDied",
					slog.Group(
						"source", source.ToSlogGroup()...,
					),
					slog.Group(
						"target", target.ToSlogGroup()...,
					))
			}
		}
	}

	return target.Alive
}

func VenomousBehavior(source *Item, target *Item) bool {
	if target.Alive {
		// TODO: Should venom to stack, or for only one instance to be on a target?
		fmt.Printf("Created debf: %v, duration: %v\n", source.Name, source.Duration)
		dbf := NewItemDebuff(target, DebuffTypeVenom, source.Duration, 1, source.Damage)
		target.debuffs = append(target.debuffs, dbf)
		fmt.Printf("Applying venom debuff to: %v\n", target.Name)
	}
	return target.Alive
}

func LargerSizeAttackingBehavior(source *Item, target *Item) bool {
	if target.Alive {
		fmt.Printf("LargerSizeAttackingBehavior, source: %v, target: %v\n", source.Size, target.Size)
		if source.Size > target.Size {
			target.TakeDamage(source.Damage*2, false) // double damage to smaller fish
			fmt.Printf("did double damage\n")
		} else {
			target.TakeDamage(source.Damage, false)
		}

		if !target.Alive {
			Env.Logger.Info("ItemDied",
				slog.Group(
					"source", source.ToSlogGroup()...,
				),
				slog.Group(
					"target", target.ToSlogGroup()...,
				))
		}
	}
	return target.Alive
}
