package scene

// import (
// 	"fishgame/environment"
// 	"fishgame/shapes"
// 	"fishgame/ui"
// 	"fishgame/util"

// 	"github.com/hajimehoshi/ebiten/v2"
// )

// type Menu struct {
// 	env          *environment.Env
// 	SceneManager *Manager
// 	bg           *ebiten.Image
// 	startBtn     *ui.Button
// }

// func NewMenuScene(sm *Manager) *Menu {
// 	menu := &Menu{
// 		env:          sm.Env,
// 		SceneManager: sm,
// 		bg:           util.LoadImage(sm.Env, "assets/bg/menu.png"),
// 	}

// 	menu.startBtn = ui.NewButton(
// 		ui.WithRect(shapes.Rectangle{X: 400, Y: 450, W: 250, H: 100}),
// 		ui.WithText("StartGame"),
// 		ui.WithClickFunc(func() {
// 			sm.SwitchTo("Play", false)
// 		}),
// 		ui.WithCenteredPos(),
// 	)

// 	return menu
// }

// func (m *Menu) Update(dt float64) {
// 	m.startBtn.Update()
// }

// func (m *Menu) Draw(screen *ebiten.Image) {
// 	screen.DrawImage(m.bg, &ebiten.DrawImageOptions{})
// 	m.startBtn.Draw(screen)
// }

// func (m *Menu) Destroy() {
// 	// Clean up menu resources here
// }

// func (m *Menu) GetName() string {
// 	return "Menu"
// }
