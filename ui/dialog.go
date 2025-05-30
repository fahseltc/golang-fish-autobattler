package ui

import (
	"fishgame/shapes"
	"fishgame/util"

	"github.com/hajimehoshi/ebiten/v2"
)

type DlgOptFunc func(*Button)

type Dialog struct {
	Rect       shapes.Rectangle // rectangle for collision detection
	Title      string
	Content    string
	Buttons    []*Button
	Active     bool
	Selected   int
	Background *ebiten.Image
}

func NewDialog(rect shapes.Rectangle, title string, content string) *Dialog {
	btn := NewButton(
		WithText("OK"),
		WithRect(shapes.Rectangle{X: 0, Y: 0, W: 100, H: 50}),
	)
	dlg := &Dialog{
		Rect:       shapes.Rectangle{X: 0, Y: 0, W: 800, H: 600},
		Title:      title,
		Content:    content,
		Buttons:    []*Button{btn},
		Active:     false,
		Selected:   999,
		Background: util.LoadImage(ENV, "assets/ui/dialog_background.png"),
	}

	return dlg
}
