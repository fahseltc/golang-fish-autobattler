package item

import "log/slog"

func (it *Item) Print() {
	Env.Logger.Info("ItemInfo",
		slog.Group("source", it.ToSlogGroup()...),
	)
}

func (it *Item) ToSlogGroup() []any {

	data := map[string]any{
		"id":   it.Id.String(),
		"name": it.Name,
		"life": map[string]any{
			"current": it.CurrentLife,
			"max":     it.Life,
		},
		"time": map[string]any{
			"current": it.CurrentTime,
			"max":     it.Duration,
		},
		"isAlive": it.Alive,
	}
	group := make([]any, 0, len(data)*2)

	for key, value := range data {
		group = append(group, key, value)
	}

	return group
}
