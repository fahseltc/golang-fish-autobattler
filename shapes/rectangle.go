package shapes

type Rectangle struct {
	X float32
	Y float32
	W float32
	H float32
}

// get center coordinates of rectangle
func (r *Rectangle) Center() (float32, float32) {
	return r.X + r.W/2, r.Y + r.H/2
}
