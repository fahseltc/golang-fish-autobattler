package environment

import (
	"log/slog"
)

type Env struct {
	slog.Logger
	*Config
	*EventBus
	*Fonts
}

func NewEnv(logger *slog.Logger, config *Config) *Env {
	eventBus := NewEventBus()
	return &Env{
		Logger:   *logger,
		Config:   config,
		EventBus: eventBus,
		Fonts:    NewFontsCollection(),
	}
}
