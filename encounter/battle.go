package encounter

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/loader"
	"fishgame/player"

	"github.com/hajimehoshi/ebiten/v2"
)

type Battle struct {
	env   *environment.Env
	Name  string
	items *item.Collection
}

func NewBattleEncounter(env *environment.Env, name string) *Battle {
	itemsReg := loader.LoadCsv(*env)
	battle := &Battle{
		env:   env,
		Name:  name,
		items: generateBattleItems(*env, itemsReg),
	}

	// todo add to slots ui?
	// for index, it := range s.Player2.Items.ActiveItems {
	// 	s.Ui.Player2Slots[index].AddItem(index, it)
	// 	it.SlotIndex = index
	// }

	return battle
}

func (battle *Battle) Update(dt float64, player *player.Player) {
	battle.items.Update(dt, player.Items)
}

func (battle *Battle) Draw(screen *ebiten.Image) {

}

func (battle *Battle) GetItems() *item.Collection {
	return battle.items
}
func (battle *Battle) IsDone() bool {
	return len(battle.items.ActiveItems) == 0
}

// TODO: this should be read from a JSON of different battles with assigned difficulties
func generateBattleItems(env environment.Env, items *item.Registry) *item.Collection {
	item1, err := items.Get("Shark")
	if err {
		panic(err)
	}
	item2, err := items.Get("Minnow")
	if err {
		panic(err)
	}
	item3, err := items.Get("Minnow")
	if err {
		panic(err)
	}
	item4, err := items.Get("Shark")
	if err {
		panic(err)
	}
	item5, err := items.Get("Whale")
	if err {
		panic(err)
	}
	itemsArr := []*item.Item{&item1, &item2, &item3, &item4, &item5}
	coll := item.NewEncounterCollection(&env, itemsArr)

	return coll
}
