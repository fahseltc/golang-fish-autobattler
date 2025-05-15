package scene

import (
	"fishgame/encounter"
	"fishgame/environment"
	"fishgame/item"
	"fishgame/loader"
	"fishgame/player"
	"fishgame/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Play struct {
	Env              *environment.Env
	SceneManager     *Manager
	State            GameState
	ItemsRegistry    *item.Registry
	Ui               *ui.UI
	Player1          *player.Player
	EncounterManager *encounter.Manager
}

func (s *Play) Init(sm *Manager) {
	s.SceneManager = sm
	s.State = PlayState
	s.ItemsRegistry = loader.LoadItemRegistry(s.Env)
	s.Ui = ui.NewUI(s.Env)

	s.Player1 = &player.Player{
		Env:   s.Env,
		Name:  "p1",
		Items: item.NewEmptyPlayerCollection(s.Env),
	}

	s.EncounterManager = encounter.NewManager(s.Env, s.Player1, s.Ui)
}

func (s *Play) Update(dt float64) {
	switch s.State {
	case PlayState:
		updateDuringPlayState(s, dt)
	case MapState:
		updateDuringMapState(s, dt)
	case InventoryState:
		// Show the inventory
	case GameOverState:
		// Show the game over screen
	case PauseState:
		// Pause the game
	}
}

func updateDuringPlayState(s *Play, dt float64) {
	if s.Ui != nil {
		s.Ui.Update()
	}
	// switch based on encounter type?
	s.Player1.Items.Update(dt, s.EncounterManager.Current.GetItems())
	s.EncounterManager.Current.Update(dt, s.Player1)
	//return nil
}

func updateDuringMapState(s *Play, dt float64) error {
	return nil
}

func (s *Play) Draw(screen *ebiten.Image) {
	switch s.State {
	case PlayState:
		if s.Ui != nil {
			s.Ui.Draw(screen)
		}
		if s.Player1 != nil {
			s.Player1.Items.Draw(s.Env, screen, 1)
		}
		s.EncounterManager.Current.Draw(screen)

	case MapState:
		return
	case InventoryState:
		return
	case GameOverState:
		return
	case PauseState:
		return
	}
}

func (s *Play) Destroy() {
	// Clean up resources if necessary
	s.Env = nil
	s.Player1 = nil
}

func (s *Play) GetName() string {
	return "Play"
}
