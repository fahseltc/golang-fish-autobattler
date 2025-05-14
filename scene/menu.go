package scene

import (
	"fishgame/environment"
	"fishgame/ui"
	"fishgame/util"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Menu struct {
	env          *environment.Env
	SceneManager *Manager
	bg           *ebiten.Image
	StartBtn     *ui.Button
}

func NewMenuScene(sm *Manager) *Menu {
	btn := ui.NewButton(sm.Env, 400, 450, 200, 100, "Start", color.Black, 30)
	btn.OnClick = func() {
		sm.SwitchTo("Play", true)
	}
	menu := &Menu{
		env:          sm.Env,
		SceneManager: sm,
		bg:           util.LoadImage(*sm.Env, "assets/bg/menu.png"),
		StartBtn:     btn,
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
