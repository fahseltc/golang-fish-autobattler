package fish

type Type int

const (
	Weapon Type = iota
	SizeBasedWeapon
	AdjacencyBasedWeapon
	Reactive
	VenomousBasedWeapon
)

var itemNames = map[Type]string{
	Weapon:               "weapon",
	SizeBasedWeapon:      "sizeBasedWeapon",
	AdjacencyBasedWeapon: "adjacencyBasedWeapon",
	Reactive:             "reactive",
	VenomousBasedWeapon:  "venomousBasedWeapon",
}

var itemTypes = map[string]Type{
	"weapon":               Weapon,
	"sizeBasedWeapon":      SizeBasedWeapon,
	"adjacencyBasedWeapon": AdjacencyBasedWeapon,
	"reactive":             Reactive,
	"venomousBasedWeapon":  VenomousBasedWeapon,
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

func (it Type) ToBehaviorFunc() func(*Fish, *Fish) bool {
	switch it {
	case Weapon:
		return AttackingBehavior
	case SizeBasedWeapon:
		return LargerSizeAttackingBehavior
	// case AdjacencyBasedWeapon:
	// 	return AdjacencyBasedWeapon
	// Reactive
	case VenomousBasedWeapon:
		return VenomousBehavior
	default:
		return AttackingBehavior
	}
}
