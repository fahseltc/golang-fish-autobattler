package fish

type Stats struct {
	Type Type
	Size Size

	MaxLife     int
	CurrentLife int

	MaxDuration     float64
	CurrentDuration float64
	// SecondaryDuration for extended poison duration?

	Damage int

	ActivateFunc func(source *Fish, target *Fish) bool
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
		CurrentDuration: 0,
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
		CurrentDuration: 0,
		Damage:          damage,
	}
	switch fishType {
	case Weapon:
		stats.ActivateFunc = AttackingBehavior
	case SizeBasedWeapon:
		stats.ActivateFunc = LargerSizeAttackingBehavior
	case VenomousBasedWeapon:
		stats.ActivateFunc = VenomousBehavior
	}

	return stats
}
