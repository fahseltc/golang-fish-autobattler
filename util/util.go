package util

import (
	"fishgame/environment"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadImage(Env *environment.Env, filePath string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(filePath)
	if err != nil {
		Env.Logger.Error("Error loading image", "filePath", filePath, "error", err)
		return nil
	}
	return img
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
