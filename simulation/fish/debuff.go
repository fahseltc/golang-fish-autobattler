package fish

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
	fish.env.Logger.Info("debuff", "created", debuff.Type, "remainingDuration", debuff.RemainingDuration, "damage", debuff.damage)
	return debuff
}

func (dbf *Debuff) Update(dt float64) {
	trigger := false
	if dbf.target.IsAlive() && dbf.RemainingDuration > 0 {
		dbf.CurrentTime += dt
		if dbf.CurrentTime >= dbf.TickRate {
			dbf.CurrentTime -= dbf.TickRate
			dbf.RemainingDuration -= dbf.TickRate
			trigger = true
		}
		// if DT is very high, the poison perhaps should have triggered twice? do we want to handle that?
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
