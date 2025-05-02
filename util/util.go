package util

import (
	"bytes"
	"fishgame/environment"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

func LoadImage(Env environment.Env, filePath string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(filePath)
	if err != nil {
		Env.Logger.Error("Error loading image", "filePath", filePath, "error", err)
		return nil
	}
	return img
}

func DrawLifeBar(screen *ebiten.Image, healthRatio float64, x, y float64) {
	healthLength := float64(64 * healthRatio)
	ebitenutil.DrawRect(screen, x, y+64, healthLength, 4, color.White)
}

func DrawProgressBar(screen *ebiten.Image, progressRatio float64, x, y float64) {
	progressLength := float64(64 * progressRatio)
	ebitenutil.DrawRect(screen, x, y+69, progressLength, 4, color.RGBA{0, 255, 0, 255})
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
