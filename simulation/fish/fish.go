package fish

import (
	"fishgame/shared/environment"
	"time"

	"github.com/google/uuid"
)

type Fish struct {
	env         *environment.Env
	Id          uuid.UUID
	Name        string
	Description string
	Stats       *Stats
	Debuffs     []*Debuff
}

func NewFish(env *environment.Env, name string, desc string, stats *Stats) *Fish {
	f := &Fish{
		env:         env,
		Id:          uuid.New(),
		Name:        name,
		Description: desc,
		Stats:       stats,
	}
	return f
}

func (f *Fish) UpdateDebuffs(dt float64) {
	for index, dbf := range f.Debuffs {
		dbf.Update(dt)
		if dbf.IsDone() {
			f.Debuffs = append(f.Debuffs[:index], f.Debuffs[index+1:]...)
		}
	}
}

// Function for this fish to take a 'param' amount of damage. Returns whether this fish is still alive afterwards or not.
func (f *Fish) TakeDamage(amt int) bool { // TODO if we add reactive fish, we need some more logic here to tell whether dmg was from a debuff or not.
	f.Stats.CurrentLife = f.Stats.CurrentLife - amt
	alive := f.IsAlive()
	if !alive {
		sendFishDiedEvent(f)
	}
	return alive
}
func (f *Fish) AddDebuff(dbf *Debuff) {
	f.Debuffs = append(f.Debuffs, dbf)
}
func (f *Fish) IsAlive() bool {
	alive := f.Stats.CurrentLife > 0
	//fmt.Printf("fish life status: %v result:%v\n", f.Stats.CurrentLife, alive)
	return alive
}
func (f *Fish) IsDead() bool {
	return f.Stats.CurrentLife <= 0
}

// Event senders
func sendFishDiedEvent(deadFish *Fish) {
	deadFish.env.EventBus.Publish(environment.Event{
		Type:      "FishDiedEvent",
		Timestamp: time.Now(),
		Data: environment.FishDiedEvent{
			Id: deadFish.Id,
		},
	})
}
