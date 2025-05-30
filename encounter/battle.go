package encounter

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/player"
	"fishgame/reward"
	"fishgame/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Battle struct {
	env   *environment.Env
	Name  string
	items *item.Collection
	Type  Type

	player  *player.Player
	rewards []*reward.Reward
	started bool

	startBtn *ui.Button
}

func NewBattleScene(env *environment.Env, encounterData jsonEncounter, player *player.Player) *Battle {
	enc := &Battle{
		env:     env,
		Name:    encounterData.Title,
		Type:    EncounterTypeBattle,
		player:  player,
		started: false,
	}

	enc.startBtn = ui.NewButton(env,
		ui.WithRect(ui.Rectangle{X: 450, Y: 30, W: 200, H: 50}),
		ui.WithText("Start Battle!"),
		ui.WithClickFunc(func() {
			enc.started = true
		}),
		ui.WithCenteredPos(),
	)

	return enc
}

func (battle *Battle) Update(dt float64, player *player.Player) {
	if battle.started {
		battle.items.Update(dt, player.Items)
	} else {
		battle.startBtn.Update()
	}
}

func (battle *Battle) Draw(screen *ebiten.Image) {
	battle.items.Draw(battle.env, screen, 2)
	if !battle.started {
		battle.startBtn.Draw(screen)
	}
}

func (battle *Battle) GetItems() *item.Collection {
	return battle.items
}

func (battle *Battle) IsStarted() bool {
	return battle.started
}

func (battle *Battle) IsDone() bool {
	return len(battle.items.ActiveItems) == 0
}

func (battle *Battle) IsGameOver() bool {
	for _, it := range battle.player.Items.ActiveItems {
		if it != nil && it.Alive {
			return false
		}
	}
	return true
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

func (battle *Battle) SetRewards(rewards []*reward.Reward) {
	battle.rewards = rewards
}
