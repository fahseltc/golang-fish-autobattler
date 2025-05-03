package item

import (
	"image"
	"image/color"
)

type Draggable struct {
	item   *Item
	hitbox *image.Alpha
	x      int
	y      int
}

func (d *Draggable) In(x, y int) bool {
	return d.hitbox.At(x-d.x, y-d.y).(color.Alpha).A > 0
}
