package encounter

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/player"
	"fishgame/reward"

	"github.com/hajimehoshi/ebiten/v2"
)

type Battle struct {
	env   *environment.Env
	Name  string
	items *item.Collection
	Type  Type

	player  *player.Player
	rewards []*reward.Reward
}

func (battle *Battle) Update(dt float64, player *player.Player) {
	battle.items.Update(dt, player.Items)
}

func (battle *Battle) Draw(screen *ebiten.Image) {
	battle.items.Draw(battle.env, screen, 2)
}

func (battle *Battle) GetItems() *item.Collection {
	return battle.items
}

func (battle *Battle) IsDone() bool {
	return len(battle.items.ActiveItems) == 0
}

func (battle *Battle) IsGameOver() bool {
	return len(battle.player.Items.ActiveItems) == 0
}

func (battle Battle) GetType() Type {
	return battle.Type
}

func (battle Battle) GetRewards() []*reward.Reward {
	return battle.rewards
}

func (battle *Battle) AddReward(reward *reward.Reward) {
	battle.rewards = append(battle.rewards, reward)
}
