package ui

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/shapes"
	"fishgame/simulation"
	"fishgame/util"
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var ENV *environment.Env

var SlotCount int = 5

var spriteScale float64
var spriteSizePx int

type UI struct {
	// Player            *player.Player
	// Player1Slots      map[int]*Slot
	// Player1Collection *item.Collection
	// Player2Slots      map[int]*Slot
	sim         *simulation.SimulationInterface
	slotImg     *ebiten.Image
	Font        text.Face
	smallFont   text.Face
	CurrencyImg *ebiten.Image

	attackLines []*AttackLine

	startSimBtn *Button
	stopSimBtn  *Button
}

func NewUI(env *environment.Env, sim simulation.SimulationInterface) *UI {
	ENV = env
	// set package consts
	spriteScale = ENV.Get("sprite.scale").(float64)
	spriteSizePx = ENV.Get("sprite.sizeInPx").(int)

	font := ENV.Fonts.Med
	smallFont := ENV.Fonts.Small

	ui := &UI{
		sim:         sim,
		slotImg:     util.LoadImage(env, "assets/slot.png"),
		CurrencyImg: util.LoadImage(env, "assets/ui/icons/fishfood.png"),
		Font:        font,
		smallFont:   smallFont,
	}

	// ui.Player1Collection = playerCollection
	// ui.Player1Slots = make(map[int]*Slot, SlotCount)
	// for index := range SlotCount {
	// 	ui.Player1Slots[index] = NewPlayerSlot(index)
	// }

	// ui.Player2Slots = make(map[int]*Slot, SlotCount)
	// for index := range SlotCount {
	// 	ui.Player2Slots[index] = NewEncounterSlot(2, index)
	// }

	env.EventBus.Subscribe("ItemAttackedEvent", ui.ItemAttackedEventHandler)
	ui.attackLines = make([]*AttackLine, 0)

	ui.startSimBtn = NewButton(
		WithRect(shapes.Rectangle{X: 250, Y: 525, W: 150, H: 75}),
		WithText("StartSim"),
		WithClickFunc(func() {
			ENV.EventBus.Publish(environment.Event{Type: "StartSimulation", Timestamp: time.Now()})
		}),
		WithCenteredPos(),
	)
	ui.stopSimBtn = NewButton(
		WithRect(shapes.Rectangle{X: 550, Y: 525, W: 150, H: 75}),
		WithText("StopSim"),
		WithClickFunc(func() {
			ENV.EventBus.Publish(environment.Event{Type: "StopSimulation", Timestamp: time.Now()})
		}),
		WithCenteredPos(),
	)

	return ui
}

func (ui *UI) Update(dt float64) {
	ui.startSimBtn.Update()
	ui.stopSimBtn.Update()
	// var draggingItem *item.Item
	// var previousSlot *Slot

	// for _, slot := range ui.Player1Slots {
	// 	if slot.item != nil && slot.item.Dragging {
	// 		draggingItem = slot.item
	// 		previousSlot = ui.Player1Slots[slot.item.SlotIndex]
	// 	}
	// 	// Clean up the slot if the item already died
	// 	if slot.item != nil && !slot.item.Alive {
	// 		slot.item.SlotIndex = 999
	// 		slot.item = nil
	// 	}
	// }

	// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	// 	mx, my := ebiten.CursorPosition()
	// 	for _, slot := range ui.Player1Slots {
	// 		if slot.item != nil {
	// 			if slot.item.Collides(mx, my) && !slot.item.Dragging {
	// 				slot.item.Dragging = true
	// 				fmt.Printf("Item: %v picked up from slot: %v\n", slot.item.Name, slot.item.SlotIndex)
	// 			}
	// 		}
	// 	}
	// }

	// if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && draggingItem != nil {
	// 	mx, my := ebiten.CursorPosition()
	// 	for index, slot := range ui.Player1Slots {
	// 		collision := slot.Collides(mx, my)
	// 		if collision.Collides {
	// 			if slot.item != nil { // the slot is occupied already
	// 				fmt.Printf("Item: %v dropped onto a occupied slot.\n", draggingItem.Name)
	// 				// we need to move ItemToMove up/down based
	// 				ui.cascadeItem(draggingItem, collision.Type, slot)

	// 			} else {
	// 				// the slot is empty, put the item in it
	// 				fmt.Printf("Item: %v dropped onto an empty slot: %v\n", draggingItem.Name, slot.index)
	// 				fmt.Printf("previousSlot: %v\n", previousSlot)
	// 				previousSlot.item = nil
	// 				fmt.Printf("previousSlotAfterNil: %v\n", previousSlot)
	// 				slot.item = draggingItem
	// 				slot.item.X = int(slot.rect.X)
	// 				slot.item.Y = int(slot.rect.Y)
	// 				slot.item.SlotIndex = index
	// 				slot.item.Dragging = false
	// 			}
	// 		} else { // there was no collision with a slot
	// 			// put the dragging item back into its initial slot
	// 			prevSlot := ui.Player1Slots[draggingItem.SlotIndex]
	// 			prevSlot.item = draggingItem
	// 			draggingItem.SlotIndex = prevSlot.index
	// 			draggingItem.X = int(prevSlot.rect.X)
	// 			draggingItem.Y = int(prevSlot.rect.Y)
	// 			draggingItem.Dragging = false
	// 		}
	// 	}
	// 	// if ui.Player1Collection.Inventory.Collides(mx, my) {
	// 	// 	fmt.Printf("inventory collision \n")

	// 	// 	ui.Player1Collection.Inventory.AddItem(draggingItem)
	// 	// 	prevSlot := ui.Player1Slots[draggingItem.SlotIndex]
	// 	// 	prevSlot.item = nil
	// 	// 	draggingItem.SlotIndex = 999
	// 	// 	draggingItem.X, draggingItem.Y = ui.Player1Collection.Inventory.GetRandomPos()
	// 	// }

	// }

	// if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
	// 	mx, my := ebiten.CursorPosition()
	// 	for _, slot := range ui.Player1Slots {
	// 		if slot.item != nil && slot.item.Dragging {
	// 			slot.item.X = mx - slot.item.OffsetX
	// 			slot.item.Y = my - slot.item.OffsetY
	// 		}
	// 	}
	// }

	// for _, line := range ui.attackLines {
	// 	line.Update(float32(dt))
	// 	if !line.enabled {
	// 		// remove it from the slice
	// 		ui.attackLines = append(ui.attackLines[:0], ui.attackLines[1:]...)
	// 	}
	// }
}

