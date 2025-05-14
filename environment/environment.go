package environment

import (
	"log/slog"
)

type Env struct {
	slog.Logger
	*Config
}

func NewEnv(logger *slog.Logger, config *Config) *Env {
	return &Env{
		Logger: *logger,
		Config: config,
	}
}

func Null() *Env {
	return NewEnv(nil, nil)
}
