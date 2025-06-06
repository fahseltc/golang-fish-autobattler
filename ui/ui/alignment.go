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
	padding := float32(10)
	switch alignment {
	case CenterAlignment:
		toBeAligned.X = source.X + (source.W-toBeAligned.W)/2
		toBeAligned.Y = source.Y + (source.H-toBeAligned.H)/2
	case TopAlignment:
		toBeAligned.X = source.X + (source.W-toBeAligned.W)/2
		toBeAligned.Y = source.Y
	case BottomAlignment: // works with initialChoice TT
		toBeAligned.X = source.X - toBeAligned.W/2
		toBeAligned.Y = source.Y + source.H/2 + padding
	case LeftAlignment: // works with in-game TT
		toBeAligned.X = source.X - toBeAligned.W - padding    // left TT is not centered
		toBeAligned.Y = source.Y + (source.H-toBeAligned.H)/2 // left TT is not centered
	case RightAlignment: // SOON works with in-game TT
		toBeAligned.X = source.X + source.W + padding         // right TT is not centered
		toBeAligned.Y = source.Y + (source.H-toBeAligned.H)/2 // right TT is not centered
	}
}
