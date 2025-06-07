package main

import (
	"fishgame/game"
	"fishgame/shared/environment"
	"log/slog"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var ENV *environment.Env

func main() {
	config := environment.NewConfig()

	ENV = environment.NewEnv(SetupLogger(), config)
	game := game.NewGame(ENV)

	ebiten.SetWindowTitle(config.Get("windowTitle").(string))
	ebiten.SetWindowSize(config.Get("resolution.ext.w").(int), config.Get("resolution.ext.h").(int))
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetTPS(config.Get("targetFPS").(int))

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
