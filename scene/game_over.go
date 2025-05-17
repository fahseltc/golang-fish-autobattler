package scene

import (
	"fishgame/environment"
	"fishgame/ui"
	"fishgame/util"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update(dt float64)
// Draw(screen *ebiten.Image)
// Destroy()
// GetName() string

type GameOver struct {
	env          *environment.Env
	bg           *ebiten.Image
	SceneManager *Manager
	restartBtn   *ui.Button
	menuBtn      *ui.Button
}

func NewGameOverScene(env *environment.Env, sm *Manager) *GameOver {
	restartBtn := ui.NewButton(env, 200, 450, 150, 150, "Restart", color.Black, 20)
	restartBtn.OnClick = func() {
		sm.SwitchTo("Play", false)
	}
	menuBtn := ui.NewButton(env, 600, 450, 150, 150, "Menu", color.Black, 20)
	menuBtn.OnClick = func() {
		sm.SwitchTo("Menu", false)
	}

	g := &GameOver{
		env:          env,
		bg:           util.LoadImage(env, "assets/bg/game_over.png"),
		SceneManager: sm,
		restartBtn:   restartBtn,
		menuBtn:      menuBtn,
	}
	return g
}

func (g *GameOver) Update(dt float64) {
	g.restartBtn.Update()
	g.menuBtn.Update()
}
func (g *GameOver) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bg, nil)
	g.restartBtn.Draw(screen)
	g.menuBtn.Draw(screen)
}
func (g *GameOver) Destroy() {}
func (g *GameOver) GetName() string {
	return "GameOver"
}
