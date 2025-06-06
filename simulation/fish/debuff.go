package fish

import (
	"fmt"
)

// this will handle DOT effects for now, but other things are possible.

type DebuffType int

const (
	DebuffTypeNone DebuffType = iota
	DebuffTypeVenom
	DebuffTypeSlow
	//... more
)

type DebuffInterface interface {
	Update(float64)
	IsDone() bool
}

type Debuff struct {
	MaxDuration       float64
	RemainingDuration float64
	CurrentTime       float64
	TickRate          float64
	damage            int
	Type              DebuffType
	target            *Fish
}

func NewItemDebuff(fish *Fish, dbt DebuffType, dur float64, tickRate float64, dam int) *Debuff {
	debuff := &Debuff{
		RemainingDuration: dur,
		CurrentTime:       0,
		TickRate:          tickRate,
		Type:              dbt,
		target:            fish,
		damage:            dam,
	}
	fmt.Printf("create debuff, remainingDur: %v, tickRate: %v\n", debuff.RemainingDuration, debuff.TickRate)
	return debuff
}

// example 20 remdur with tick rate 2
// many frames occur where we just need to subtract dt from remainingduration
// at some point we have removed '2' worth of dt time from remainingdur, and need to trigger

func (dbf *Debuff) Update(dt float64) {
	trigger := false
	if dbf.target.IsAlive() && dbf.RemainingDuration > 0 {
		dbf.CurrentTime += dt
		if dbf.CurrentTime >= dbf.TickRate {
			dbf.CurrentTime -= dbf.TickRate
			dbf.RemainingDuration -= dbf.TickRate
			trigger = true
		}
	}

	if trigger {
		switch dbf.Type {
		case DebuffTypeVenom:
			dbf.target.TakeDamage(dbf.damage)
			return
		case DebuffTypeSlow:
			return
		default:
			return
		}
	}
}

func (dbf *Debuff) IsDone() bool {
	return dbf.RemainingDuration <= 0
}
