package item

import "log/slog"

// activate func(*Item, *Item) bool

// func ReactiveItem(target *Item) bool {
// 	if it.Type == Reactive {
// 		target.TakeDamage(it)
// 	}
// 	return target.Alive
// }

func AttackingItem(source *Item, target *Item) bool {
	// print into and args
	Env.Logger.Info("ItemAttacked", "source", source.Name, "target", target.Name, "damage", source.Damage)
	//fmt.Printf("SourceItem: '%v' deals '%v' Damage to Target: '%s'\n", source.Name, source.Damage, target.Name)
	if target.Alive {
		targetAlive := target.TakeDamage(source)
		if !targetAlive {
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