func (ui *UI) Draw(screen *ebiten.Image) {
	ui.startSimBtn.Draw(screen)
	ui.stopSimBtn.Draw(screen)
	// mx, my := ebiten.CursorPosition()

	// // Draw player slots / progress bars
	// for _, slot := range ui.Player1Slots {
	// 	op := &ebiten.DrawImageOptions{}
	// 	op.GeoM.Scale(spriteScale, spriteScale)
	// 	op.GeoM.Translate(float64(slot.rect.X), float64(slot.rect.Y))

	// 	screen.DrawImage(ui.slotImg, op)
	// 	if slot.item != nil {
	// 		DrawLifeBar(screen, float64(slot.item.CurrentLife)/float64(slot.item.Life), float64(slot.item.X), float64(slot.item.Y))
	// 		DrawProgressBar(screen, float64(slot.item.CurrentTime)/float64(slot.item.Duration), float64(slot.item.X), float64(slot.item.Y))
	// 	}
	// 	if slot.item != nil {
	// 		DrawCenteredText(screen, ui.Font, slot.item.Name,
	// 			int(slot.rect.X)+spriteSizePx/2,
	// 			int(slot.rect.Y)+spriteSizePx/2,
	// 			nil)
	// 		DrawCenteredText(screen, ui.Font, fmt.Sprintf("%v", slot.item.SlotIndex),
	// 			int(slot.rect.X)+spriteSizePx/2,
	// 			int(slot.rect.Y)+spriteSizePx, nil)
	// 	}

	// }

	// // Draw tooltips on top of items
	// for _, slot := range ui.Player1Slots {
	// 	if slot.item != nil && slot.item.Alive {
	// 		slot.DrawTooltip(screen, ui, mx, my, 1)
	// 	}
	// }
	// for _, slot := range ui.Player2Slots {
	// 	if slot.item != nil && slot.item.Alive {
	// 		DrawLifeBar(screen, float64(slot.item.CurrentLife)/float64(slot.item.Life), float64(slot.item.X), float64(slot.item.Y))
	// 		DrawProgressBar(screen, float64(slot.item.CurrentTime)/float64(slot.item.Duration), float64(slot.item.X), float64(slot.item.Y))
	// 	}
	// 	op := &ebiten.DrawImageOptions{}
	// 	op.GeoM.Scale(spriteScale, spriteScale)
	// 	op.GeoM.Translate(float64(slot.rect.X), float64(slot.rect.Y))
	// 	screen.DrawImage(ui.slotImg, op)
	// }

	// // Draw enemy slots
	// for _, slot := range ui.Player2Slots {
	// 	slot.DrawTooltip(screen, ui, mx, my, 2)
	// }

	// // draw attack lines
	// for _, line := range ui.attackLines {
	// 	if line.enabled {
	// 		line.Draw(screen)
	// 	}
	// }
}

// func (ui *UI) setItemSlot(it *item.Item, slot *Slot) {
// 	it.X = int(slot.rect.X)
// 	it.Y = int(slot.rect.Y)
// 	it.SlotIndex = slot.index
// 	ui.Player1Collection.ActiveItems[it.SlotIndex] = it
// 	it.Dragging = false
// 	slot.item = it
// 	fmt.Printf("Placed item: %v in slot %d\n", it.Name, slot.index)
// }

