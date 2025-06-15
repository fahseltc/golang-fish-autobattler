package ui

import (
	"fishgame/simulation/player"
	"fishgame/ui/util"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Currency struct {
	//bg     *ebiten.Image
	sprite *Sprite
	player *player.Player
}

func NewCurrency(player *player.Player) *Currency {
	invent := &Currency{
		//bg:     util.LoadImage("ui/panel/grey_panel.png"),
		sprite: NewCurrencySprite(),
		player: player,
	}
	return invent
}

func (curr *Currency) Draw(screen *ebiten.Image) {
	curr.sprite.Draw(screen)
	font := ENV.Fonts.Large

	util.DrawCenteredText(screen, font, fmt.Sprintf("%v", curr.player.GetCurrencyAmount()), int(curr.sprite.Rect.X+75), int((curr.sprite.Rect.H/2)+12), nil)

}
