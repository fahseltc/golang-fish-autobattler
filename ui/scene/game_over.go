package scene

import (
	"fishgame/shared/environment"
	"fishgame/ui/elements"
	"fishgame/ui/shapes"
	"fishgame/ui/util"

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
	restartBtn   *elements.Button
	menuBtn      *elements.Button
}

func NewGameOverScene(env *environment.Env, sm *Manager) *GameOver {
	restartBtn := elements.NewButton(
		elements.WithRect(shapes.Rectangle{X: 200, Y: 450, W: 150, H: 150}),
		elements.WithText("Restart"),
		elements.WithClickFunc(func() {
			sm.SwitchTo("Play", false)
		}),
		elements.WithCenteredPos(),
	)

	menuBtn := elements.NewButton(
		elements.WithRect(shapes.Rectangle{X: 600, Y: 450, W: 150, H: 150}),
		elements.WithText("Menu"),
		elements.WithClickFunc(func() {
			sm.SwitchTo("Menu", false)
		}),
		elements.WithCenteredPos(),
	)

	g := &GameOver{
		env:          env,
		bg:           util.LoadImage("bg/game_over.png"),
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
