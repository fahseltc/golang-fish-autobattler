package ui

import (
	"fishgame/encounter"
	"fishgame/shared/environment"
	"fishgame/simulation/fish"
	"fishgame/simulation/simulation"
	"fishgame/ui/shapes"
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type DialogInterface interface {
	Draw(screen *ebiten.Image)
	Update()
	IsCompleted() bool
}

var ENV *environment.Env

type UI struct {
	Env *environment.Env

	sim          simulation.SimulationInterface
	encounterMgr *encounter.Manager

	sprites        map[string]*Sprite
	playerSlots    map[int]*Slot
	encounterSlots map[int]*Slot
	inventory      *Sprite
	attackLines    []*AttackLine
	dialogs        []DialogInterface

	startSimBtn *Button
	stopSimBtn  *Button

	enabled bool

	draggingSprite       *Sprite
	draggedFromInventory bool
}

func InitEnv(env *environment.Env) {
	ENV = env
}

func NewUI(env *environment.Env, sim simulation.SimulationInterface, encounterMgr *encounter.Manager) *UI {
	ENV = env

	ui := &UI{
		Env:                  env,
		sim:                  sim,
		encounterMgr:         encounterMgr,
		sprites:              make(map[string]*Sprite),
		playerSlots:          make(map[int]*Slot),
		encounterSlots:       make(map[int]*Slot),
		enabled:              true,
		draggedFromInventory: false,
	}

	ui.startSimBtn = NewButton(
		WithRect(shapes.Rectangle{X: 570, Y: 270, W: 75, H: 50}),
		WithText("Start"),
		WithClickFunc(func() {
			env.EventBus.Publish(environment.Event{Type: "StartSimulationEvent", Timestamp: time.Now()})
		}),
		WithCenteredPos(),
	)
	ui.stopSimBtn = NewButton(
		WithRect(shapes.Rectangle{X: 570, Y: 330, W: 75, H: 50}),
		WithText("Stop"),
		WithClickFunc(func() {
			env.EventBus.Publish(environment.Event{Type: "StopSimulationEvent", Timestamp: time.Now()})
		}),
		WithCenteredPos(),
	)

	for i := 0; i <= 4; i++ {
		ui.playerSlots[i] = NewPlayerSlot(i)
		ui.encounterSlots[i] = NewEncounterSlot(i)
	}

	ui.inventory = NewInventorySprite()

	// StartSimulationEvent - no data associated
	// StopSimulationEvent - no data associated
	// FishAttackedEvent
	// FishDiedEvent
	// GameOverEvent
	// EncounterDoneEvent
	ENV.EventBus.Subscribe("FishAttackedEvent", ui.handleFishAttackedEvent)
	//ENV.EventBus.Subscribe("FishDiedEvent", func(event environment.Event) {})
	//ENV.EventBus.Subscribe("GameOverEvent", func(event environment.Event) {})
	//ENV.EventBus.Subscribe("EncounterDoneEvent", func(event environment.Event) {})
	ENV.EventBus.Subscribe("DisableUiEvent", ui.handleDisableUiEvent)
	ENV.EventBus.Subscribe("EnableUiEvent", ui.handleEnableUiEvent)
	ENV.EventBus.Subscribe("NextEncounterEvent", ui.handleNextEncounterEvent)

	return ui
}

func (ui *UI) Update(dt float64) {
	ui.startSimBtn.Update()
	ui.stopSimBtn.Update()
	if ui.enabled {
		ui.updatePlayerFish()
		ui.updateEncounterFish()
		ui.updateSpritePositionsFromSim()

		for _, line := range ui.attackLines {
			line.Update(dt)
		}

		for i, dialog := range ui.dialogs {
			if dialog.IsCompleted() {
				ui.dialogs = append(ui.dialogs[:i], ui.dialogs[i+1:]...)
				if len(ui.dialogs) == 0 {
					ENV.EventBus.Publish(environment.Event{
						Type: "NextEncounterEvent",
						Data: environment.NextEncounterEvent{
							EncounterType: "battle",
						},
					})
				}
			} else {
				dialog.Update()
			}
		}
	}
}
func (ui *UI) updateSpritePositionsFromSim() {
	if ui.draggingSprite == nil {
		for uuid, sprite := range ui.sprites {
			simFishIndex, simFish := ui.sim.Player_GetFish().ById(uuid)
			if simFish == nil {
				simFishIndex, simFish = ui.sim.Encounter_GetFish().ById(uuid)
			}
			if simFish != nil {
				sprite.SetPosition(simFishIndex)
			}
		}
	}
}

func (ui *UI) deleteOrAddOrFindFish(fish *fish.Fish, index int) *Sprite {
	if fish.IsDead() { // the sim fish is dead, remove its sprite
		delete(ui.sprites, fish.Id.String())
		return nil
	} else if ui.sprites[fish.Id.String()] == nil { // the sim fish is new and needs a sprite made
		var sprite *Sprite
		if index == 999 { // fish in inventory
			sprite = NewInventoryFishSprite(fish)
		} else {
			sprite = NewPlayerFishSprite(fish, index)
		}
		ui.sprites[fish.Id.String()] = sprite
		return sprite
	} else if ui.sprites[fish.Id.String()] != nil { // the sim fish is already added to the list of sprites
		return ui.sprites[fish.Id.String()]
	} else { // the fish is already dead
		return nil
	}
}

func (ui *UI) updatePlayerFish() {
	if ui.sim.IsInitialized() {
		// Handle fish in slots
		for index, fish := range ui.sim.Player_GetFish().GetAllFish() {
			if fish != nil {
				sprite := ui.deleteOrAddOrFindFish(fish, index)
				if sprite == nil {
					continue
				}

				// Handle Clicks
				if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && !ui.sim.IsEnabled() {
					mx, my := ebiten.CursorPosition()
					if sprite.Rect.Collides(float32(mx), float32(my)) {
						fmt.Println("click collides w/ slot fish")
						sprite.Dragging = true
						sprite.SavePositionBeforeDrag()
						ui.draggingSprite = sprite
					}
				}
			}
		}
		// Handle fish in inventory
		for _, fish := range ui.sim.Player_GetInventory().GetAll() {
			if fish != nil {
				sprite := ui.deleteOrAddOrFindFish(fish, 999)
				if sprite == nil {
					continue
				}
				sprite.toolTip.ChangeAlignment(shapes.BottomAlignment) // ensure all inventory fish tooltips are downward
				// create sprite for fish in inventory
				if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && !ui.sim.IsEnabled() {
					mx, my := ebiten.CursorPosition()
					if sprite.Rect.Collides(float32(mx), float32(my)) {
						fmt.Println("click collides w/ inventory fish")
						ui.draggedFromInventory = true
						sprite.Dragging = true
						sprite.SavePositionBeforeDrag()
						ui.draggingSprite = sprite
					}
				}
			}
		}
	}

	// Handle mouse pressed
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && ui.draggingSprite != nil {
		mx, my := ebiten.CursorPosition()
		ui.draggingSprite.MoveCentered(mx, my)
	}

	// Handle mouse released
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && ui.draggingSprite != nil {
		mx, my := ebiten.CursorPosition()
		for _, slot := range ui.playerSlots {
			if slot.rect.Collides(float32(mx), float32(my)) {
				ui.draggingSprite.SetPosition(slot.index)
				idx, draggingFish := ui.sim.GetFishByID(ui.draggingSprite.Id.String())
				//fmt.Printf("idx:%v fishID: %v\n", idx, draggingFish.Id.String())
				targetSlot := slot.index
				if draggingFish != nil {
					if ui.draggedFromInventory {
						ui.sim.Player_GetFish().AddFish(draggingFish, targetSlot)
						ui.draggingSprite.toolTip.ChangeAlignment(shapes.LeftAlignment)
						ui.draggedFromInventory = false
					} else {
						ui.sim.Player_GetFish().MoveFish(idx, targetSlot)
					}
				}

				ui.draggingSprite.Dragging = false
				ui.draggingSprite = nil
				return
			}
		}

		if ui.inventory.Rect.Collides(float32(mx), float32(my)) {
			ui.sim.Player_StoreExistingFish(ui.draggingSprite.Id.String())
			ui.draggingSprite.Dragging = false
			ui.draggingSprite.toolTip.ChangeAlignment(shapes.BottomAlignment)
			ui.draggingSprite = nil
		} else {
			ui.draggingSprite.ResetToPositionBeforeDrag()
			ui.draggedFromInventory = false
			ui.draggingSprite.Dragging = false
			ui.draggingSprite = nil
		}
	}
}

func (ui *UI) updateEncounterFish() {
	for index, fish := range ui.sim.Encounter_GetFish().GetAllFish() {
		if fish != nil {
			id := fish.Id
			if fish.IsDead() {
				delete(ui.sprites, id.String())
			} else if ui.sprites[id.String()] == nil {
				sprite := NewEncounterFishSprite(fish, index)
				ui.sprites[id.String()] = sprite
			}
		}
	}
}

func (ui *UI) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
	for _, slot := range ui.playerSlots {
		slot.Draw(screen)
	}
	for _, slot := range ui.encounterSlots {
		slot.Draw(screen)
	}
	ui.startSimBtn.Draw(screen)
	ui.stopSimBtn.Draw(screen)
	ui.inventory.Draw(screen)

	if ui.enabled {
		for _, sprite := range ui.sprites {
			sprite.Draw(screen)
		}
		for _, line := range ui.attackLines {
			line.Draw(screen)
		}
		for _, sprite := range ui.sprites {
			sprite.DrawToolTip(screen)
		}
		for _, dialog := range ui.dialogs {
			dialog.Draw(screen)
		}
	}
}
func (ui *UI) slotIndexToScreenPos(index int, leftSide bool) (int, int) {
	var xPos float32
	if leftSide {
		xPos = float32(ENV.Config.Get("slot.playerColX").(int))
	} else {
		xPos = float32(ENV.Config.Get("slot.encounterColX").(int))
	}
	yPadding := ENV.Config.Get("slot.topPad").(int)
	betweenSlotPadding := ENV.Config.Get("slot.betweenPad").(int)
	spriteSizePx := ENV.Config.Get("sprite.sizeInPx").(int)
	spriteScale := ENV.Config.Get("sprite.scale").(float64)
	height := float64(spriteSizePx) * spriteScale
	yPos := float32(index)*float32(int(height)+betweenSlotPadding) + float32(yPadding)
	return int(xPos), int(yPos)
}

