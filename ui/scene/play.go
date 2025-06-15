package scene

import (
	"fishgame/encounter"
	"fishgame/shared/environment"
	"fishgame/simulation/collection"
	"fishgame/simulation/fish"
	"fishgame/simulation/player"
	"fishgame/simulation/simulation"
	"fishgame/ui/ui"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Play struct {
	SceneManager     *Manager
	EncounterManager *encounter.Manager
	Ui               *ui.UI
	Player           *player.Player

	Simulation simulation.SimulationInterface

	CurrentState PlayState
}

func NewPlayScene(sm *Manager) *Play {
	statsReg := fish.NewFishStatsRegistry(ENV)
	encMgr := encounter.NewEncounterManager(ENV, statsReg)
	_, err := encMgr.GetCurrent()
	if err != nil {
		log.Fatal(err)
	}

	player := player.NewPlayer(ENV, "player1")
	player.Fish = collection.NewCollection(ENV)
	sim := simulation.NewSimulation(ENV, player)

	ui := ui.NewUI(ENV, sim, encMgr)

	// ENV.EventBus.Publish(environment.Event{
	// 	Type: "EnableUiEvent",
	// })

	ENV.EventBus.Publish(environment.Event{
		Type: "NextEncounterEvent",
		Data: environment.NextEncounterEvent{
			EncounterType: "initial",
		},
	})

	//ENV.EventBus.Unsubscribe("EnableUiEvent")

	playScene := &Play{
		SceneManager:     sm,
		EncounterManager: encMgr,
		Ui:               ui,
		Player:           player,
		Simulation:       sim,
		CurrentState:     EncounterState,
	}
	ENV.EventBus.Subscribe("GameOverEvent", playScene.handleGameOverEvent)
	return playScene
}

func (s *Play) Update(dt float64) {
	// enc, _ := s.EncounterManager.GetCurrent()
	// switch enc.GetType() {
	// case encounter.EncounterTypeInitial:

	// case encounter.EncounterTypeBattle:
	// case encounter.EncounterTypeChoice:
	// case encounter.EncounterTypeShop:
	// case encounter.EncounterTypeUnknown:
	// }
	s.Simulation.Update(dt)
	s.Ui.Update(dt)
	//s.EncounterManager.Current.Update(dt, s.Player1)
	//s.Ui.Update(dt)
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
	s.Ui.Draw(screen)
	// // switch s.CurrentState {
	// // case PreparingState:
	// // 	s.Env.Logger.Info("Preparing state, nothing to draw")
	// // case EncounterState:
	// // 	//s.drawEncounterState(screen)
	// // case RewardState:

	// // default:
	// // 	s.Env.Logger.Error("Unknown state in Play scene", "state", s.CurrentState)
	// // }

	// s.Ui.Draw(screen)
	// s.Ui.DrawPlayerCurrency(screen, s.Player1.Currency)

	// // s.EncounterManager.Current.Draw(screen)
	// // if s.CurrentState == RewardState {
	// // 	for _, reward := range s.EncounterManager.Current.GetRewards() {
	// // 		if !reward.Obtained {
	// // 			reward.Draw(screen)
	// // 		}
	// // 	}
	// // }
}

func (s *Play) Destroy() {
	// Clean up resources if necessary

	s.Player = nil
}

func (s *Play) GetName() string {
	return "Play"
}

func (s *Play) handleGameOverEvent(event environment.Event) {
	s.SceneManager.SwitchTo("GameOver", true)
}
