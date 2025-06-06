package ui

import (
	"fishgame/shared/environment"
	"fishgame/simulation/simulation"
	images "fishgame/ui/images"
	"fishgame/ui/shapes"
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var ENV *environment.Env

type UI struct {
	Env           *environment.Env
	imageRegistry *images.Registry

	sim simulation.SimulationInterface

	sprites     map[string]*Sprite
	playerSlots map[int]*Slot
	attackLines []*AttackLine

	startSimBtn *Button
	stopSimBtn  *Button

	enabled bool

	draggingSprite *Sprite
}

func NewUI(env *environment.Env, sim simulation.SimulationInterface) *UI {
	ENV = env

	ui := &UI{
		Env:           env,
		imageRegistry: images.NewRegistry(),
		sim:           sim,
		sprites:       make(map[string]*Sprite),
		playerSlots:   make(map[int]*Slot),
		enabled:       false,
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
	}

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
	ENV.EventBus.Subscribe("EnableUiEvent", ui.HandleEnableUiEvent)

	return ui
}

func (ui *UI) Update(dt float64) {
	ui.startSimBtn.Update()
	ui.stopSimBtn.Update()
	if ui.enabled {
		ui.updatePlayerFish()
		ui.updateEncounterFish()
		ui.updateSpritePositionsFromSim()

		for i, line := range ui.attackLines {
			if line.Duration == 0 {
				line.enabled = false
				ui.attackLines = append(ui.attackLines[:i], ui.attackLines[i+1:]...)
			}
			line.Update(dt)
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

func (ui *UI) updatePlayerFish() {
	for index, fish := range ui.sim.Player_GetFish().GetAllFish() {
		if fish != nil {
			id := fish.Id
			if fish.IsDead() { // the sim fish is dead, remove its sprite
				delete(ui.sprites, id.String())
			} else if ui.sprites[id.String()] == nil { // the sim fish is new and needs a sprite made
				sprite := NewPlayerFishSprite(ui.imageRegistry, fish, index)
				ui.sprites[id.String()] = sprite
			} else if ui.sprites[id.String()] != nil { // the sim fish is already added to the list of sprites
				sprite := ui.sprites[id.String()]

				// Handle Clicks
				if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && !ui.sim.IsEnabled() {
					mx, my := ebiten.CursorPosition()
					if sprite.Rect.Collides(float32(mx), float32(my)) {
						fmt.Println("click collides")
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
				//fmt.Printf("released into slot: %v\n", slot.index)
				// if ui.draggingSprite.Rect.Collides(float32(mx), float32(my)) {
				// 	fmt.Printf("it collides with itself, do nothing\n")
				// }

				ui.draggingSprite.SetPosition(slot.index)
				idx, draggingFish := ui.sim.Player_GetFish().ById(ui.draggingSprite.Id.String())
				//fmt.Printf("idx:%v fishID: %v\n", idx, draggingFish.Id.String())
				targetSlot := slot.index
				if draggingFish != nil {
					ui.sim.Player_GetFish().MoveFish(idx, targetSlot)
					// if the slot had a fish, we need to move its sprite
				}

				ui.draggingSprite.Dragging = false
				ui.draggingSprite = nil
				return
			}
		}
		// if mx,my are on top of slots or inventory, place the item there
		// otherwise, put it bakc where it came from
		fmt.Println("released")
		ui.draggingSprite.ResetToPositionBeforeDrag()
		ui.draggingSprite.Dragging = false
		ui.draggingSprite = nil
	}

}

func (ui *UI) updateEncounterFish() {
	for index, fish := range ui.sim.Encounter_GetFish().GetAllFish() {
		if fish != nil {
			id := fish.Id
			if fish.IsDead() {
				delete(ui.sprites, id.String())
			} else if ui.sprites[id.String()] == nil {
				sprite := NewEncounterFishSprite(ui.imageRegistry, fish, index)
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
	ui.startSimBtn.Draw(screen)
	ui.stopSimBtn.Draw(screen)
	if ui.enabled {
		for _, sprite := range ui.sprites {
			sprite.Draw(screen)
		}
		for _, line := range ui.attackLines {
			line.Draw(screen)
		}
	}
}

// event handlers
func (ui *UI) HandleEnableUiEvent(event environment.Event) {
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
	sourceX, sourceY := ui.slotIndexToScreenPos(sourceIndex, true)  // TODO - need ower to tell whether its left or right side.
	targetX, targetY := ui.slotIndexToScreenPos(targetIndex, false) // TODO - need ower to tell whether its left or right side.

	al := NewAttackLine(sourceX, sourceY, targetX, targetY, float32(sourceFish.Stats.MaxDuration))
	ui.attackLines = append(ui.attackLines, al)
}

func (ui *UI) slotIndexToScreenPos(index int, leftSide bool) (int, int) {
	var xPos float32
	if leftSide {
		xPos = float32(ENV.Config.Get("playerSlotColumnX").(int))
	} else {
		xPos = float32(ENV.Config.Get("encounterSlotColumnX").(int))
	}
	yPadding := ENV.Config.Get("slotYpadding").(int)
	betweenSlotPadding := ENV.Config.Get("betweenSlotPadding").(int)
	spriteSizePx := ENV.Config.Get("spriteSizePx").(int)
	spriteScale := ENV.Config.Get("spriteScale").(float64)
	height := float64(spriteSizePx) * spriteScale
	yPos := float32(index)*float32(int(height)+betweenSlotPadding) + float32(yPadding)
	return int(xPos), int(yPos)
}
