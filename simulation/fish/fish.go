package fish

import (
	"fmt"

	"github.com/google/uuid"
)

type Fish struct {
	Id          uuid.UUID
	Name        string
	Description string
	Stats       *Stats
	//activate
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
		targetAlive := f.Stats.ActivateFunc(f, target)
		if !targetAlive {
			fmt.Printf("target fish died")
		}
		f.Stats.CurrentDuration -= f.Stats.MaxDuration
	}
}

// Function for this fish to take a 'param' amount of damage. Returns whether this fish is still alive afterwards or not.
func (f *Fish) TakeDamage(amt int) bool { // TODO if we add reactive fish, we need some more logic here to tell whether dmg was from a debuff or not.
	f.Stats.CurrentLife = f.Stats.CurrentLife - amt
	return f.IsAlive()
}
func (f *Fish) IsAlive() bool {
	alive := f.Stats.CurrentLife > 0
	//fmt.Printf("fish life status: %v result:%v\n", f.Stats.CurrentLife, alive)
	return alive
}
func (f *Fish) IsDead() bool {
	return f.Stats.CurrentLife <= 0
}
