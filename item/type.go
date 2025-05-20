package item

type Type int

const (
	Weapon Type = iota
	SizeBasedWeapon
	AdjacencyBasedWeapon
	Reactive
	Venomous
)

var itemNames = map[Type]string{
	Weapon:               "weapon",
	SizeBasedWeapon:      "sizeBasedWeapon",
	AdjacencyBasedWeapon: "adjacencyBasedWeapon",
	Reactive:             "reactive",
	Venomous:             "venomous",
}

var itemTypes = map[string]Type{
	"weapon":               Weapon,
	"sizeBasedWeapon":      SizeBasedWeapon,
	"adjacencyBasedWeapon": AdjacencyBasedWeapon,
	"reactive":             Reactive,
	"venomous":             Venomous,
}

func (it Type) String() string {
	return itemNames[it]
}

func TypeFromString(s string) Type {
	if t, ok := itemTypes[s]; ok {
		return t
	} else {
		return Weapon
	}
}
