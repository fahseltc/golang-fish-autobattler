package fish

type Stats struct {
	//Type Type
	//Size string

	MaxLife     int
	CurrentLife int

	MaxDuration     float64
	CurrentDuration float64

	Damage int

	// Activate func?
}

func NewStats(life int, duration int, damage int) *Stats {
	return &Stats{
		MaxLife:         life,
		CurrentLife:     life,
		MaxDuration:     float64(duration),
		CurrentDuration: float64(duration),
		Damage:          damage,
	}
}
