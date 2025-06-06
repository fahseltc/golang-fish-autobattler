package game

import (
	"fishgame/shared/environment"
	"fishgame/ui/scene"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var ENV *environment.Env

type Game struct {
	Env            *environment.Env
	LastUpdateTime time.Time
	sceneManager   *scene.Manager
}

func NewGame(env *environment.Env) *Game {
	ENV = env

	mgr := scene.NewSceneManager(env)
	return &Game{
		Env:          env,
		sceneManager: mgr,
	}
}

func (g *Game) Update() error {
	// Pt1: Calculate DT
	if g.LastUpdateTime.IsZero() {
		g.LastUpdateTime = time.Now()
	}
	dt := time.Since(g.LastUpdateTime).Seconds()

	// Update Game Objects
	if g.sceneManager.Current == nil {
		g.sceneManager.SwitchTo("Menu", false)
	}
	g.sceneManager.Current.Update(dt)

	// Pt2: Calculate DT
	g.LastUpdateTime = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Current.Draw(screen)
	//g.ui.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return ENV.Config.Get("internalResolutionW").(int), ENV.Config.Get("internalResolutionH").(int)
}
