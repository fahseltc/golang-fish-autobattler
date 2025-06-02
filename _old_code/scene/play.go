package scene

import (
	"fishgame/encounter"
	"fishgame/environment"
	"fishgame/item"
	"fishgame/loader"
	"fishgame/player"
	"fishgame/simulation"
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

	Simulation simulation.SimulationInterface

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

	s.EncounterManager = encounter.NewManager(s.Env, s.Player1, s.Ui)
	s.Simulation = simulation.NewSimulation(s.Env, s.Player1, s.EncounterManager.Current)

	s.Ui = ui.NewUI(s.Env, s.Simulation)

	s.CurrentState = EncounterState
}

func (s *Play) Update(dt float64) {
	s.Simulation.Update(dt)
	//s.EncounterManager.Current.Update(dt, s.Player1)
	s.Ui.Update(dt)
	// switch s.CurrentState {
	// case PreparingState:
	// 	//s.updatePreparingState(dt)
	// case EncounterState:
	// 	updateDuringEncounterState(s, dt)
	// case RewardState:
	// 	updateDuringRewardState(s, dt)
	// default:
	// 	s.Env.Logger.Error("Unknown state in Play scene", "state", s.CurrentState)
	// }
}

// func updateDuringEncounterState(s *Play, dt float64) {
// 	if s.Ui != nil {
// 		s.Ui.Update(dt)
// 	}

// 	encItems := s.EncounterManager.Current.GetItems()
// 	if encItems != nil && s.EncounterManager.Current.IsStarted() {
// 		s.Player1.Items.Update(dt, encItems)
// 	}

// 	s.EncounterManager.Current.Update(dt, s.Player1)

// 	if s.EncounterManager.Current.IsDone() && len(s.EncounterManager.Current.GetRewards()) > 0 {
// 		s.Env.Logger.Info("Encounter done, switching to Reward state")
// 		s.CurrentState = RewardState
// 		return
// 	}

// 	if s.EncounterManager.Current.IsGameOver() {
// 		s.Env.Logger.Info("GameOver")
// 		s.SceneManager.SwitchTo("GameOver", true)
// 	}
// }

// func updateDuringRewardState(s *Play, dt float64) {
// 	for i, reward := range s.EncounterManager.Current.GetRewards() {
// 		if !reward.Obtained {
// 			reward.Update(s.Player1)
// 		} else {
// 			rewards := s.EncounterManager.Current.GetRewards()
// 			// update the list of rewards in the encounter to remove this one
// 			s.EncounterManager.Current.SetRewards(append(rewards[:i], rewards[i+1:]...))
// 		}
// 	}
// 	if len(s.EncounterManager.Current.GetRewards()) == 0 {
// 		s.Env.Logger.Info("All rewards obtained, switching to Encounter state")
// 		s.Ui.ClearSlots()
// 		s.CurrentState = EncounterState // todo change to preparing state
// 		s.EncounterManager.NextEncounter()
// 	}
// }

func (s *Play) Draw(screen *ebiten.Image) {
	// switch s.CurrentState {
	// case PreparingState:
	// 	s.Env.Logger.Info("Preparing state, nothing to draw")
	// case EncounterState:
	// 	//s.drawEncounterState(screen)
	// case RewardState:

	// default:
	// 	s.Env.Logger.Error("Unknown state in Play scene", "state", s.CurrentState)
	// }

	s.Player1.Items.Draw(s.Env, screen, 1)

	s.Ui.Draw(screen)
	s.Ui.DrawPlayerCurrency(screen, s.Player1.Currency)

	s.EncounterManager.Current.Draw(screen)
	if s.CurrentState == RewardState {
		for _, reward := range s.EncounterManager.Current.GetRewards() {
			if !reward.Obtained {
				reward.Draw(screen)
			}
		}
	}
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
