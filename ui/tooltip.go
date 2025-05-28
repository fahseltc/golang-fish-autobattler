package ui

import (
	"fishgame/item"
	"fishgame/util"
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TooltipInterface interface {
	OnHover(screen *ebiten.Image)
}

type Tooltip struct {
	x  float32
	y  float32
	w  float32
	h  float32
	bg *ebiten.Image
}

type InitialTooltip struct {
	tt   *Tooltip
	font text.Face
	item *item.Item
}

func NewInitialTooltip(x, y, w, h float32, item *item.Item) *InitialTooltip {
	font, _ := util.LoadFont(20)

	centeredX := x - 0.5*w
	centeredY := y + 0.3*h

	tt := &InitialTooltip{
		tt: &Tooltip{
			x:  centeredX,
			y:  centeredY,
			w:  w,
			h:  h,
			bg: util.LoadImage(nil, "assets/ui/panel/grey_panel.png"),
		},
		font: font,
		item: item,
	}
	return tt
}

func (i *InitialTooltip) OnHover(screen *ebiten.Image) {
	// draw bg

	if i.tt.bg != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(i.tt.w)/float64(i.tt.bg.Bounds().Dx()), float64(i.tt.h)/float64(i.tt.bg.Bounds().Dy()))
		op.GeoM.Translate(float64(i.tt.x), float64(i.tt.y))

		screen.DrawImage(i.tt.bg, op)
	}

	titleX := float32(i.tt.x) + float32(i.tt.h)*0.7
	titleY := float32(i.tt.y) + float32(i.tt.w)*0.05

	txtColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}

	bigFont, _ := util.LoadFont(30)
	txtSpacing := float32(20)

	DrawCenteredText(screen, bigFont, i.item.Name, int(titleX), int(titleY+10), txtColor)

	hpString := fmt.Sprintf("HP: %v/%v", i.item.CurrentLife, i.item.Life)
	DrawCenteredText(screen, i.font, hpString, int(titleX), int(titleY+(txtSpacing*3)), txtColor)
	dpsString := fmt.Sprintf("DPS: %.2f", i.item.Dps())
	DrawCenteredText(screen, i.font, dpsString, int(titleX), int(titleY+(txtSpacing*4)), txtColor)
	dmgString := fmt.Sprintf("Damage: %v", i.item.Damage)
	DrawCenteredText(screen, i.font, dmgString, int(titleX), int(titleY+(txtSpacing*5)), txtColor)
	durationString := fmt.Sprintf("Duration: %v", i.item.Duration)
	DrawCenteredText(screen, i.font, durationString, int(titleX), int(titleY+(txtSpacing*6)), txtColor)

	typeString := fmt.Sprintf("%v", strings.Title(i.item.Type.String()))
	DrawCenteredText(screen, i.font, typeString, int(titleX), int(titleY+(txtSpacing*7)), txtColor)

	DrawCenteredText(screen, i.font, i.item.Description, int(titleX), int(titleY+(txtSpacing*8)), txtColor)
}
