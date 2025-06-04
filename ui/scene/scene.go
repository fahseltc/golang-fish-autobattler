package scene

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	//Init()
	Update(dt float64)
	Draw(screen *ebiten.Image)
	Destroy()
	GetName() string
}
