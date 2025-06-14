package scene

import (
	"fishgame/shared/environment"
	"fishgame/ui/ui"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

var ENV *environment.Env

type Manager struct {
	Scenes  map[string]Scene
	Current Scene
}

func NewSceneManager(Env *environment.Env) *Manager {
	ENV = Env
	manager := &Manager{}
	ui.InitEnv(ENV)
	manager.Scenes = make(map[string]Scene)
	//	manager.Scenes["Play"] = nil
	//	manager.Scenes["GameOver"] = nil
	ENV.EventBus.Subscribe("GameOverEvent", manager.handleGameOverEvent)

	return manager
}

func (sm *Manager) SwitchTo(scene string, destroyOld bool) {
	newScene := sm.Scenes[scene]
	if destroyOld && sm.Current != nil {
		sm.Current.Destroy()
	}
	if scene == "Play" {
		playScene := NewPlayScene(sm)
		sm.Scenes[playScene.GetName()] = playScene
		newScene = playScene
	}
	if scene == "Menu" {
		menuScene := NewMenuScene(sm)
		sm.Scenes[menuScene.GetName()] = menuScene
		newScene = menuScene
	}
	if scene == "GameOver" {
		gameOverScene := NewGameOverScene(ENV, sm)
		sm.Scenes[gameOverScene.GetName()] = gameOverScene
		newScene = gameOverScene
	}
	sm.Current = newScene
}

func (sm *Manager) Update(dt float64) error {
	if sm.Current != nil {
		sm.Current.Update(dt)
		return nil
	}
	return fmt.Errorf("scenemanager has no current scene to update")
}

func (sm *Manager) Draw(screen *ebiten.Image) {
	// todo: handle transitions between scenes
	if sm.Current != nil {
		sm.Current.Draw(screen)
	}
}

// Event handlers

func (sm *Manager) handleGameOverEvent(event environment.Event) {

}
