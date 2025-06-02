package environment

import (
	"log/slog"
	"os"
)

type Env struct {
	slog.Logger
	*Config
	*EventBus
}

func NewEnv(logger *slog.Logger, config *Config) *Env {
	eventBus := NewEventBus()
	if logger == nil {
		handler := slog.NewJSONHandler(os.Stdout, nil)
		logger = slog.New(handler)
	}
	return &Env{
		Logger:   *logger,
		Config:   config,
		EventBus: eventBus,
	}
}
