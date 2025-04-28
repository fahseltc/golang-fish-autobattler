package environment

import "log/slog"

type Env struct {
	slog.Logger
	*Registry
}

func NewEnv(logger *slog.Logger, registry *Registry) *Env {
	return &Env{
		Logger:   *logger,
		Registry: registry,
	}
}

func Null() *Env {
	return NewEnv(nil, nil)
}
