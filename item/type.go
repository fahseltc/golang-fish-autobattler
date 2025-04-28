package item

type Type int

const (
	Weapon Type = iota
	Defensive
	Reactive
	Misc
)

var itemNames = map[Type]string{
	Weapon:    "weapon",
	Defensive: "defensive",
	Misc:      "misc",
	Reactive:  "reactive",
}

func (it Type) String() string {
	return itemNames[it]
}
