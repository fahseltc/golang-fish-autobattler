package inventory

import "fishgame-sim/fish"

type Inventory struct {
	fish map[string]*fish.Fish
}

func NewInventory() *Inventory {
	return &Inventory{
		fish: make(map[string]*fish.Fish),
	}
}

func (inv *Inventory) Add(fish *fish.Fish) {
	inv.fish[fish.Id.String()] = fish
}

func (inv *Inventory) Get(id string) *fish.Fish {
	fish := inv.fish[id]
	if fish != nil {
		inv.fish[id] = nil
		return fish
	}
	return nil
}

func (inv *Inventory) GetAll() []*fish.Fish {
	fishes := make([]*fish.Fish, 0, len(inv.fish))
	for _, f := range inv.fish {
		fishes = append(fishes, f)
	}
	return fishes
}

func (inv *Inventory) GetCount() int {
	return len(inv.fish)
}
