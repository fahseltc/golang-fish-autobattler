package scene

import (
	"fishgame/encounter"
	"fishgame/environment"
	"fishgame/item"
	"fishgame/loader"
	"fishgame/player"
	"fishgame/ui"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Play struct {
	Env          *environment.Env
	SceneManager *Manager
	//State            GameState
	ItemsRegistry    *item.Registry
	Ui               *ui.UI
	Player1          *player.Player
	EncounterManager *encounter.Manager
}

func (s *Play) Init(sm *Manager) {
	s.SceneManager = sm
	itemReg, _ := loader.GetFishRegistry(s.Env)
	s.ItemsRegistry = itemReg.Reg
	s.Ui = ui.NewUI(s.Env)

	s.Player1 = &player.Player{
		Env:   s.Env,
		Name:  "p1",
		Items: item.NewEmptyPlayerCollection(s.Env),
	}

	s.EncounterManager = encounter.NewManager(s.Env, s.Player1, s.Ui)
}

func (s *Play) Update(dt float64) {
	updateDuringPlayState(s, dt)
}

func updateDuringPlayState(s *Play, dt float64) {
	if s.Ui != nil {
		s.Ui.Update()
	}
	// switch based on encounter type?
	encItems := s.EncounterManager.Current.GetItems()
	if encItems != nil {
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
		s.EncounterManager.NextEncounter()
	}
	if s.EncounterManager.Current.IsGameOver() {
		s.Env.Logger.Info("GameOver")
		s.SceneManager.SwitchTo("GameOver", true)
	}
}

func updateDuringGameOverState(s *Play, dt float64) error {
	return nil
}

func (s *Play) Draw(screen *ebiten.Image) {
	if s.Ui != nil {
		s.Ui.Draw(screen)
	}
	if s.Player1 != nil {
		s.Player1.Items.Draw(s.Env, screen, 1)
	}
	s.EncounterManager.Current.Draw(screen)

	// Draw fish food currency UI
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(32, 32)
	screen.DrawImage(s.Ui.CurrencyImg, op)
	ui.DrawCenteredText(screen, s.Ui.Font, fmt.Sprintf("%v", s.Player1.Currency), 120, 64)
}

func (s *Play) Destroy() {
	// Clean up resources if necessary
	s.Env = nil
	s.Player1 = nil
}

func (s *Play) GetName() string {
	return "Play"
}
