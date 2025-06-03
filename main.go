package main

// UI / game init
// Simulation init
// Hook simulation -> UI
// loop until end
import (
	"log/slog"
	"os"
	"time"

	"fishgame-sim/simulation"
)

func main() {
	simulation.NewSimulation(nil, nil, nil)
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
