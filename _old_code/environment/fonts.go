package environment

import (
	"bytes"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

type Fonts struct {
	Small text.Face
	Med   text.Face
	Large text.Face
}

func NewFontsCollection() *Fonts {
	fonts := &Fonts{}
	s, _ := loadFont(12)
	fonts.Small = s
	m, _ := loadFont(20)
	fonts.Med = m
	l, _ := loadFont(30)
	fonts.Large = l
	return fonts
}

func loadFont(size float64) (text.Face, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Error loading font: %w", err)
	}

	return &text.GoTextFace{
		Source: s,
		Size:   size,
	}, nil
}
