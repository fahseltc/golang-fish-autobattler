package ui

import (
	"fishgame/environment"
	"fishgame/util"

	"github.com/hajimehoshi/ebiten/v2"
)

type DlgOptFunc func(*Button)

type Dialog struct {
	Env        *environment.Env
	Rect       Rectangle // rectangle for collision detection
	Title      string
	Content    string
	Buttons    []*Button
	Active     bool
	Selected   int
	Background *ebiten.Image
}

func NewDialog(env *environment.Env, rect Rectangle, title string, content string) *Dialog {
	btn := NewButton(env,
		WithText("OK"),
		WithRect(Rectangle{X: 0, Y: 0, W: 100, H: 50}),
	)
	dlg := &Dialog{
		Env:        env,
		Rect:       Rectangle{X: 0, Y: 0, W: 800, H: 600},
		Title:      title,
		Content:    content,
		Buttons:    []*Button{btn},
		Active:     false,
		Selected:   999,
		Background: util.LoadImage(env, "assets/ui/dialog_background.png"),
	}

	return dlg
}
