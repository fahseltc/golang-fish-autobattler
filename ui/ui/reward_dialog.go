package ui

import (
	"fishgame/encounter"
	"fishgame/shared/environment"
	"fishgame/simulation/fish"
	"fishgame/simulation/simulation"
	"fishgame/ui/shapes"
	"fishgame/ui/util"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type RewardDialog struct {
	bg            *ebiten.Image
	title         string
	rect          *shapes.Rectangle
	fishImgs      []*ebiten.Image
	btns          []*Button
	selectionMade bool
}

func NewRewardDialog(enc encounter.EncounterInterface, sim simulation.SimulationInterface) *RewardDialog {
	bg := util.LoadImage("ui/panel/grey_panel.png")
	rect := &shapes.Rectangle{X: 100, Y: 100, W: 600, H: 400}
	scaled := util.ScaleImage(bg, rect.W, rect.H)
	dlg := &RewardDialog{
		bg:            scaled,
		title:         enc.GetTitle(),
		rect:          rect,
		selectionMade: false,
	}

	reg := sim.GetFishRegistry()

	for index, reward := range enc.GetRewards() {
		if len(reward.Fish) > 0 {
			btn := dlg.createFishRewardBtn(index, reg, reward, enc, sim)
			dlg.btns = append(dlg.btns, btn)
		}
		if reward.Currency != 0 {
			// add currenct btns
		}

	}
	return dlg
}

func (dlg *RewardDialog) createFishRewardBtn(index int, reg *fish.FishStatsRegistry, reward *encounter.Reward, enc encounter.EncounterInterface, sim simulation.SimulationInterface) *Button {
	firstFish, _ := reg.GetFish(reward.Fish[0])

	img := getFishImage(firstFish)
	dlg.fishImgs = append(dlg.fishImgs, img)

	xOffset := float32(index * 150)
	btnRect := shapes.Rectangle{X: 250 + xOffset, Y: 350, W: 125, H: 50}

	btn := NewButton(
		WithRect(btnRect),
		WithText(firstFish.Name),
		WithClickFunc(func() {
			for _, fishName := range enc.GetRewards()[index].Fish {
				fish, _ := reg.GetFish(fishName)
				sim.Player_GetInventory().Add(fish)
			}
			// TODO: handle currency rewards
			dlg.selectionMade = true
			ENV.EventBus.Publish(environment.Event{
				Type: "NextEncounterEvent",
				Data: environment.NextEncounterEvent{
					EncounterType: "reward",
				},
			})
		}),
		WithCenteredPos(),
		WithToolTip(
			NewFishToolTip(
				ENV,
				btnRect,
				shapes.BottomAlignment,
				firstFish,
			),
		),
	)
	return btn
}

func (dlg *RewardDialog) Update() {
	for _, btn := range dlg.btns {
		btn.Update()
	}
}

func (dlg *RewardDialog) Draw(screen *ebiten.Image) {
	// Draw BG
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(dlg.rect.X), float64(dlg.rect.Y))
	screen.DrawImage(dlg.bg, opts)
	// Draw Title
	util.DrawCenteredText(screen, ENV.Fonts.Med, dlg.title, 400, 200, color.RGBA{R: 0, G: 0, B: 0, A: 255})

	// Draw buttons
	for index, btn := range dlg.btns {
		btn.Draw(screen)
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(btn.rect.X+0.2*btn.rect.W), float64(btn.rect.Y-80))
		screen.DrawImage(dlg.fishImgs[index], opts)
	}
}

func (dlg *RewardDialog) IsCompleted() bool {
	return dlg.selectionMade
}
