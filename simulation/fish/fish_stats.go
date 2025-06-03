package fish

type Stats struct {
	Type Type
	Size Size

	MaxLife     int
	CurrentLife int

	MaxDuration     float64
	CurrentDuration float64

	Damage int

	ActivateFunc func(source *Fish, target *Fish) bool

	// Activate func?
}

type FishSize int

const (
	Small  FishSize = iota
	Medium FishSize = iota
)

func NewWeaponStats(life int, duration int, damage int) *Stats {
	return &Stats{
		Type:            Weapon,
		Size:            SizeMedium,
		MaxLife:         life,
		CurrentLife:     life,
		MaxDuration:     float64(duration),
		CurrentDuration: float64(duration),
		Damage:          damage,

		ActivateFunc: AttackingBehavior,
	}
}

func NewStats(fishType Type, size Size, life int, duration int, damage int) *Stats {
	stats := &Stats{
		Type:            fishType,
		Size:            size,
		MaxLife:         life,
		CurrentLife:     life,
		MaxDuration:     float64(duration),
		CurrentDuration: float64(duration),
		Damage:          damage,
	}
	switch fishType {
	case Weapon:
		stats.ActivateFunc = AttackingBehavior
	case SizeBasedWeapon:
		stats.ActivateFunc = LargerSizeAttackingBehavior
	}

	return stats
}
