package btn

import (
	"fishgame/shapes"
	"fishgame/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type BtnOptFunc func(*Button)

type Button struct {
	rect shapes.Rectangle // rectangle for collision detection

	text string
	font text.Face

	currentImg *ebiten.Image
	defaultImg *ebiten.Image
	pressedImg *ebiten.Image

	OnClick func()
	ToolTip TooltipInterface
}

//
// NewButton creates a new Button with the given environment and options.
//

func NewButton(opts ...BtnOptFunc) *Button {
	btn := defaultBtnOpts()
	for _, opt := range opts {
		opt(&btn)
	}
	return &btn
}

func defaultBtnOpts() Button {
	defaultWidth := float32(250.0)
	defaultHeight := float32(100.0)
	font := ENV.Fonts.Med
	defaultImg := util.LoadImage(ENV, "assets/ui/btn/green_button.png")
	defaultImg = util.ScaleImage(defaultImg, defaultWidth, defaultHeight)
	pressed := util.LoadImage(ENV, "assets/ui/btn/green_button_pressed.png")
	pressed = util.ScaleImage(pressed, defaultWidth, defaultHeight)
	return Button{
		rect: shapes.Rectangle{
			X: 0,
			Y: 0,
			W: 250,
			H: 100,
		},
		font:       font,
		text:       "DefaultText",
		currentImg: defaultImg,
		defaultImg: defaultImg,
		pressedImg: pressed,
	}
}
func WithText(txt string) BtnOptFunc {
	return func(btn *Button) {
		btn.text = txt
	}
}
func WithRect(rect shapes.Rectangle) BtnOptFunc {
	return func(btn *Button) {
		btn.rect = rect
		defaultImg := util.LoadImage(ENV, "assets/ui/btn/green_button.png")
		defaultImg = util.ScaleImage(defaultImg, rect.W, rect.H)
		pressed := util.LoadImage(ENV, "assets/ui/btn/green_button_pressed.png")
		pressed = util.ScaleImage(pressed, rect.W, rect.H)

		btn.currentImg = defaultImg
		btn.defaultImg = defaultImg
		btn.pressedImg = pressed
	}
}
func WithClickFunc(f func()) BtnOptFunc {
	return func(btn *Button) {
		btn.OnClick = f
	}
}
func WithToolTip(tt TooltipInterface) BtnOptFunc {
	return func(btn *Button) {
		btn.ToolTip = tt
		btn.ToolTip.GetAlignment().Align(btn.rect, tt.GetRect())
	}
}
func WithCenteredPos() BtnOptFunc {
	return func(btn *Button) {
		centeredX := btn.rect.X - 0.5*btn.rect.W
		centeredY := btn.rect.Y - 0.5*btn.rect.H
		btn.rect.X = centeredX
		btn.rect.Y = centeredY
	}
}

//
// Class Functions
//

func (btn *Button) Draw(screen *ebiten.Image) {
	if btn.ToolTip != nil && btn.MouseCollides() {
		btn.ToolTip.OnHover(screen)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(btn.rect.X), float64(btn.rect.Y))
	screen.DrawImage(btn.currentImg, op)

	// draw text centered
	centerX, centerY := btn.GetCenter()
	util.DrawCenteredText(screen, btn.font, btn.text, centerX, centerY, nil)
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
	collides := mx > int(btn.rect.X) && mx < int(btn.rect.X+btn.rect.W) && my > int(btn.rect.Y) && my < int(btn.rect.Y+btn.rect.H)
	return collides
}

func (btn *Button) GetCenter() (x, y int) {
	centerX := btn.rect.X + btn.rect.W/2
	centerY := btn.rect.Y + btn.rect.H/2
	return int(centerX), int(centerY)
}
