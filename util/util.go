package util

import (
	"fishgame/environment"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func LoadImage(Env environment.Env, filePath string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(filePath, 0)
	if err != nil {
		Env.Logger.Error("Error loading image", "filePath", filePath, "error", err)
		return nil
	}
	return img
}
