package ui

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/shapes"
	"fishgame/tooltip"
	"fishgame/util"
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TooltipInterface interface {
	OnHover(screen *ebiten.Image)
	GetAlignment() tooltip.Alignment
	GetRect() *shapes.Rectangle
}

type Tooltip struct {
	rect shapes.Rectangle
	bg   *ebiten.Image
}

func (tt *Tooltip) Draw(screen *ebiten.Image) {
	if tt.bg != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(tt.rect.W)/float64(tt.bg.Bounds().Dx()), float64(tt.rect.H)/float64(tt.bg.Bounds().Dy()))
		op.GeoM.Translate(float64(tt.rect.X), float64(tt.rect.Y))

		screen.DrawImage(tt.bg, op)
	}
}

type InitialToolTip struct {
	env        *environment.Env
	tt         *Tooltip
	font       text.Face
	item       *item.Item
	alignment  tooltip.Alignment
	hpIcon     *ebiten.Image
	damageIcon *ebiten.Image
}

func NewInitialToolTip(env *environment.Env, x, y, w, h float32, item *item.Item, alignment tooltip.Alignment) *InitialToolTip {
	font := env.Fonts.Med

	centeredX := x - 0.5*w
	centeredY := y + 0.3*h

	tt := &InitialToolTip{
		tt: &Tooltip{
			rect: shapes.Rectangle{
				X: centeredX,
				Y: centeredY,
				W: w,
				H: h,
			},
			bg: util.LoadImage(nil, "assets/ui/tooltip/grey_panel.png"),
		},
		hpIcon:     util.LoadImage(nil, "assets/ui/tooltip/shield.png"),
		damageIcon: util.LoadImage(nil, "assets/ui/tooltip/damage_panel.png"),
		alignment:  alignment,
		font:       font,
		item:       item,
	}
	return tt
}

func (i *InitialToolTip) OnHover(screen *ebiten.Image) {
	txtColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	i.tt.Draw(screen)

	// draw fish sprite
	if i.item.Sprite != nil {
		op := &ebiten.DrawImageOptions{}
		// op.GeoM.Scale(float64(i.tt.rect.W)/float64(i.item.Sprite.Bounds().Dx()), float64(i.tt.rect.H)/float64(i.item.Sprite.Bounds().Dy()))
		op.GeoM.Translate(float64(i.tt.rect.X+10), float64(i.tt.rect.Y))
		screen.DrawImage(i.item.Sprite, op)
	}

	// draw hp icon
	if i.hpIcon != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i.tt.rect.X+i.percentW(75)), float64(i.tt.rect.Y+i.percentW(3)))
		screen.DrawImage(i.hpIcon, op)
	}

	titleX := float32(i.tt.rect.X) + i.percentH(85)
	titleY := float32(i.tt.rect.Y) + i.percentW(5)

	bigFont := ENV.Fonts.Large
	txtSpacing := float32(20)

	// Draw fish name
	DrawCenteredText(screen, bigFont, i.item.Name, int(titleX), int(titleY+10), txtColor)

	// draw HP text
	hpFont := ENV.Fonts.Large
	hpString := fmt.Sprintf("%v", i.item.Life)
	DrawCenteredText(screen, hpFont, hpString, int(i.tt.rect.X+i.percentW(86)), int(i.tt.rect.Y+i.percentW(15)), txtColor)

	// draw damage panel
	if i.damageIcon != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i.tt.rect.X+i.percentW(75)), float64(i.tt.rect.Y+i.percentW(25)))
		screen.DrawImage(i.damageIcon, op)
	}
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

func (i *InitialToolTip) GetAlignment() Alignment {
	return i.alignment
}

func (i *InitialToolTip) GetRect() *shapes.Rectangle {
	return &i.tt.rect
}

func (i *InitialToolTip) percentW(percent int) float32 {
	return i.tt.rect.W * (float32(percent) / 100.0)
}

func (i *InitialToolTip) percentH(percent int) float32 {
	return i.tt.rect.H * (float32(percent) / 100.0)
}
