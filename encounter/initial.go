package encounter

import "github.com/hajimehoshi/ebiten/v2"

type Initial struct {
	text string
}

func (i Initial) Update() {
}

func (i *Initial) Draw(screen *ebiten.Image) {
	// get 3 starter fishes, put buttons of them on screen, wait for user to click a button

}
