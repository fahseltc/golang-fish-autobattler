package encounter

import (
	"fishgame/environment"
	"fishgame/util"

	"github.com/hajimehoshi/ebiten/v2"
)

// type Encounter interface {
// 	Update()
// 	Draw()
// }

// const (
// 	Initial Type = iota
// 	Shop
// 	Choice
// 	Battle
// )

type Encounter struct {
	encType Type
	bg      *ebiten.Image
}

func NewEncounter(env *environment.Env) *Encounter {
	return &Encounter{
		encType: Initial,
		bg:      util.LoadImage(*env, "assets/bg/ocean.png"),
	}
}

func (enc *Encounter) Update() {

}

func (enc *Encounter) Draw(screen *ebiten.Image) {
	screen.DrawImage(enc.bg, nil)
	// background
	// switch based on type
}
