package reward

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/player"
)

type Reward struct {
	env      *environment.Env
	Type     Type
	obtained bool
	Item     *item.Item
	Currency int
}

func NewReward(env *environment.Env, newType Type, it *item.Item, curr int) *Reward {
	r := &Reward{
		env:      env,
		Type:     newType,
		obtained: false,
		Item:     it,
		Currency: curr,
	}
	return r
}

func (r *Reward) Obtain(player *player.Player) bool {
	if !r.obtained {
		switch r.Type {
		case Item:
			if r.Item != nil {
				result := player.Items.AddItem(r.Item)
				if result {
					r.obtained = true
					r.env.Logger.Info("player obtained reward item", "item", r.Item.Name)
				}
				return result
			}
			return false

		case Currency:
			if r.Currency != 0 {
				player.Currency += r.Currency
				r.obtained = true
				r.env.Logger.Info("player obtained reward currency", "amount", r.Currency)
				return true
			}
			return false
		}

	}
	return false
}
