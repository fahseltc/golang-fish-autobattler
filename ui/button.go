package ui

import (
	"fishgame/environment"
	"fishgame/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Button struct {
	env    *environment.Env
	x      float32
	y      float32
	width  float32
	height float32

	text string
	font text.Face

	currentImg *ebiten.Image
	defaultImg *ebiten.Image
	pressedImg *ebiten.Image

	OnClick func()
	ToolTip TooltipInterface
}

func NewButton(env *environment.Env, x, y, w, h float32, txt string, fntSize float64) *Button {
	font, _ := util.LoadFont(fntSize)

	// make button centered on given coordinates
	centeredX := x - 0.5*w
	centeredY := y - 0.5*h

	defaultImg := util.LoadImage(env, "assets/ui/btn/green_button.png")
	defaultImg = util.ScaleImage(defaultImg, w, h)
	pressed := util.LoadImage(env, "assets/ui/btn/green_button_pressed.png")
	pressed = util.ScaleImage(pressed, w, h)

	btn := &Button{
		env:        env,
		x:          centeredX,
		y:          centeredY,
		width:      w,
		height:     h,
		text:       txt,
		font:       font,
		currentImg: defaultImg,
		defaultImg: defaultImg,
		pressedImg: pressed,
	}
	// btn.tt = NewInitialTooltip(btn, int(centeredX), int(centeredY), int(w), int(h), nil)

	return btn
}

func (btn *Button) Draw(screen *ebiten.Image) {
	if btn.ToolTip != nil && btn.MouseCollides() {
		btn.ToolTip.OnHover(screen)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(btn.x), float64(btn.y))
	screen.DrawImage(btn.currentImg, op)

	// draw text centered
	centerX, centerY := btn.GetCenter()
	DrawCenteredText(screen, btn.font, btn.text, centerX, centerY, nil)
}

func (btn *Button) Update() {
	if btn.OnClick != nil && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && btn.MouseCollides() {
		btn.currentImg = btn.pressedImg
	}
	if btn.OnClick != nil && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && btn.MouseCollides() {
		btn.OnClick()
		btn.currentImg = btn.defaultImg
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
