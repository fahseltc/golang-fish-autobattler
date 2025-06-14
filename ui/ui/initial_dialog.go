package ui

import (
	"fishgame/encounter"
	"fishgame/simulation/fish"
	"fishgame/simulation/simulation"
	"fishgame/ui/shapes"
	"fishgame/ui/util"
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Initial struct {
	bg            *ebiten.Image
	title         string
	rect          *shapes.Rectangle
	fishImgs      []*ebiten.Image
	btns          []*Button
	selectionMade bool
}

func NewInitialDialog(enc encounter.EncounterInterface, sim simulation.SimulationInterface) *Initial {
	bg := util.LoadImage("ui/panel/grey_panel.png")
	rect := &shapes.Rectangle{X: 100, Y: 100, W: 600, H: 400}
	scaled := util.ScaleImage(bg, rect.W, rect.H)
	dlg := &Initial{
		bg:            scaled,
		title:         enc.GetTitle(),
		rect:          rect,
		selectionMade: false,
	}

	reg := sim.GetFishRegistry()

	for index, reward := range enc.GetRewards() {
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
				dlg.selectionMade = true
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
		dlg.btns = append(dlg.btns, btn)
	}
	return dlg
}

func getFishImage(fish *fish.Fish) *ebiten.Image {
	spriteScale := ENV.Config.Get("sprite.scale").(float64)
	img := util.LoadImage(fmt.Sprintf("fish/%v.png", strings.ToLower(fish.Name)))
	if img == nil {
		img = util.LoadImage("TEXTURE_MISSING.png")
	}
	w, h := img.Size()
	scaled := ebiten.NewImage(int(float64(w)*spriteScale), int(float64(h)*spriteScale))
	op := &ebiten.DrawImageOptions{} // Draw original onto the new image with scaling

	scaled.DrawImage(img, op)
	return scaled
}

func (ini *Initial) createBtn() {

}

func (ini *Initial) Update() {
	for _, btn := range ini.btns {
		btn.Update()
	}
}

func (ini *Initial) Draw(screen *ebiten.Image) {
	// Draw BG
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(ini.rect.X), float64(ini.rect.Y))
	screen.DrawImage(ini.bg, opts)
	// Draw Title
	util.DrawCenteredText(screen, ENV.Fonts.Med, ini.title, 400, 200, color.RGBA{R: 0, G: 0, B: 0, A: 255})

	// Draw buttons
	for index, btn := range ini.btns {
		btn.Draw(screen)
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(btn.rect.X+0.2*btn.rect.W), float64(btn.rect.Y-80))
		screen.DrawImage(ini.fishImgs[index], opts)
	}
}

func (ini *Initial) IsCompleted() bool {
	return ini.selectionMade
}
