package ui

import (
	"fishgame/shared/environment"
	"fishgame/simulation/simulation"
	images "fishgame/ui/images"
	"fishgame/ui/shapes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

var ENV *environment.Env

type UI struct {
	Env           *environment.Env
	imageRegistry *images.Registry

	sim simulation.SimulationInterface

	sprites map[string]*Sprite
}

func NewUI(env *environment.Env, sim simulation.SimulationInterface) *UI {
	ENV = env
	// StartSimulationEvent - no data associated
	// StopSimulationEvent - no data associated
	// FishAttackedEvent
	// FishDiedEvent
	// GameOverEvent
	// EncounterDoneEvent
	ENV.EventBus.Subscribe("FishAttackedEvent", func(event environment.Event) {})
	ENV.EventBus.Subscribe("FishDiedEvent", func(event environment.Event) {})
	ENV.EventBus.Subscribe("GameOverEvent", func(event environment.Event) {})
	ENV.EventBus.Subscribe("EncounterDoneEvent", func(event environment.Event) {})
	ui := &UI{
		Env:           env,
		imageRegistry: images.NewRegistry(),
		sim:           sim,
		sprites:       make(map[string]*Sprite),
	}
	return ui
}

func (ui *UI) Update() {
	ui.updatePlayerFish()
	ui.updateEncounterFish()
}

func (ui *UI) updatePlayerFish() {
	for index, fish := range ui.sim.Player_GetFish().GetAllFish() {
		if fish != nil {
			id := fish.Id
			if fish.IsDead() {
				delete(ui.sprites, id.String())
			} else if ui.sprites[id.String()] == nil {
				// its a new fish, make the sprite
				img := ui.imageRegistry.Images[fmt.Sprintf("fish/%v.png", fish.Name)]

				if img == nil {
					// load some dumb bright default img
					img = ui.imageRegistry.Images["TEXTURE_MISSING.png"]
				}
				// for player fish, they are in static locations based on the collection slot they are in - CALCULATABLE!
				sprite := NewSprite(&id, shapes.Rectangle{X: 1, Y: float32(index) * 128, W: 10, H: 10}, img)
				//sprite.setlocation ...
				ui.sprites[id.String()] = sprite
			}
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
				// its a new fish, make the sprite
				img := ui.imageRegistry.Images[fmt.Sprintf("fish/%v.png", fish.Name)]

				if img == nil {
					// load some dumb bright default img
					img = ui.imageRegistry.Images["TEXTURE_MISSING.png"]
				}
				// for player fish, they are in static locations based on the collection slot they are in - CALCULATABLE!
				sprite := NewSprite(&id, shapes.Rectangle{X: 600, Y: float32(index) * 128, W: 10, H: 10}, img)
				//sprite.setlocation ...
				ui.sprites[id.String()] = sprite
			}
		}
	}
}

func (ui *UI) Draw(screen *ebiten.Image) {
	for _, sprite := range ui.sprites {
		sprite.Draw(screen)
	}
}
