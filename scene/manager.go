package scene

import (
	"fishgame/environment"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Manager struct {
	Env     *environment.Env
	Scenes  map[string]Scene
	Current Scene
	Next    Scene
}

func NewSceneManager(Env *environment.Env) *Manager {
	manager := &Manager{
		Env: Env,
	}
	manager.Init()
	return manager
}

func (sm *Manager) Init() {
	sm.Scenes = make(map[string]Scene)

	menuScene := NewMenuScene(sm)
	sm.Scenes[menuScene.GetName()] = menuScene

	playScene := &Play{Env: sm.Env}
	playScene.Init(sm)
	sm.Scenes[playScene.GetName()] = playScene

	sm.Current = menuScene
	sm.Next = playScene
}

func (sm *Manager) SwitchTo(scene string, destroyOld bool) {
	newScene := sm.Scenes[scene]
	if newScene == nil {
		sm.Env.Logger.Error("Scene not found", "scene", scene)
		return
	}
	if destroyOld && sm.Current != nil {
		sm.Current.Destroy()
	}
	sm.Current = newScene
}

func (sm *Manager) SwitchToNext() {
	if sm.Next != nil {
		sm.Current.Destroy()
		sm.Current = sm.Next
		sm.Next = nil
	}
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
