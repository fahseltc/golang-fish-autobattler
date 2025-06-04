package game

import (
	//"fishgame/scene"
	"fishgame/shared/environment"
	"fishgame/simulation/collection"
	"fishgame/simulation/fish"
	"fishgame/simulation/player"
	"fishgame/simulation/simulation"
	"fishgame/ui/ui"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var ENV *environment.Env

type Game struct {
	Env *environment.Env
	//SceneManager   *scene.Manager
	LastUpdateTime time.Time
	sim            simulation.SimulationInterface
	ui             *ui.UI
}

func NewGame(env *environment.Env) *Game {
	ENV = env
	player := player.NewPlayer(ENV, "player1")
	player.Fish = collection.NewCollection(ENV)
	player.Fish.AddFish(fish.NewFish(ENV, "fish1", "testfish1", fish.NewWeaponStats(20, 1, 5)), 0)
	player.Fish.AddFish(fish.NewFish(ENV, "fish2", "testfish2", fish.NewWeaponStats(30, 1, 5)), 1)
	player.Fish.AddFish(fish.NewFish(ENV, "fish3", "testfish3", fish.NewWeaponStats(10, 2, 10)), 2)

	encounterFish := collection.NewCollection(ENV)
	encounterFish.AddFish(fish.NewFish(ENV, "fish4", "testfish4", fish.NewWeaponStats(100, 2, 5)), 0)
	encounterFish.AddFish(fish.NewFish(ENV, "fish5", "testfish5", fish.NewWeaponStats(35, 3, 20)), 4)
	encounterFish.AddFish(fish.NewFish(ENV, "fish6", "testfish6", fish.NewWeaponStats(35, 3, 20)), 2)
	sim := simulation.NewSimulation(ENV, player, encounterFish)
	sim.Enable()
	return &Game{
		Env: env,
		sim: sim,
		ui:  ui.NewUI(ENV, sim),
		//SceneManager: scene.NewSceneManager(env),
	}
}

func (g *Game) Update() error {
	if g.LastUpdateTime.IsZero() {
		g.LastUpdateTime = time.Now()
	}

	dt := time.Since(g.LastUpdateTime).Seconds()
	g.sim.Update(dt)
	g.ui.Update()

	g.LastUpdateTime = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 800, 600
}
