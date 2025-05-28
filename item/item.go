package item

import (
	"fishgame/environment"
	"fishgame/util"
	"fmt"
	"image"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

var Env *environment.Env

type Item struct {
	Id          uuid.UUID `json:"id"`
	coll        *Collection
	Name        string        `json:"name"`
	Type        Type          `json:"type"`
	Description string        `json:"description"`
	Life        int           `json:"max_life"`
	CurrentLife int           `json:"current_life"`
	Size        Size          `json:"size"`
	Alive       bool          `json:"alive"`
	Sprite      *ebiten.Image `json:"-"`
	hitbox      *image.Alpha

	Duration    float64 `json:"duration"`
	CurrentTime float64 `json:"current_time"`
	Damage      int     `json:"damage"`

	Activate     func(*Item, *Item, *BehaviorProps) bool `json:"-"`
	HitLastFrame bool                                    `json:"-"`
	debuffs      []DebuffInterface

	X, Y      int
	Dragging  bool
	SlotIndex int

	OffsetX int
	OffsetY int
}

func NewItem(env *environment.Env, coll *Collection, name string, iType Type, sz Size, desc string, life int, duration float64, damage int, activate func(*Item, *Item, *BehaviorProps) bool) *Item {
	it := new(Item)
	Env = env
	it.Id = uuid.New()
	it.coll = coll
	it.Name = name
	it.Alive = true
	it.Size = sz

	spriteScale := env.Get("spriteScale").(float64)
	originalSprite := util.LoadImage(env, fmt.Sprintf("assets/fish/%s.png", strings.ToLower(it.Name)))
	w, h := originalSprite.Size()
	scaled := ebiten.NewImage(int(float64(w)*spriteScale), int(float64(h)*spriteScale))
	op := &ebiten.DrawImageOptions{} // Draw original onto the new image with scaling
	op.GeoM.Scale(spriteScale, spriteScale)
	scaled.DrawImage(originalSprite, op)

	it.Sprite = scaled

	it.Life = life
	it.CurrentLife = life

	it.Type = iType
	it.Description = desc

	it.Duration = duration
	it.CurrentTime = 0
	it.Damage = damage

	it.Activate = activate
	it.HitLastFrame = false
	it.OffsetX = int(float64(32) * spriteScale)
	it.OffsetY = int(float64(32) * spriteScale)

	return it
}

func (it *Item) RegenerateUuid() {
	it.Id = uuid.New()
}

func (it *Item) Update(dt float64, enemyItems *Collection, ic *Collection, index int) bool {
	if it == nil {
		return false
	}
	if it.hitbox == nil { // run once on first loop to generate hitbox once.(cant be done before update loop in ebiten) TODO move this
		bounds := it.Sprite.Bounds()
		ebitenAlphaImage := image.NewAlpha(bounds)
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			for i := bounds.Min.X; i < bounds.Max.X; i++ {
				ebitenAlphaImage.Set(i, j, it.Sprite.At(i, j))
			}
		}

		it.hitbox = ebitenAlphaImage
	}
	// Check if the item is alive
	if !it.Alive || it.CurrentLife <= 0 {
		it.CurrentLife = 0
		it.Alive = false
		return false
	}

	// Update items debuffs
	it.updateDebuffs(dt)
	if !it.Alive {
		return false
	}

	// Update Items counters
	it.CurrentTime += dt
	if it.CurrentTime >= it.Duration {
		it.CurrentTime -= it.Duration
		index, target := enemyItems.GetRandomActive()

		if target != nil && it.Activate != nil {
			Env.EventBus.Publish(environment.Event{
				Type:      "ItemAttackedEvent",
				Timestamp: time.Now(),
				Data: ItemAttackedEvent{
					Source: it,
					Target: target,
					Damage: it.Damage,
				},
			})
			// trigger all weapon item
			if it.Type.String() == "weapon" || it.Type.String() == "sizeBasedWeapon" {
				if !it.Activate(it, target, nil) {
					// remove the item from the enemy's active items and add it to the inactive items
					enemyItems.ActiveItems = append(enemyItems.ActiveItems[:index], enemyItems.ActiveItems[index+1:]...)
					enemyItems.InactiveItems = append(enemyItems.InactiveItems, target)
				}
			}
			// trigger reactive item that was targetted
			if target.Type.String() == "reactive" && it.HitLastFrame {
				it.Activate(it, target, nil)
				it.HitLastFrame = false
			}

			// trigger Venomous and size-based items
			if it.Type.String() == "venomousBasedWeapon" {
				it.Activate(it, target, nil)
			}

			// trigger adjancency-based weapons
			if it.Type.String() == "adjacencyBasedWeapon" {

				props := &BehaviorProps{
					data: make(map[string]any),
				}
				// Bug here
				props.data["itemAbove"] = ic.GetItemAbove(it.SlotIndex)
				props.data["itemBelow"] = ic.GetItemBelow(it.SlotIndex)
				it.Activate(it, target, props)
			}
		}
	}
	return true
}
func (it *Item) updateDebuffs(dt float64) {
	if len(it.debuffs) > 0 {
		for index, dbf := range it.debuffs {
			if dbf.IsDone() {
				// remove the finished debuff from the list
				it.debuffs = append(it.debuffs[:index], it.debuffs[index+1:]...)
			} else {
				dbf.Update(dt)
			}
		}
	}
}

func (it *Item) TakeDamage(amt int, debuff bool) bool {
	it.CurrentLife -= amt
	if it.CurrentLife <= 0 {
		it.CurrentLife = 0
		it.Alive = false
	}
	if !debuff { // only trigger reactive stuff on hits, not debuff ticks
		it.HitLastFrame = true
	}

	return it.Alive
}

func (it *Item) AddDebuff(dbf DebuffInterface) bool {
	if it.Alive {
		it.debuffs = append(it.debuffs, dbf)
		return true
	}
	return false
}

func (it *Item) Collides(x, y int) bool {
	collides := x > int(it.X) && x < int(it.X+it.Sprite.Bounds().Dx()) && y > int(it.Y) && y < int(it.Y+it.Sprite.Bounds().Dy())
	return collides
}

func (it *Item) Dps() float32 {
	if it.Damage > 0 && it.Duration > 0 {
		return float32(it.Damage) / float32(it.Duration)
	}
	return 0
}
