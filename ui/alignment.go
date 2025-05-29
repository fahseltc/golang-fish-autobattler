package ui

type Alignment int

const (
	CenterAlignment Alignment = iota
	TopAlignment
	BottomAlignment
	LeftAlignment
	RightAlignment
)

func (alignment Alignment) Align(source Rectangle, toBeAligned *Rectangle) {
	padding := float32(10)
	switch alignment {
	case CenterAlignment:
		toBeAligned.X = source.X + (source.W-toBeAligned.W)/2
		toBeAligned.Y = source.Y + (source.H-toBeAligned.H)/2
	case TopAlignment:
		toBeAligned.X = source.X + (source.W-toBeAligned.W)/2
		toBeAligned.Y = source.Y
	case BottomAlignment:
		// toBeAligned.X = source.X + (source.W-toBeAligned.W)/2
		// toBeAligned.Y = source.Y + source.H - toBeAligned.H
		toBeAligned.X = source.X - toBeAligned.W/2
		toBeAligned.Y = source.Y + source.H/2 + padding
	case LeftAlignment:
		toBeAligned.X = source.X
		toBeAligned.Y = source.Y + (source.H-toBeAligned.H)/2
	case RightAlignment:
		toBeAligned.X = source.X + source.W - toBeAligned.W
		toBeAligned.Y = source.Y + (source.H-toBeAligned.H)/2
	}
}
