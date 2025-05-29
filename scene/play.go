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
	ItemsRegistry    *item.Registry
	Ui               *ui.UI
	Player1          *player.Player
	EncounterManager *encounter.Manager

	CurrentState PlayState
}

func (s *Play) Init(sm *Manager) {
	s.SceneManager = sm
	itemReg, _ := loader.GetFishRegistry(s.Env)
	s.ItemsRegistry = itemReg.Reg

	s.Player1 = &player.Player{
		Env:   s.Env,
		Name:  "p1",
		Items: item.NewEmptyPlayerCollection(s.Env),
	}
	s.Player1.Inventory = player.NewInventory(s.Player1)
	s.Ui = ui.NewUI(s.Env, s.Player1.Items)

	s.EncounterManager = encounter.NewManager(s.Env, s.Player1, s.Ui)

	s.CurrentState = EncounterState
}

func (s *Play) Update(dt float64) {
	switch s.CurrentState {
	case PreparingState:
		//s.updatePreparingState(dt)
	case EncounterState:
		updateDuringPlayingState(s, dt)
	case RewardState:
		// s.Env.Logger.Info("GameOver state reached, switching to GameOver scene")
		// s.SceneManager.SwitchTo("GameOver", true)
	default:
		s.Env.Logger.Error("Unknown state in Play scene", "state", s.CurrentState)
	}
}

func updateDuringPlayingState(s *Play, dt float64) {
	if s.Ui != nil {
		s.Ui.Update(dt)
	}
	// switch based on encounter type?
	encItems := s.EncounterManager.Current.GetItems()
	if encItems != nil && s.EncounterManager.Current.IsStarted() {
		s.Player1.Items.Update(dt, encItems)
	}
	s.EncounterManager.Current.Update(dt, s.Player1)

	if s.EncounterManager.Current.IsDone() {
		for _, reward := range s.EncounterManager.Current.GetRewards() {
			res := reward.Obtain(s.Player1)
			if !res {
				s.Env.Logger.Error("unable to add item", "itemName", reward.Item.Name)
			}
		}
		s.Ui.ClearSlots()
		s.EncounterManager.NextEncounter()
	}

	if s.EncounterManager.Current.IsGameOver() {
		s.Env.Logger.Info("GameOver")
		s.SceneManager.SwitchTo("GameOver", true)
	}
}

func (s *Play) Draw(screen *ebiten.Image) {
	switch s.CurrentState {
	case PreparingState:
		s.Env.Logger.Info("Preparing state, nothing to draw")
	case EncounterState:
		//s.drawEncounterState(screen)
	case RewardState:
		s.Env.Logger.Info("Reward state, nothing to draw")
	default:
		s.Env.Logger.Error("Unknown state in Play scene", "state", s.CurrentState)
	}

	s.Player1.Inventory.Draw(screen)
	s.Player1.Items.Draw(s.Env, screen, 1)

	s.Ui.Draw(screen)
	s.Ui.DrawPlayerCurrency(screen, s.Player1.Currency)

	s.EncounterManager.Current.Draw(screen)

}

func (s *Play) Destroy() {
	// Clean up resources if necessary
	s.Env.EventBus.Unsubscribe("ItemAttackedEvent")
	s.Env = nil
	s.Player1 = nil
}

func (s *Play) GetName() string {
	return "Play"
}
