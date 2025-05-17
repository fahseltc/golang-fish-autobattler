package item

import "log/slog"

// activate func(*Item, *Item) bool

// func ReactiveItem(target *Item) bool {
// 	if it.Type == Reactive {
// 		target.TakeDamage(it)
// 	}
// 	return target.Alive
// }

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
