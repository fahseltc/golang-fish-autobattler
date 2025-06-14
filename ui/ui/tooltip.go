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
	ChangeAlignment(alignment shapes.Alignment)
	GetAlignment() shapes.Alignment
	GetRect() *shapes.Rectangle
}

type Tooltip struct {
	rect      shapes.Rectangle
	bg        *ebiten.Image
	fonts     *environment.Fonts
	alignment shapes.Alignment
}

type FishToolTip struct {
	tooltip     *Tooltip
	fish        *fish.Fish
	hpIcon      *ebiten.Image
	damagePanel *ebiten.Image
	textArea    *TextArea
}

func NewFishToolTip(env *environment.Env, rect shapes.Rectangle, alignment shapes.Alignment, fish *fish.Fish) *FishToolTip {
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
	durationTxt := fmt.Sprintf("%vs", tt.fish.Stats.MaxDuration)
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

func (tt *FishToolTip) ChangeAlignment(alignment shapes.Alignment) {
	tt.tooltip.alignment = alignment
}
func (tt *FishToolTip) GetAlignment() shapes.Alignment {
	return tt.tooltip.alignment
}
func (tt *FishToolTip) GetRect() *shapes.Rectangle {
	return &tt.tooltip.rect
}
