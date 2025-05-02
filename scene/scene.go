package scene

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Init(sm *Manager)
	Update(dt float64) error
	Draw(screen *ebiten.Image)
	Destroy()
	GetName() string
}