// event handlers
func (ui *UI) handleEnableUiEvent(event environment.Event) {
	ui.enabled = true
}
func (ui *UI) handleDisableUiEvent(event environment.Event) {
	ui.enabled = false
}
func (ui *UI) handleFishAttackedEvent(event environment.Event) {
	attackedEvent := event.Data.(environment.FishAttackedEvent)
	// attackedEvent.Type // todo - add type to line effect
	sourceIndex, sourceFish := ui.sim.GetFishByID(attackedEvent.SourceId.String())
	targetIndex, _ := ui.sim.GetFishByID(attackedEvent.TargetId.String())

	if targetIndex == 999 || sourceFish == nil {
		return
	}
	// determine if the SourceID is a player fish or an encounter fish
	isSourcePlayerOwned := ui.sim.IsPlayerFish(attackedEvent.SourceId.String())
	sourceX, sourceY := ui.slotIndexToScreenPos(sourceIndex, isSourcePlayerOwned)

	// determine if the TargetID is a player fish or an encounter fish
	isTargetPlayerOwned := ui.sim.IsPlayerFish(attackedEvent.TargetId.String())
	targetX, targetY := ui.slotIndexToScreenPos(targetIndex, isTargetPlayerOwned)
	al := NewAttackLine(sourceX, sourceY, targetX, targetY, float32(sourceFish.Stats.MaxDuration))
	ui.attackLines = append(ui.attackLines, al)
}

// Just do one-off things that need to happen when the encounter starts
func (ui *UI) handleNextEncounterEvent(event environment.Event) {
	previousEncType := encounter.TypeFromString(event.Data.(environment.NextEncounterEvent).EncounterType)
	switch previousEncType {
	case encounter.EncounterTypeInitial:
		enc, err := ui.encounterMgr.GetCurrent()
		if err != nil {
			log.Fatalf("unable to get current encounter")
		}
		rewards := enc.GetRewards()
		if len(rewards) > 0 {
			dlg := NewInitialDialog(enc, ui.sim)
			ui.dialogs = append(ui.dialogs, dlg)
		}
	case encounter.EncounterTypeBattle: // set up next battle
		enc, err := ui.encounterMgr.GetNext()
		if err != nil {
			log.Fatalf("unable to get next encounter")
		}
		ui.sim.Encounter_SetFish(enc.GetCollection())

	}

}
