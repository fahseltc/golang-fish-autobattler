package scene

type PlayState int

const (
	PausedState PlayState = iota
	PreparingState
	EncounterState
	RewardState
)
