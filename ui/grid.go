package ui

type Grid struct {
	rowCount int
	colCount int
}

func NewGrid(rows, cols int) *Grid {
	grid := &Grid{
		rowCount: rows,
		colCount: cols,
	}

	return grid
}

// func (grid *Grid) Draw(screen *ebiten.Image) {
// 	for row := range grid.rowCount {
// 		for col := range grid.colCount {
// 			//vector.DrawFilledRect(screen, ttx, tty, float32(slot.width), float32(slot.width), color.RGBA{128, 128, 128, 255}, true)

// 		}
// 	}

// }
