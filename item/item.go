package item

import (
	"fishgame/environment"

	"github.com/google/uuid"
)

var Env *environment.Env

// add json tags for the struct
type Item struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Type        Type      `json:"type"`
	Life        int       `json:"max_life"`
	CurrentLife int       `json:"current_life"`
	Alive       bool      `json:"alive"`

	Duration    float32 `json:"duration"`
	CurrentTime float32 `json:"current_time"`
	Damage      int     `json:"damage"`

	Activate func(*Item, *Item) bool `json:"-"`
	React    func(*Item, *Item) bool `json:"-"`
	Reacted  bool                    `json:"reacted"`
}

func NewItem(env environment.Env, name string, life int, iType Type, duration float32, damage int, activate func(*Item, *Item) bool, react func(*Item, *Item) bool) *Item {
	it := new(Item)
	Env = &env
	it.Id = uuid.New()
	it.Name = name
	it.Alive = true

	it.Life = life
	it.CurrentLife = life

	it.Type = iType

	it.Duration = duration
	it.CurrentTime = 0
	it.Damage = damage

	it.Activate = activate
	it.React = react
	it.Reacted = false

	return it
}

func (it *Item) Update(dt float32, enemyItems *Collection) bool {
	// Check if the item is alive
	if !it.Alive || it.CurrentLife <= 0 {
		it.CurrentLife = 0
		it.Alive = false
		return false
	}

	// Update
	it.CurrentTime += dt
	if it.CurrentTime >= it.Duration {
		it.Reacted = false
		it.CurrentTime -= it.Duration
		index, target := enemyItems.GetRandomActive()
		if !it.Activate(it, target) { // return value false means the target item just died
			// remove the item from the enemy's active items and add it to the inactive items
			enemyItems.ActiveItems = append(enemyItems.ActiveItems[:index], enemyItems.ActiveItems[index+1:]...)
			enemyItems.InactiveItems = append(enemyItems.InactiveItems, target)
		}
	}
	it.Print()
	return true
}

func (it *Item) TakeDamage(source *Item) bool {
	it.CurrentLife -= source.Damage
	if it.CurrentLife <= 0 {
		it.CurrentLife = 0
		it.Alive = false
	}
	// if !it.Reacted {
	// 	// logic isnt great, needs some work!
	// 	it.React(it, source)
	// 	it.Reacted = true
	// }

	return it.Alive
}
