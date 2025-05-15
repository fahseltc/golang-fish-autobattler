package encounter

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/player"

	"github.com/hajimehoshi/ebiten/v2"
)

type Battle struct {
	env   *environment.Env
	Name  string
	items *item.Collection
	Type  Type

	player *player.Player
}

// func NewBattleEncounter(env *environment.Env, player *player.Player, name string) *Battle {
// 	// itemsReg := loader.LoadItemRegistry(env)
// 	// battle := &Battle{
// 	// 	env:    env,
// 	// 	Name:   name,
// 	// 	items:  generateBattleItems(*env, itemsReg),
// 	// 	Type:   EncounterTypeBattle,
// 	// 	player: player,
// 	// }

// 	// todo add to slots ui?
// 	// for index, it := range s.Player2.Items.ActiveItems {
// 	// 	s.Ui.Player2Slots[index].AddItem(index, it)
// 	// 	it.SlotIndex = index
// 	// }

// 	return battle
// }

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
func (battle Battle) GetType() Type {
	return battle.Type
}
