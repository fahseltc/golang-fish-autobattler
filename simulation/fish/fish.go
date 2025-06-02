package fish

import "github.com/google/uuid"

type Fish struct {
	Id          uuid.UUID
	Name        string
	Description string
	Stats       *Stats
}

func NewFish(name string, desc string, stats *Stats) *Fish {
	f := &Fish{
		Id:          uuid.New(),
		Name:        name,
		Description: desc,
		Stats:       stats,
	}
	return f
}

func (f *Fish) Update(dt float64, target *Fish) {
	f.Stats.CurrentDuration += dt
	if f.Stats.CurrentDuration >= f.Stats.MaxDuration {
		target.TakeDamage(f.Stats.Damage)
		f.Stats.CurrentDuration -= f.Stats.MaxDuration
	}
}

// Function for this fish to take a 'param' amount of damage. Returns whether this fish is still alive afterwards or not.
func (f *Fish) TakeDamage(amt int) bool {
	f.Stats.CurrentLife -= amt
	return f.IsAlive()
}
func (f *Fish) IsAlive() bool {
	return f.Stats.CurrentLife > 0
}
func (f *Fish) IsDead() bool {
	return f.Stats.CurrentLife <= 0
}
