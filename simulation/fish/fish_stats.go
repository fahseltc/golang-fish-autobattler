package fish

import "github.com/google/uuid"

type Stats struct {
	uuid uuid.UUID
	Type Type
	Size Size

	MaxLife     int
	CurrentLife int

	MaxDuration     float64
	CurrentDuration float64
	// SecondaryDuration for extended poison duration?

	Damage int

	ActivateFunc func(source *Fish, target *Fish, index int, ownCollection []*Fish) bool
}

type FishSize int

const (
	Small  FishSize = iota
	Medium FishSize = iota
)

func NewWeaponStats(life int, duration int, damage int) Stats {
	return Stats{
		uuid:            uuid.New(),
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

func NewStats(fishType Type, size Size, life int, duration float64, damage int) Stats {
	stats := Stats{
		uuid:            uuid.New(),
		Type:            fishType,
		Size:            size,
		MaxLife:         life,
		CurrentLife:     life,
		MaxDuration:     float64(duration),
		CurrentDuration: 0,
		Damage:          damage,
	}

	stats.ActivateFunc = fishType.ToBehaviorFunc()

	return stats
}
