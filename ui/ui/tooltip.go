package ui

import (
	"fishgame/shared/environment"
	"fishgame/simulation/fish"
	"fishgame/ui/shapes"
	"fishgame/ui/util"
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type TooltipInterface interface {
	OnHover(screen *ebiten.Image)
	ReAlign(sprite *Sprite)
	ChangeAlignment(alignment Alignment)
	//GetAlignment() Alignment
	//GetRect() *shapes.Rectangle
}

type Tooltip struct {
	rect      shapes.Rectangle
	bg        *ebiten.Image
	fonts     *environment.Fonts
	alignment Alignment
}

type FishToolTip struct {
	tooltip     *Tooltip
	fish        *fish.Fish
	hpIcon      *ebiten.Image
	damagePanel *ebiten.Image
	textArea    *TextArea
}

func NewFishToolTip(env *environment.Env, rect shapes.Rectangle, alignment Alignment, fish *fish.Fish) *FishToolTip {
	ttW := float32(env.Config.Get("tooltip.fish.w").(int))
	ttH := float32(env.Config.Get("tooltip.fish.h").(int))
	ttRect := &shapes.Rectangle{X: rect.X, Y: rect.Y, W: ttW, H: ttH}
	scaledBg := util.ScaleImage(util.LoadImage("ui/tooltip/grey_panel.png"), ttW, ttH)
	//hpIcon := util.ScaleImage(util.LoadImage("ui/tooltip/shield.png"), 50, 50)
	alignment.Align(rect, ttRect)
	tt := &FishToolTip{
		tooltip: &Tooltip{
			rect:      *ttRect,
			bg:        scaledBg,
			fonts:     env.Fonts,
			alignment: alignment,
		},
		fish:        fish,
		hpIcon:      util.LoadImage("ui/tooltip/shield.png"),
		damagePanel: util.LoadImage("ui/tooltip/damage_panel_plain.png"),
		textArea:    NewTextArea(fish.Description, shapes.Rectangle{X: ttRect.X + 78, Y: ttRect.Y + 30, W: 110, H: 110}),
	}
	return tt
}

func (tt *FishToolTip) OnHover(screen *ebiten.Image) {
	font := tt.tooltip.fonts.XSmall
	txtColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}

	// draw BG image
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(tt.tooltip.rect.X), float64(tt.tooltip.rect.Y))
	screen.DrawImage(tt.tooltip.bg, opts)

	// Draw fish name
	nameX := int(tt.tooltip.rect.X + 130)
	nameY := int(tt.tooltip.rect.Y + 18)
	util.DrawCenteredText(screen, ENV.Fonts.Small, strings.Title(tt.fish.Name), nameX, nameY, txtColor)

	// draw HP icon
	hpX := int(tt.tooltip.rect.X + 14)
	hpY := int(tt.tooltip.rect.Y + 8)
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(hpX), float64(hpY))
	screen.DrawImage(tt.hpIcon, opts)
	// Draw HP label
	util.DrawCenteredText(screen, font, "HP", hpX+26, hpY+12, txtColor)
	// Draw current HP text
	curHpStr := fmt.Sprintf("%d", tt.fish.Stats.CurrentLife)
	util.DrawCenteredText(screen, font, curHpStr, hpX+26, hpY+23, txtColor)
	// Draw max HP text
	maxHpStr := fmt.Sprintf("%d", tt.fish.Stats.MaxLife)
	util.DrawCenteredText(screen, font, maxHpStr, hpX+26, hpY+36, txtColor)

	// Draw damage panel BG
	dmgPanelX := int(tt.tooltip.rect.X + 8)
	dmgPanelY := int(tt.tooltip.rect.Y + 65)
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(1, 0.62)
	opts.GeoM.Translate(float64(dmgPanelX), float64(dmgPanelY))

	screen.DrawImage(tt.damagePanel, opts)

	// Draw Damage Label
	util.DrawCenteredText(screen, font, "DMG", dmgPanelX+32, dmgPanelY+9, txtColor)

	// Draw Damage Label
	dmgTxt := fmt.Sprintf("%d", tt.fish.Stats.Damage)
	util.DrawCenteredText(screen, font, dmgTxt, dmgPanelX+32, dmgPanelY+19, txtColor)

	// Draw duration label
	util.DrawCenteredText(screen, font, "DUR", dmgPanelX+32, dmgPanelY+34, txtColor)
	// Draw duration text
	durationTxt := fmt.Sprintf("%vsec", tt.fish.Stats.MaxDuration)
	util.DrawCenteredText(screen, font, durationTxt, dmgPanelX+32, dmgPanelY+44, txtColor)

	// Draw DPS label
	util.DrawCenteredText(screen, font, "DPS", dmgPanelX+32, dmgPanelY+58, txtColor)
	// Draw dps text
	dps := float32(tt.fish.Stats.Damage) / float32(tt.fish.Stats.MaxDuration)
	dpsTxt := fmt.Sprintf("%.2f", dps)
	util.DrawCenteredText(screen, font, dpsTxt, dmgPanelX+32, dmgPanelY+68, txtColor)

	// draw textArea
	if tt.textArea != nil {
		tt.textArea.Draw(screen, tt.tooltip.rect)
	}
}

