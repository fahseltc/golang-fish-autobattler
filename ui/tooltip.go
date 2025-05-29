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
	GetAlignment() Alignment
	GetRect() *Rectangle
}

type Tooltip struct {
	rect Rectangle
	bg   *ebiten.Image
}

type ItemTooltip struct {
	tt        *Tooltip
	font      text.Face
	item      *item.Item
	alignment Alignment
}

func NewItemToolTip(x, y, w, h float32, item *item.Item, alignment Alignment) *ItemTooltip {
	font, _ := util.LoadFont(20)

	centeredX := x - 0.5*w
	centeredY := y + 0.3*h

	tt := &ItemTooltip{
		tt: &Tooltip{
			rect: Rectangle{
				X: centeredX,
				Y: centeredY,
				W: w,
				H: h,
			},
			bg: util.LoadImage(nil, "assets/ui/panel/grey_panel.png"),
		},
		alignment: alignment,
		font:      font,
		item:      item,
	}
	return tt
}

func (i *ItemTooltip) OnHover(screen *ebiten.Image) {
	// draw bg

	if i.tt.bg != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(i.tt.rect.W)/float64(i.tt.bg.Bounds().Dx()), float64(i.tt.rect.H)/float64(i.tt.bg.Bounds().Dy()))
		op.GeoM.Translate(float64(i.tt.rect.X), float64(i.tt.rect.Y))

		screen.DrawImage(i.tt.bg, op)
	}

	titleX := float32(i.tt.rect.X) + float32(i.tt.rect.H)*0.7
	titleY := float32(i.tt.rect.Y) + float32(i.tt.rect.W)*0.05

	txtColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}

	bigFont, _ := util.LoadFont(30)
	txtSpacing := float32(20)

	DrawCenteredText(screen, bigFont, i.item.Name, int(titleX), int(titleY+10), txtColor)

	// HP
	hpString := fmt.Sprintf("HP: %v/%v", i.item.CurrentLife, i.item.Life)
	DrawCenteredText(screen, i.font, hpString, int(titleX), int(titleY+(txtSpacing*3)), txtColor)
	// DPS
	dpsString := fmt.Sprintf("DPS: %.2f", i.item.Dps())
	DrawCenteredText(screen, i.font, dpsString, int(titleX), int(titleY+(txtSpacing*4)), txtColor)
	// Damage
	dmgString := fmt.Sprintf("Damage: %v", i.item.Damage)
	DrawCenteredText(screen, i.font, dmgString, int(titleX), int(titleY+(txtSpacing*5)), txtColor)
	// Duration
	durationString := fmt.Sprintf("Duration: %v", i.item.Duration)
	DrawCenteredText(screen, i.font, durationString, int(titleX), int(titleY+(txtSpacing*6)), txtColor)
	// Type
	typeString := fmt.Sprintf("%v", strings.Title(i.item.Type.String()))
	DrawCenteredText(screen, i.font, typeString, int(titleX), int(titleY+(txtSpacing*7)), txtColor)
	// Description
	DrawCenteredText(screen, i.font, i.item.Description, int(titleX), int(titleY+(txtSpacing*8)), txtColor)
}

func (i *ItemTooltip) GetAlignment() Alignment {
	return i.alignment
}

func (i *ItemTooltip) GetRect() *Rectangle {
	return &i.tt.rect
}
