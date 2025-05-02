package main

import (
	"fishgame/environment"
	"fishgame/game"
	"log/slog"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	logger := SetupLogger()
	registry := environment.NewRegistry()
	registry.Add("targetFPS", 30)
	registry.Add("width", 800)
	registry.Add("height", 600)

	env := environment.NewEnv(logger, registry)

	game := game.NewGame(env)

	ebiten.SetWindowTitle("Fish Game")
	ebiten.SetWindowSize(registry.Get("width").(int), registry.Get("height").(int))
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetTPS(registry.Get("targetFPS").(int))

	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}

func SetupLogger() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					return slog.Attr{
						Key:   slog.TimeKey,
						Value: slog.Int64Value(t.Unix()),
					}
				}
			}
			return a
		},
	})

	return slog.New(handler)
}
