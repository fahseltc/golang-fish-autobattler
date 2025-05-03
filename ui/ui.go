package ui

import (
	"fishgame/environment"
	"fishgame/util"

	"github.com/hajimehoshi/ebiten/v2"
)

var SlotCount int = 5

type UI struct {
	env          *environment.Env
	Player1Slots map[int]*Slot
	Player2Slots map[int]*Slot
	slotImg      *ebiten.Image
}

func NewUI(env *environment.Env) *UI {
	ui := &UI{
		env:     env,
		slotImg: util.LoadImage(*env, "assets/slot.png"),
	}

	ui.Player1Slots = make(map[int]*Slot, SlotCount)
	for index := range SlotCount {
		ui.Player1Slots[index] = NewSlot(env, 1, index)
	}

	ui.Player2Slots = make(map[int]*Slot, SlotCount)
	for index := range SlotCount {
		ui.Player2Slots[index] = NewSlot(env, 2, index)
	}
	return ui
}

func (ui *UI) Update() {
	// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	// 	mx, my := ebiten.CursorPosition()
	// 	fmt.Printf("Cursor position: Mouse Pressed: %v, %v\n", mx, my)
	// 	// for _, item := range ui.Items {
	// 	// 	if mx >= item.X && mx <= item.X+item.Width && my >= item.Y && my <= item.Y+item.Height {
	// 	// 		item.Dragging = true
	// 	// 		item.OffsetX = mx - item.X
	// 	// 		item.OffsetY = my - item.Y
	// 	// 		break
	// 	// 	}
	// 	// }
	// }

	// if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
	// 	mx, my := ebiten.CursorPosition()
	// 	// for _, item := range ui.Items {
	// 	// 	if item.Dragging {
	// 	// 		item.Dragging = false
	// 	// 		for _, slot := range ui.Slots {
	// 	// 			if mx >= slot.x && mx <= slot.x+slot.width && my >= slot.y && my <= slot.y+slot.height {
	// 	// 				item.X = slot.x
	// 	// 				item.Y = slot.y
	// 	// 				slot.item = item
	// 	// 				break
	// 	// 			}
	// 	// 		}
	// 	// 	}
	// 	// }
	// }

	// if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
	// 	mx, my := ebiten.CursorPosition()
	// 	for _, item := range ui.Items {
	// 		if item.Dragging {
	// 			item.X = mx - item.OffsetX
	// 			item.Y = my - item.OffsetY
	// 		}
	// 	}
	// }
}

func (ui *UI) Draw(screen *ebiten.Image) {
	for _, slot := range ui.Player1Slots {
		if slot.item != nil {
			DrawLifeBar(screen, float64(slot.item.CurrentLife)/float64(slot.item.Life), float64(slot.item.X), float64(slot.item.Y))
			DrawProgressBar(screen, float64(slot.item.CurrentTime)/float64(slot.item.Duration), float64(slot.item.X), float64(slot.item.Y))
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(slot.x), float64(slot.y))
		screen.DrawImage(ui.slotImg, op)
	}
	for _, slot := range ui.Player2Slots {
		if slot.item != nil {
			DrawLifeBar(screen, float64(slot.item.CurrentLife)/float64(slot.item.Life), float64(slot.item.X), float64(slot.item.Y))
			DrawProgressBar(screen, float64(slot.item.CurrentTime)/float64(slot.item.Duration), float64(slot.item.X), float64(slot.item.Y))
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(slot.x), float64(slot.y))
		screen.DrawImage(ui.slotImg, op)
	}
}
