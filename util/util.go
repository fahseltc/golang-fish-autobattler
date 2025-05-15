package util

import (
	"bytes"
	"fishgame/environment"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

func LoadImage(Env *environment.Env, filePath string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(filePath)
	if err != nil {
		Env.Logger.Error("Error loading image", "filePath", filePath, "error", err)
		return nil
	}
	return img
}

func LoadFont(size float64) (text.Face, error) {
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
