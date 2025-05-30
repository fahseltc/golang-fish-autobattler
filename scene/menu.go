package scene

import (
	"fishgame/environment"
	"fishgame/shapes"
	"fishgame/ui"
	"fishgame/util"

	"github.com/hajimehoshi/ebiten/v2"
)

type Menu struct {
	env          *environment.Env
	SceneManager *Manager
	bg           *ebiten.Image
	StartBtn     *ui.Button
}

func NewMenuScene(sm *Manager) *Menu {
	startBtn := ui.NewButton(
		ui.WithRect(shapes.Rectangle{X: 400, Y: 450, W: 250, H: 100}),
		ui.WithText("Start"),
		ui.WithClickFunc(func() {
			sm.SwitchTo("Play", false)
		}),
		ui.WithCenteredPos(),
	)

	menu := &Menu{
		env:          sm.Env,
		SceneManager: sm,
		bg:           util.LoadImage(sm.Env, "assets/bg/menu.png"),
		StartBtn:     startBtn,
	}
	return menu
}

func (m *Menu) Update(dt float64) {
	m.StartBtn.Update()
}

func (m *Menu) Draw(screen *ebiten.Image) {
	screen.DrawImage(m.bg, &ebiten.DrawImageOptions{})
	m.StartBtn.Draw(screen)
}

func (m *Menu) Destroy() {
	// Clean up menu resources here
}

func (m *Menu) GetName() string {
	return "Menu"
}
