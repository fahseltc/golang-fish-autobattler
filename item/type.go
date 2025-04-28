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
	Reactive:  "reactive",
	Misc:      "misc",
}

var itemTypes = map[string]Type{
	"weapon":    Weapon,
	"defensive": Defensive,
	"reactive":  Reactive,
	"misc":      Misc,
}

func (it Type) String() string {
	return itemNames[it]
}

func TypeFromString(s string) Type {
	if t, ok := itemTypes[s]; ok {
		return t
	} else {
		return Misc
	}
}