// func (ui *UI) cascadeItem(itemToMove *item.Item, direction CollisionType, targetSlot *Slot) {
// 	// switch direction {
// 	// case CollisionTopHalf: // move the item downwards
// 	// 	fmt.Println("cascadeItem: CollisionTopHalf")
// 	// 	//index := itemToMove.SlotIndex

// 	// case CollisionBottomHalf: // move the item upwards
// 	// 	fmt.Println("cascadeItem: CollisionBottomHalf")
// 	ui.cascadeDownRecursively(0, itemToMove, targetSlot)
// 	//}
// }
// func (ui *UI) cascadeDownRecursively(traversedCount int, currentItem *item.Item, targetSlot *Slot) {
// 	if traversedCount >= SlotCount {
// 		fmt.Println("No more slots available to cascade items.")
// 		return
// 	}

// 	if currentItem != nil && currentItem.SlotIndex == targetSlot.index { // handle case where the item is dropped onto its own slot.
// 		fmt.Println("item was dropped onto itself.")
// 		currentItem.Dragging = false
// 		return
// 	}

// 	//currentIndex := currentItem.SlotIndex
// 	fmt.Printf("cascadeDownRecursively: Item: %v, currently in slot: %v - moving DOWN to slot: %v\n", currentItem.Name, currentItem.SlotIndex, targetSlot.index)
// 	// if currentIndex >= SlotCount {
// 	// 	currentIndex = currentIndex % SlotCount
// 	// }

// 	if targetSlot.item == nil {
// 		// Found an empty slot, place the item here
// 		fmt.Println("target slot item is nil, placing here")
// 		//oldSlot := ui.Player1Slots[currentItem.SlotIndex]
// 		ui.setItemSlot(currentItem, targetSlot)
// 		//oldSlot.item = nil
// 	} else if targetSlot.item != nil && targetSlot.item.Dragging {
// 		// slot is occupied, but its the target dragging item
// 		fmt.Println("Slot is occupied by the dragging item")
// 		ui.setItemSlot(currentItem, targetSlot)
// 	} else {
// 		// Slot is occupied, move the current item to the next slot recursively
// 		// there is still a bug when dragging an item upwards
// 		newItemToMove := targetSlot.item
// 		if newItemToMove == nil {
// 			fmt.Println("newItemToMove is NIL")
// 			return
// 		}
// 		newIndex := newItemToMove.SlotIndex + 1
// 		if newIndex >= SlotCount {
// 			newIndex = newIndex % SlotCount
// 		}
// 		newTargetSlot := ui.Player1Slots[newIndex]
// 		if newTargetSlot == nil {
// 			fmt.Printf("TARGET SLOT IS NIL  -  SHOULD NEVER HAPPEN - index was: %v\n", newIndex)
// 		}

// 		ui.cascadeDownRecursively(traversedCount+1, newItemToMove, newTargetSlot)
// 		oldSlot := ui.Player1Slots[currentItem.SlotIndex]
// 		ui.setItemSlot(currentItem, targetSlot)
// 		oldSlot.item = nil
// 	}
// }

// func (ui *UI) ClearSlots() {
// 	ui.Player2Slots = make(map[int]*Slot, SlotCount)
// 	for index := range SlotCount {
// 		ui.Player2Slots[index] = NewEncounterSlot(2, index)
// 	}
// }

func (ui *UI) DrawPlayerCurrency(screen *ebiten.Image, currency int) {
	// Draw fish food currency UI
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(32, 32)
	screen.DrawImage(ui.CurrencyImg, op)
	DrawCenteredText(screen, ui.Font, fmt.Sprintf("%v", currency), 120, 64, nil)
}

func (ui *UI) ItemAttackedEventHandler(event environment.Event) {
	ItemAttackedEvent, ok := event.Data.(item.ItemAttackedEvent)
	if !ok {
		fmt.Println("Invalid ItemAttackedEvent data")
		return
	}

	// Handle the event
	fmt.Printf("ItemAttacked Recieved in UI layer: %v\n", ItemAttackedEvent.Source)

	ui.attackLines = append(ui.attackLines, NewAttackLine(
		ItemAttackedEvent.Source.X,
		ItemAttackedEvent.Source.Y,
		ItemAttackedEvent.Target.X,
		ItemAttackedEvent.Target.Y,
		float32(ItemAttackedEvent.Source.Duration)-0.5,
	))

	//point1 := ItemAttackedEvent.Source.SlotIndex
	//point2 := ItemAttackedEvent.Target.SlotIndex
	// we dont have access to the Screen here, so we need to create a line image in an array and draw it later in the Draw method)
}
