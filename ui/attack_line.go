package ui

import (
	"fishgame/environment"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type AttackLine struct {
	enabled  bool
	SourceX  int
	SourceY  int
	TargetX  int
	TargetY  int
	Duration float32
}

func NewAttackLine(env *environment.Env, sourceX, sourceY, targetX, targetY int, duration float32) *AttackLine {
	spriteSizePx := float64(env.Get("spriteSizePx").(int))
	spriteScale := env.Get("spriteScale").(float64)
	if sourceX > targetX {
		// If the source is on the right side, adjust the targetX to end at the right edge of the sprite
		targetX += int(spriteSizePx * spriteScale)
	} else {
		// If the source is on the left side, adjust the sourceX to start from the left edge of the sprite
		sourceX += int(spriteSizePx * spriteScale)
	}
	return &AttackLine{
		enabled:  true,
		SourceX:  sourceX,
		SourceY:  sourceY + int(spriteSizePx*spriteScale)/2,
		TargetX:  targetX,
		TargetY:  targetY + int(spriteSizePx*spriteScale)/2,
		Duration: duration,
	}
}

func (a *AttackLine) Update(dt float32) {
	if a.enabled {
		a.Duration -= dt
		if a.Duration < 0 {
			a.Duration = 0
			a.enabled = false
		}
	}
}

func (a *AttackLine) Draw(screen *ebiten.Image) {
	if !a.enabled {
		return
	}

	// Draw the attack line from Source to Target
	lineColor := color.RGBA{200, 0, 0, 255} // this doesnt support alpha
	vector.StrokeLine(screen, float32(a.SourceX), float32(a.SourceY), float32(a.TargetX), float32(a.TargetY), 2, lineColor, true)
	//func DrawFilledCircle(dst *ebiten.Image, cx, cy, r float32, clr color.Color, antialias bool)
	// draw a set number of small circles along the line
	// for i := 0; i < 5; i++ {
	// 	t := float32(i) / 4.0 // 0 to 1
	// 	x := float32(a.SourceX) + t*float32(a.TargetX-a.SourceX)
	// 	y := float32(a.SourceY) + t*float32(a.TargetY-a.SourceY)
	// 	vector.DrawFilledCircle(screen, x, y, 3, lineColor, false)
	// }
}
