package ui

import (
	"fishgame/ui/shapes"
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TextArea struct {
	rect  shapes.Rectangle
	text  string
	lines []string

	overflows bool
}

func NewTextArea(text string, rect shapes.Rectangle) *TextArea {
	ta := &TextArea{
		rect:      rect,
		text:      text,
		overflows: false,
	}
	ta.splitTextOntoLines()
	return ta
}

func (ta *TextArea) splitTextOntoLines() {
	font := ENV.Fonts.XSmall
	// Split the text into words
	words := strings.Split(ta.text, " ") // if we use commas or something else, this will have bugs
	var currentLine string
	var totalHeight float64

	for _, word := range words {
		testLine := currentLine
		if testLine != "" {
			testLine += " "
		}
		testLine += word

		// Measure the width of the testLine
		tw, th := text.Measure(testLine, font, 2)
		totalHeight += th

		if tw > float64(ta.rect.W) && currentLine != "" {
			ta.lines = append(ta.lines, currentLine)
			currentLine = word
		} else {
			currentLine = testLine
		}
		fmt.Printf("totalHeight:%v", totalHeight)

	}
	if currentLine != "" {
		ta.lines = append(ta.lines, currentLine)
	}
}

func (ta *TextArea) Draw(screen *ebiten.Image, ttrect shapes.Rectangle) {
	ta.rect.X = ttrect.X + 78
	ta.rect.Y = ttrect.Y + 30
	ebitenutil.DrawRect(screen, float64(ta.rect.X), float64(ta.rect.Y), float64(ta.rect.W), float64(ta.rect.H), color.RGBA{200, 200, 155, 255})
	font := ENV.Fonts.XSmall
	_, th := text.Measure(ta.text, font, 2)
	th += 2
	y := float64(ta.rect.Y) + th

	for _, line := range ta.lines {
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(float64(ta.rect.X+4), y)
		opts.ColorScale.ScaleWithColor(color.RGBA{R: 0, G: 0, B: 0, A: 255})
		text.Draw(screen, line, font, opts)
		y += th
	}
}
