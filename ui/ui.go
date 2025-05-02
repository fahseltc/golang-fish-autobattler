package ui

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/util"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Slot struct {
	X, Y   int
	Item   *item.Item
	Width  int
	Height int
}

// type Item struct {
// 	X, Y     int
// 	Width    int
// 	Height   int
// 	Dragging bool
// 	OffsetX  int
// 	OffsetY  int
// }

type UI struct {
	Env     *environment.Env
	Slots   []*Slot
	Items   []*item.Item
	SlotImg *ebiten.Image
}

func NewUI(env *environment.Env) *UI {
	return &UI{
		Env:     env,
		Slots:   []*Slot{},
		Items:   []*item.Item{},
		SlotImg: util.LoadImage(*env, "assets/slot.png"),
	}
}

func (ui *UI) AddSlot(x int, y int, item *item.Item) {
	slot := &Slot{X: x, Y: y, Width: 64, Height: 64}
	slot.Item = item
	ui.Slots = append(ui.Slots, slot)
	ui.Items = append(ui.Items, item)
}

// func (ui *UI) AddItem(x int, y int) {
// 	ui.Items = append(ui.Items, &Item{X: x, Y: y, Width: 64, Height: 64})
// }

func (ui *UI) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		fmt.Printf("Cursor position: Mouse Pressed: %v, %v\n", mx, my)
		for _, item := range ui.Items {
			if mx >= item.X && mx <= item.X+item.Width && my >= item.Y && my <= item.Y+item.Height {
				item.Dragging = true
				item.OffsetX = mx - item.X
				item.OffsetY = my - item.Y
				break
			}
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		for _, item := range ui.Items {
			if item.Dragging {
				item.Dragging = false
				for _, slot := range ui.Slots {
					if mx >= slot.X && mx <= slot.X+slot.Width && my >= slot.Y && my <= slot.Y+slot.Height {
						item.X = slot.X
						item.Y = slot.Y
						slot.Item = item
						break
					}
				}
			}
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		for _, item := range ui.Items {
			if item.Dragging {
				item.X = mx - item.OffsetX
				item.Y = my - item.OffsetY
			}
		}
	}
}

func (ui *UI) Draw(screen *ebiten.Image) {
	for _, slot := range ui.Slots {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(slot.X), float64(slot.Y))
		screen.DrawImage(ui.SlotImg, op)
	}
}