func (tt *FishToolTip) ReAlign(sprite *Sprite) {
	tt.tooltip.alignment.Align(sprite.Rect, &tt.tooltip.rect)
}

func (tt *FishToolTip) ChangeAlignment(alignment Alignment) {
	tt.tooltip.alignment = alignment
}

// type InitialToolTip struct {
// 	env        *environment.Env
// 	tt         *Tooltip
// 	font       text.Face
// 	fish       *fish.Fish
// 	alignment  Alignment
// 	hpIcon     *ebiten.Image
// 	damageIcon *ebiten.Image
// }

// func NewInitialToolTip(env *environment.Env, x, y, w, h float32, fish *fish.Fish, alignment Alignment) *InitialToolTip {
// 	font := env.Fonts.Med

// 	centeredX := x - 0.5*w
// 	centeredY := y + 0.3*h

// 	tt := &InitialToolTip{
// 		tt: &Tooltip{
// 			rect: shapes.Rectangle{
// 				X: centeredX,
// 				Y: centeredY,
// 				W: w,
// 				H: h,
// 			},
// 			bg: util.LoadImage("ui/tooltip/grey_panel.png"),
// 		},
// 		hpIcon:     util.LoadImage("ui/tooltip/shield.png"),
// 		damageIcon: util.LoadImage("ui/tooltip/damage_panel.png"),
// 		alignment:  alignment,
// 		font:       font,
// 		item:       item,
// 	}
// 	return tt
// }

// func (i *InitialToolTip) OnHover(screen *ebiten.Image) {
// 	txtColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}
// 	i.tt.Draw(screen)

// 	// draw fish sprite
// 	if i.item.Sprite != nil {
// 		op := &ebiten.DrawImageOptions{}
// 		// op.GeoM.Scale(float64(i.tt.rect.W)/float64(i.item.Sprite.Bounds().Dx()), float64(i.tt.rect.H)/float64(i.item.Sprite.Bounds().Dy()))
// 		op.GeoM.Translate(float64(i.tt.rect.X+10), float64(i.tt.rect.Y))
// 		screen.DrawImage(i.item.Sprite, op)
// 	}

// 	// draw hp icon
// 	if i.hpIcon != nil {
// 		op := &ebiten.DrawImageOptions{}
// 		op.GeoM.Translate(float64(i.tt.rect.X+i.percentW(75)), float64(i.tt.rect.Y+i.percentW(3)))
// 		screen.DrawImage(i.hpIcon, op)
// 	}

// 	titleX := float32(i.tt.rect.X) + i.percentH(85)
// 	titleY := float32(i.tt.rect.Y) + i.percentW(5)

// 	bigFont := ENV.Fonts.Large
// 	txtSpacing := float32(20)

// 	// Draw fish name
// 	util.DrawCenteredText(screen, bigFont, i.item.Name, int(titleX), int(titleY+10), txtColor)

// 	// draw HP text
// 	hpFont := ENV.Fonts.Large
// 	hpString := fmt.Sprintf("%v", i.item.Life)
// 	util.DrawCenteredText(screen, hpFont, hpString, int(i.tt.rect.X+i.percentW(86)), int(i.tt.rect.Y+i.percentW(15)), txtColor)

// 	// draw damage panel
// 	if i.damageIcon != nil {
// 		op := &ebiten.DrawImageOptions{}
// 		op.GeoM.Translate(float64(i.tt.rect.X+i.percentW(75)), float64(i.tt.rect.Y+i.percentW(25)))
// 		screen.DrawImage(i.damageIcon, op)
// 	}
// 	// DPS
// 	dpsString := fmt.Sprintf("DPS: %.2f", i.item.Dps())
// 	util.DrawCenteredText(screen, i.font, dpsString, int(titleX), int(titleY+(txtSpacing*4)), txtColor)
// 	// Damage
// 	dmgString := fmt.Sprintf("Damage: %v", i.item.Damage)
// 	util.DrawCenteredText(screen, i.font, dmgString, int(titleX), int(titleY+(txtSpacing*5)), txtColor)
// 	// Duration
// 	durationString := fmt.Sprintf("Duration: %v", i.item.Duration)
// 	util.DrawCenteredText(screen, i.font, durationString, int(titleX), int(titleY+(txtSpacing*6)), txtColor)
// 	// Type
// 	typeString := fmt.Sprintf("%v", strings.Title(i.item.Type.String()))
// 	util.DrawCenteredText(screen, i.font, typeString, int(titleX), int(titleY+(txtSpacing*7)), txtColor)
// 	// Description
// 	util.DrawCenteredText(screen, i.font, i.item.Description, int(titleX), int(titleY+(txtSpacing*8)), txtColor)
// }

// func (i *InitialToolTip) GetAlignment() Alignment {
// 	return i.alignment
// }

// func (i *InitialToolTip) GetRect() *shapes.Rectangle {
// 	return &i.tt.rect
// }

// func (i *InitialToolTip) percentW(percent int) float32 {
// 	return i.tt.rect.W * (float32(percent) / 100.0)
// }

// func (i *InitialToolTip) percentH(percent int) float32 {
// 	return i.tt.rect.H * (float32(percent) / 100.0)
// }
