package scene

type GameState int

const (
	PlayState GameState = iota
	MapState
	PauseState
	InventoryState
	GameOverState
)
