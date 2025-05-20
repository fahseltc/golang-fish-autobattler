package ui

import (
	"fishgame/item"
	"fishgame/util"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Tooltip struct {
	btn *Button
	x   int
	y   int
	w   int
	h   int
}

type InitialTooltip struct {
	tt   *Tooltip
	font text.Face
	item *item.Item
}

func NewInitialTooltip(btn *Button, x, y, w, h int, item *item.Item) *InitialTooltip {
	font, _ := util.LoadFont(20)
	tt := &InitialTooltip{
		tt: &Tooltip{
			btn: btn,
			x:   x,
			y:   y,
			w:   w,
			h:   h,
		},
		font: font,
		item: item,
	}
	return tt
}

func (i *InitialTooltip) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(i.tt.x), float32(i.tt.y), float32(i.tt.w), float32(i.tt.h), color.RGBA{128, 128, 128, 255}, true)

	titleX := float32(i.tt.x) + float32(i.tt.h)*0.5
	titleY := float32(i.tt.y) + float32(i.tt.w)*0.15
	//DrawCenteredText(screen, i.tt.btn., slot.item.Name, int(titleX), int(titleY))

	ttstring := fmt.Sprintf("DPS: %.2f", i.item.Dps())
	DrawCenteredText(screen, i.font, ttstring, int(titleX), int(titleY+15))

	hpstring := fmt.Sprintf("HP: %v/%v", i.item.CurrentLife, i.item.Life)
	DrawCenteredText(screen, i.font, hpstring, int(titleX), int(titleY+30))

	DrawCenteredText(screen, i.font, i.item.Description, int(titleX), int(titleY+45))
}
