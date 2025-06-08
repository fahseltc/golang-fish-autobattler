package ui

import "fishgame/ui/shapes"

type Alignment int

const (
	CenterAlignment Alignment = iota
	TopAlignment
	BottomAlignment
	LeftAlignment
	RightAlignment
)

func (alignment Alignment) Align(source shapes.Rectangle, toBeAligned *shapes.Rectangle) {
	padding := float32(ENV.Config.Get("tooltip.pad").(int))
	switch alignment {
	case CenterAlignment:
		toBeAligned.X = source.X + (source.W-toBeAligned.W)/2
		toBeAligned.Y = source.Y + (source.H-toBeAligned.H)/2
	case TopAlignment:
		toBeAligned.X = source.X + (source.W-toBeAligned.W)/2
		toBeAligned.Y = source.Y
	case BottomAlignment: // WORKS
		toBeAligned.X = source.X + (source.W-toBeAligned.W)/2
		toBeAligned.Y = source.Y + source.H + padding
	case LeftAlignment: // WORKS
		toBeAligned.X = source.X - toBeAligned.W - padding
		toBeAligned.Y = source.Y + (source.H-toBeAligned.H)/2
	case RightAlignment:
		toBeAligned.X = source.X + source.W + padding         // right TT is not centered
		toBeAligned.Y = source.Y + (source.H-toBeAligned.H)/2 // right TT is not centered
	}
}
