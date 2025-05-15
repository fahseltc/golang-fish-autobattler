package encounter

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/player"
	"fishgame/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Initial struct {
	env     *environment.Env
	manager *Manager
	Type    Type

	player *player.Player

	text    string
	bg      *ebiten.Image
	font    *text.Face
	buttons []*ui.Button

	itemChosen bool
}

func (i Initial) Update(dt float64, p *player.Player) {
	for _, button := range i.buttons {
		button.Update()
	}
	if i.itemChosen {
		i.manager.NextEncounter()
	}
}

func (i Initial) Draw(screen *ebiten.Image) {
	screen.DrawImage(i.bg, nil)
	ui.DrawCenteredText(screen, *i.font, i.text, 400, 100)
	// get 3 starter fishes, put buttons of them on screen, wait for user to click a button
	for _, button := range i.buttons {
		button.Draw(screen)
	}
}

func (i Initial) GetItems() *item.Collection {
	return nil
}
func (i Initial) IsDone() bool {
	return i.itemChosen
}

func (i Initial) GetType() Type {
	return i.Type
}
