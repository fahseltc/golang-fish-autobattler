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

func ScaleImage(input *ebiten.Image, newWidth float32, newHeight float32) *ebiten.Image {
	imgW, imgH := input.Size()
	//example: input image is 300,200 and we want to draw it at 200,100 size
	wScale := newWidth / float32(imgW)
	hScale := newHeight / float32(imgH)

	scaledImg := ebiten.NewImage(int(newWidth), int(newHeight))
	op := &ebiten.DrawImageOptions{} // Draw original onto the new image with scaling
	op.GeoM.Scale(float64(wScale), float64(hScale))
	scaledImg.DrawImage(input, op)
	return scaledImg
}
