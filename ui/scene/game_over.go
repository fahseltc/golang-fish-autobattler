package scene

// import (
// 	"fishgame/shared/environment"
// 	"fishgame/ui/btn"
// 	"fishgame/ui/shapes"
// 	"fishgame/ui/util"

// 	"github.com/hajimehoshi/ebiten/v2"
// )

// // Update(dt float64)
// // Draw(screen *ebiten.Image)
// // Destroy()
// // GetName() string

// type GameOver struct {
// 	env          *environment.Env
// 	bg           *ebiten.Image
// 	SceneManager *Manager
// 	restartBtn   *btn.Button
// 	menuBtn      *btn.Button
// }

// func NewGameOverScene(env *environment.Env, sm *Manager) *GameOver {
// 	restartBtn := btn.NewButton(
// 		btn.WithRect(shapes.Rectangle{X: 200, Y: 450, W: 150, H: 150}),
// 		btn.WithText("Restart"),
// 		btn.WithClickFunc(func() {
// 			sm.SwitchTo("Play", false)
// 		}),
// 		btn.WithCenteredPos(),
// 	)

// 	menuBtn := btn.NewButton(
// 		btn.WithRect(shapes.Rectangle{X: 600, Y: 450, W: 150, H: 150}),
// 		btn.WithText("Menu"),
// 		btn.WithClickFunc(func() {
// 			sm.SwitchTo("Menu", false)
// 		}),
// 		btn.WithCenteredPos(),
// 	)

// 	g := &GameOver{
// 		env:          env,
// 		bg:           util.LoadImage(env, "assets/bg/game_over.png"),
// 		SceneManager: sm,
// 		restartBtn:   restartBtn,
// 		menuBtn:      menuBtn,
// 	}
// 	return g
// }

// func (g *GameOver) Update(dt float64) {
// 	g.restartBtn.Update()
// 	g.menuBtn.Update()
// }
// func (g *GameOver) Draw(screen *ebiten.Image) {
// 	screen.DrawImage(g.bg, nil)
// 	g.restartBtn.Draw(screen)
// 	g.menuBtn.Draw(screen)
// }
// func (g *GameOver) Destroy() {}
// func (g *GameOver) GetName() string {
// 	return "GameOver"
// }
