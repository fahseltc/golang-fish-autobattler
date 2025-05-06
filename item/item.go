package item

import (
	"fishgame/environment"
	"fishgame/util"
	"fmt"
	"image"
	"image/color"
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
	hitbox      *image.Alpha

	Duration    float64 `json:"duration"`
	CurrentTime float64 `json:"current_time"`
	Damage      int     `json:"damage"`

	Activate func(*Item, *Item) bool `json:"-"`
	// React        func(*Item, *Item) bool `json:"-"`
	HitLastFrame bool `json:"-"`

	X, Y      int
	Width     int
	Height    int
	Dragging  bool
	SlotIndex int

	OffsetX int
	OffsetY int
}

func NewItem(env environment.Env, name string, iType Type, life int, duration float64, damage int, activate func(*Item, *Item) bool) *Item {
	it := new(Item)
	Env = &env
	it.Id = uuid.New()
	it.Name = name
	it.Alive = true

	it.Sprite = util.LoadImage(env, fmt.Sprintf("assets/fish/%s.png", strings.ToLower(it.Name)))

	it.Life = life
	it.CurrentLife = life

	it.Type = iType

	it.Duration = duration
	it.CurrentTime = 0
	it.Damage = damage

	it.Activate = activate
	it.HitLastFrame = false
	it.OffsetX = 32
	it.OffsetY = 32

	return it
}

func (it *Item) RegenerateUuid() {
	it.Id = uuid.New()
}

func (it *Item) Update(dt float64, enemyItems *Collection) bool {
	if it.hitbox == nil {
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

	// Update
	it.CurrentTime += dt
	if it.CurrentTime >= it.Duration {
		it.CurrentTime -= it.Duration
		//index, target := enemyItems.GetRandomActive()

		// if target != nil {
		// 	// trigger weapon item
		// 	if it.Activate != nil && it.Type.String() == "weapon" {
		// 		if !it.Activate(it, target) {
		// 			// remove the item from the enemy's active items and add it to the inactive items
		// 			enemyItems.ActiveItems = append(enemyItems.ActiveItems[:index], enemyItems.ActiveItems[index+1:]...)
		// 			enemyItems.InactiveItems = append(enemyItems.InactiveItems, target)
		// 		}
		// 	}
		// 	// trigger reactive item
		// 	if it.Activate != nil && target.Type.String() == "reactive" && it.HitLastFrame {
		// 		it.Activate(it, target)
		// 		it.HitLastFrame = false
		// 	}
		// }
	}
	it.handleDrag()
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

func (it *Item) handleDrag() {
	// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	// 	mx, my := ebiten.CursorPosition()
	// 	fmt.Printf("Cursor position: Mouse Pressed: %v, %v\n", mx, my)
	// 	if it.Collides(mx, my) {
	// 		fmt.Println("Collides")
	// 		it.Dragging = true
	// 		it.OffsetX = mx - it.Sprite.Bounds().Max.X
	// 		it.OffsetY = my - it.Sprite.Bounds().Max.X
	// 	}
	// }

	// if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
	// 	mx, my := ebiten.CursorPosition()
	// 	if it.Dragging {
	// 		it.Dragging = false
	// 		//if mx >= slot.X && mx <= slot.X+slot.Width && my >= slot.Y && my <= slot.Y+slot.Height {
	// 		// TODO: Find slot to put it in
	// 		it.X = mx
	// 		it.Y = my
	// 	}
	// }

	// if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
	// 	mx, my := ebiten.CursorPosition()
	// 	it.Sprite.Pos = mx - it.OffsetX
	// 	it.Y = my - it.OffsetY
	// }
}

func (it *Item) Collides(x, y int) bool {
	if it.hitbox == nil {
		return false
	}

	collides := it.hitbox.At(x-it.X, y-it.Y).(color.Alpha).A > 0
	// if collides {
	// 	fmt.Printf("collision: %v\n", collides)
	// }
	return collides
}
