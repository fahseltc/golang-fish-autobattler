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
	s.ItemsRegistry = loader.LoadCsv(*s.Env)
	s.Ui = ui.NewUI(s.Env)

	s.Player1 = generatePlayer1(*s.Env, s.ItemsRegistry)
	//s.Player2 = generatePlayer2(*s.Env, s.ItemsRegistry)

	for index, it := range s.Player1.Items.ActiveItems {
		s.Ui.Player1Slots[index].AddItem(index, it)
		it.SlotIndex = index
	}

	s.EncounterManager = encounter.NewManager(s.Env, s.Player1)
}

func (s *Play) Update(dt float64) {
	switch s.State {
	case PlayState:
		updatePlayState(s, dt)
	case MapState:
		updateMapState(s, dt)
	case InventoryState:
		// Show the inventory
	case GameOverState:
		// Show the game over screen
	case PauseState:
		// Pause the game
	}
}

func updatePlayState(s *Play, dt float64) {
	if s.Ui != nil {
		s.Ui.Update()
	}
	if s.EncounterManager.Current.Type == encounter.EncounterTypeBattle {
		s.EncounterManager.Current.Behavior.GetItems()
		s.Player1.Items.Update(dt, s.EncounterManager.Current.Behavior.GetItems())
		s.EncounterManager.Update(dt)
		if len(s.Player1.Items.ActiveItems) == 0 {
			fmt.Println("Player 1 has no active items. GAME OVER")
		}
	}
	//return nil
}

func updateMapState(s *Play, dt float64) error {
	return nil
}

func (s *Play) Draw(screen *ebiten.Image) {
	switch s.State {
	case PlayState:
		// Draw the UI
		if s.Ui != nil {
			s.Ui.Draw(screen)
		}
		if s.Player1 != nil {
			s.Player1.Items.Draw(*s.Env, screen, 1)
		}
		if s.EncounterManager.Current.Behavior.GetItems() != nil {
			s.EncounterManager.Current.Behavior.GetItems().Draw(*s.Env, screen, 2)
		}

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

func generatePlayer1(env environment.Env, items *item.Registry) *player.Player {
	item1, err := items.Get("Goldfish")
	if err {
		panic(err)
	}
	item2, err := items.Get("Eel")
	if err {
		panic(err)
	}
	item3, err := items.Get("Shark")
	if err {
		panic(err)
	}
	item4, err := items.Get("Puffer")
	if err {
		panic(err)
	}
	itemsArr := []*item.Item{&item1, &item2, &item3, &item4}
	coll := item.NewPlayerCollection(&env, itemsArr)
	p := &player.Player{
		Env:   env,
		Name:  "Player 1",
		Items: coll,
	}
	fmt.Printf("Player 1 active items length: %d\n", len(p.Items.ActiveItems))

	return p
}
