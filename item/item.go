package item

import (
	"fishgame/environment"
	"fishgame/util"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

var Env *environment.Env

// add json tags for the struct
type Item struct {
	Id          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Type        Type          `json:"type"`
	Life        int           `json:"max_life"`
	CurrentLife int           `json:"current_life"`
	Alive       bool          `json:"alive"`
	Sprite      *ebiten.Image `json:"-"`

	Duration    float64 `json:"duration"`
	CurrentTime float64 `json:"current_time"`
	Damage      int     `json:"damage"`

	Activate func(*Item, *Item) bool `json:"-"`
	// React        func(*Item, *Item) bool `json:"-"`
	HitLastFrame bool `json:"-"`
}

func NewItem(env environment.Env, name string, iType Type, life int, duration float64, damage int, activate func(*Item, *Item) bool) *Item {
	it := new(Item)
	Env = &env
	it.Id = uuid.New()
	it.Name = name
	it.Alive = true

	it.Sprite = util.LoadImage(env, fmt.Sprintf("assets/%s.png", strings.ToLower(it.Name)))

	it.Life = life
	it.CurrentLife = life

	it.Type = iType

	it.Duration = duration
	it.CurrentTime = 0
	it.Damage = damage

	it.Activate = activate
	it.HitLastFrame = false

	return it
}

func (it *Item) RegenerateUuid() {
	it.Id = uuid.New()
}

func (it *Item) Update(dt float64, enemyItems *Collection) bool {
	// Check if the item is alive
	if !it.Alive || it.CurrentLife <= 0 {
		it.CurrentLife = 0
		it.Alive = false
		return false
	}

	// Update
	it.CurrentTime += dt
	if it.CurrentTime >= it.Duration {
		it.CurrentTime -= it.Duration
		index, target := enemyItems.GetRandomActive()

		if target != nil {
			// trigger weapon item
			if it.Activate != nil && it.Type.String() == "weapon" {
				if !it.Activate(it, target) {
					// remove the item from the enemy's active items and add it to the inactive items
					enemyItems.ActiveItems = append(enemyItems.ActiveItems[:index], enemyItems.ActiveItems[index+1:]...)
					enemyItems.InactiveItems = append(enemyItems.InactiveItems, target)
				}
			}
			// trigger reactive item
			if it.Activate != nil && target.Type.String() == "reactive" && it.HitLastFrame {
				it.Activate(it, target)
				it.HitLastFrame = false
			}
		}
	}
	//it.Print()
	return true
}

func (it *Item) TakeDamage(source *Item) bool {
	it.CurrentLife -= source.Damage
	if it.CurrentLife <= 0 {
		it.CurrentLife = 0
		it.Alive = false
	}
	it.HitLastFrame = true

	return it.Alive
}
