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

	sprites map[string]*Sprite

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

	// StartSimulationEvent - no data associated
	// StopSimulationEvent - no data associated
	// FishAttackedEvent
	// FishDiedEvent
	// GameOverEvent
	// EncounterDoneEvent
	//ENV.EventBus.Subscribe("FishAttackedEvent", func(event environment.Event) {})
	//ENV.EventBus.Subscribe("FishDiedEvent", func(event environment.Event) {})
	//ENV.EventBus.Subscribe("GameOverEvent", func(event environment.Event) {})
	//ENV.EventBus.Subscribe("EncounterDoneEvent", func(event environment.Event) {})
	ENV.EventBus.Subscribe("DisableUiEvent", ui.handleDisableUiEvent)
	ENV.EventBus.Subscribe("EnableUiEvent", ui.HandleEnableUiEvent)
	return ui
}

func (ui *UI) Update() {
	ui.startSimBtn.Update()
	ui.stopSimBtn.Update()
	if ui.enabled {
		ui.updatePlayerFish()
		ui.updateEncounterFish()
	}
}

func (ui *UI) updatePlayerFish() {
	//var draggingIndex int

	for index, fish := range ui.sim.Player_GetFish().GetAllFish() {
		if fish != nil {
			id := fish.Id
			if fish.IsDead() {
				delete(ui.sprites, id.String())
			} else if ui.sprites[id.String()] == nil {
				sprite := NewPlayerFishSprite(ui.imageRegistry, fish, index)
				ui.sprites[id.String()] = sprite
			} else if ui.sprites[id.String()] != nil {
				// the fish shouldnt be removed and is already added
				sprite := ui.sprites[id.String()]
				if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && !ui.sim.IsEnabled() {
					mx, my := ebiten.CursorPosition()
					if sprite.Rect.Collides(float32(mx), float32(my)) {
						fmt.Println("click collides")
						sprite.Dragging = true
						sprite.SavePositionBeforeDrag()
						ui.draggingSprite = sprite

						//draggingIndex = index
					}
				}
			}
		}
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && ui.draggingSprite != nil {
		mx, my := ebiten.CursorPosition()
		ui.draggingSprite.MoveCentered(mx, my)
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && ui.draggingSprite != nil {
		//mx, my := ebiten.CursorPosition()
		// if mx,my are on top of slots or inventory, place the item there
		// otherwise, put it bakc where it came from
		fmt.Println("released")
		ui.draggingSprite.ResetToPositionBeforeDrag()
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
	ui.startSimBtn.Draw(screen)
	ui.stopSimBtn.Draw(screen)
	if ui.enabled {
		for _, sprite := range ui.sprites {
			sprite.Draw(screen)
		}
	}
}

// func (ui *UI) SetSimulation(sim simulation.SimulationInterface) {
// 	ui.sim = sim
// }

// event handlers
func (ui *UI) HandleEnableUiEvent(event environment.Event) {
	ui.enabled = true
}

func (ui *UI) handleDisableUiEvent(event environment.Event) {
	ui.enabled = false
}
