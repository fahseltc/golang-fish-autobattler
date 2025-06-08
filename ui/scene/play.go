package scene

import (
	"fishgame/shared/environment"
	"fishgame/simulation/collection"
	"fishgame/simulation/fish"
	"fishgame/simulation/player"
	"fishgame/simulation/simulation"
	"fishgame/ui/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Play struct {
	SceneManager *Manager
	//ItemsRegistry *item.Registry
	Ui     *ui.UI
	Player *player.Player
	//EncounterManager *encounter.Manager

	Simulation simulation.SimulationInterface

	CurrentState PlayState
}

func NewPlayScene(sm *Manager) *Play {

	//itemReg, _ := loader.GetFishRegistry(s.Env)
	//s.ItemsRegistry = itemReg.Reg

	player := player.NewPlayer(ENV, "player1")
	player.Fish = collection.NewCollection(ENV)
	player.Fish.AddFish(fish.NewFish(ENV, "goldfish", "He's very golden and smoll and doesn't do much!", fish.NewWeaponStats(20, 1, 5)), 0)
	player.Fish.AddFish(fish.NewFish(ENV, "cod", "testfish2", fish.NewWeaponStats(30, 1, 5)), 1)
	//player.Fish.AddFish(fish.NewFish(ENV, "eel", "testfish3", fish.NewWeaponStats(10, 2, 10)), 2)
	player.Fish.AddFish(fish.NewFish(ENV, "whale", "testfish4", fish.NewWeaponStats(10, 2, 10)), 3)
	player.Fish.AddFish(fish.NewFish(ENV, "minnow", "testfish5", fish.NewWeaponStats(10, 2, 10)), 4)

	encounterFish := collection.NewCollection(ENV)

	encounterFish.AddFish(fish.NewFish(ENV, "octopus", "testfish6", fish.NewWeaponStats(35, 3, 20)), 1)
	encounterFish.AddFish(fish.NewFish(ENV, "puffer", "testfish7", fish.NewWeaponStats(35, 3, 20)), 2)
	encounterFish.AddFish(fish.NewFish(ENV, "sunfish", "testfish8", fish.NewWeaponStats(35, 3, 20)), 4)

	sim := simulation.NewSimulation(ENV, player, encounterFish)
	//s.EncounterManager = encounter.NewManager(s.Env, s.Player1, s.Ui)
	ui := ui.NewUI(ENV, sim)

	ENV.EventBus.Publish(environment.Event{
		Type: "EnableUiEvent",
	})

	playScene := &Play{
		SceneManager: sm,
		Ui:           ui,
		Player:       player,
		Simulation:   sim,
		CurrentState: EncounterState,
	}
	return playScene
}

func (s *Play) Update(dt float64) {
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
