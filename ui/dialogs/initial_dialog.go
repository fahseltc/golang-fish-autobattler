package dialogs

import (
	"fishgame/encounter"
	"fishgame/simulation/simulation"
	"fishgame/ui/elements"
	"fishgame/ui/shapes"
	"fishgame/ui/util"

	"github.com/hajimehoshi/ebiten/v2"
)

type Initial struct {
	bg            *ebiten.Image
	rect          *shapes.Rectangle
	btns          []*elements.Button
	selectionMade bool
}

func NewInitialDialog(rewards []*encounter.Reward, sim simulation.SimulationInterface) *Initial {
	bg := util.LoadImage("ui/panel/grey_panel.png")
	rect := &shapes.Rectangle{X: 150, Y: 50, W: 500, H: 500}
	scaled := util.ScaleImage(bg, rect.W, rect.H)
	dlg := &Initial{
		bg:            scaled,
		rect:          rect,
		selectionMade: false,
	}

	btn := elements.NewButton(
		elements.WithRect(shapes.Rectangle{X: 250, Y: 270, W: 75, H: 50}),
		elements.WithText("Start"),
		elements.WithClickFunc(func() {
			reg := sim.GetFishRegistry()
			fish, _ := reg.GetFish(rewards[0].Fish[0])
			sim.Player_GetInventory().Add(fish)
			dlg.selectionMade = true
		}),
		elements.WithCenteredPos(),
	)
	dlg.btns = append(dlg.btns, btn)
	return dlg
}

func (ini *Initial) Update() {
	for _, btn := range ini.btns {
		btn.Update()
	}
}

func (ini *Initial) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(ini.rect.X), float64(ini.rect.Y))
	screen.DrawImage(ini.bg, opts)

	for _, btn := range ini.btns {
		btn.Draw(screen)
	}
}

func (ini *Initial) IsCompleted() bool {
	return ini.selectionMade
}
