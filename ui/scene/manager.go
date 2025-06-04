package scene

// import (
// 	"fishgame/environment"
// 	"fmt"

// 	"github.com/hajimehoshi/ebiten/v2"
// )

// type Manager struct {
// 	Env     *environment.Env
// 	Scenes  map[string]Scene
// 	Current Scene
// }

// func NewSceneManager(Env *environment.Env) *Manager {
// 	manager := &Manager{
// 		Env: Env,
// 	}
// 	manager.Init()
// 	return manager
// }

// func (sm *Manager) Init() {
// 	sm.Scenes = make(map[string]Scene)

// 	menuScene := NewMenuScene(sm)
// 	sm.Scenes[menuScene.GetName()] = menuScene

// 	gameOverScene := NewGameOverScene(sm.Env, sm)
// 	sm.Scenes[gameOverScene.GetName()] = gameOverScene

// 	sm.Current = menuScene
// }

// func (sm *Manager) SwitchTo(scene string, destroyOld bool) {
// 	newScene := sm.Scenes[scene]
// 	if destroyOld && sm.Current != nil {
// 		sm.Current.Destroy()
// 	}
// 	if scene == "Play" {
// 		playScene := &Play{Env: sm.Env}
// 		playScene.Init(sm)
// 		sm.Scenes[playScene.GetName()] = playScene
// 		newScene = playScene
// 	}
// 	sm.Current = newScene
// }

// func (sm *Manager) Update(dt float64) error {
// 	if sm.Current != nil {
// 		sm.Current.Update(dt)
// 		return nil
// 	}
// 	return fmt.Errorf("scenemanager has no current scene to update")
// }

// func (sm *Manager) Draw(screen *ebiten.Image) {
// 	// todo: handle transitions between scenes
// 	if sm.Current != nil {
// 		sm.Current.Draw(screen)
// 	}
// }
