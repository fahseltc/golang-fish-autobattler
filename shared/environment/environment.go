package environment

import (
	"log/slog"
	"os"

	"github.com/google/uuid"
)

type Env struct {
	uuid.UUID
	slog.Logger
	*Config
	*EventBus
	*Fonts
}

func NewEnv(logger *slog.Logger, config *Config) *Env {
	eventBus := NewEventBus()
	if logger == nil {
		handler := slog.NewJSONHandler(os.Stdout, nil)
		logger = slog.New(handler)
	}
	return &Env{
		UUID:     uuid.New(),
		Logger:   *logger,
		Config:   config,
		EventBus: eventBus,
		Fonts:    NewFontsCollection(),
	}
}
