package item

type Properties struct {
	Name        string
	Type        Type
	Description string
	Size        Size
	Sprite      string

	Life        int
	CurrentLife int

	Alive bool

	Duration        float64
	CurrentDuration float64

	Damage int
}
