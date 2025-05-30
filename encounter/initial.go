package encounter

import (
	"fishgame/item"
	"fishgame/player"
	"fishgame/reward"
	"fishgame/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Initial struct {
	manager *Manager
	Type    Type

	player *player.Player

	text    string
	bg      *ebiten.Image
	font    *text.Face
	buttons []*ui.Button
	rewards []*reward.Reward

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
	if !i.itemChosen {
		screen.DrawImage(i.bg, nil)
		ui.DrawCenteredText(screen, *i.font, i.text, 400, 100, nil)
		// get 3 starter fishes, put buttons of them on screen, wait for user to click a button
		for _, button := range i.buttons {
			button.Draw(screen)
		}
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
func (i Initial) IsGameOver() bool {
	return false
}
func (i Initial) IsStarted() bool {
	return true
}
func (i *Initial) GetRewards() []*reward.Reward {
	return i.rewards
}
func (i *Initial) AddReward(reward *reward.Reward) {
	i.rewards = append(i.rewards, reward)
}

func (i *Initial) SetRewards(rewards []*reward.Reward) {
	i.rewards = rewards
}
