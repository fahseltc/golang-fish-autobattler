package game

import (
	"fishgame/environment"
	"fishgame/scene"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Env            *environment.Env
	SceneManager   *scene.Manager
	LastUpdateTime time.Time
}

func NewGame(env *environment.Env) *Game {
	return &Game{
		Env:          env,
		SceneManager: scene.NewSceneManager(env),
	}
}

func (g *Game) Update() error {
	if g.LastUpdateTime.IsZero() {
		g.LastUpdateTime = time.Now()
	}

	dt := time.Since(g.LastUpdateTime).Seconds()
	err := g.SceneManager.Update(dt) // todo: handle errors and pass up?
	if err != nil {
		g.Env.Logger.Error("Error updating scene", "error", err)
		return err
	}
	g.LastUpdateTime = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.SceneManager.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 800, 600
}
