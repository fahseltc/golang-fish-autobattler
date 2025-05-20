package ui

import (
	"fishgame/environment"
	"fishgame/util"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	env    *environment.Env
	x      float32
	y      float32
	width  float32
	height float32

	text  string
	font  text.Face
	color color.Color
	bg    *ebiten.Image
	// todo bg image?
	OnClick func()
	//tt      *Tooltip
}

func NewButton(env *environment.Env, x, y, w, h float32, txt string, clr color.Color, fntSize float64, img *ebiten.Image) *Button {
	font, _ := util.LoadFont(fntSize)

	// make button centered on given coordinates
	centeredX := x - 0.5*w
	centeredY := y - 0.5*h
	var btnImg *ebiten.Image
	if img != nil {
		btnImg = util.ScaleImage(img, w, h)
	}

	btn := &Button{
		env:    env,
		x:      centeredX,
		y:      centeredY,
		width:  w,
		height: h,
		text:   txt,
		font:   font,
		color:  clr,
		bg:     btnImg,
	}
	// btn.tt = NewInitialTooltip(btn, int(centeredX), int(centeredY), int(w), int(h), nil)

	return btn
}

func (btn *Button) Draw(screen *ebiten.Image) {
	if btn.bg == nil {
		vector.DrawFilledRect(screen, btn.x, btn.y, btn.width, btn.height, btn.color, true)
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(btn.x), float64(btn.y))
		screen.DrawImage(btn.bg, op)
	}

	// draw text centered
	centerX, centerY := btn.GetCenter()
	DrawCenteredText(screen, btn.font, btn.text, centerX, centerY)
}

func (btn *Button) Update() {
	if btn.OnClick != nil && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && btn.MouseCollides() {
		btn.OnClick()
	}
}

func (btn *Button) MouseCollides() bool {
	mx, my := ebiten.CursorPosition()
	collides := mx > int(btn.x) && mx < int(btn.x+btn.width) && my > int(btn.y) && my < int(btn.y+btn.height)
	return collides
}

func (btn *Button) GetCenter() (x, y int) {
	centerX := btn.x + btn.width/2
	centerY := btn.y + btn.height/2
	return int(centerX), int(centerY)
}
