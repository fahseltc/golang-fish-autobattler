package fish

type Type int

const (
	Weapon Type = iota
	SizeBasedWeapon
	AdjacencyBasedWeapon
	Reactive
	VenomousBasedWeapon
	SoloBasedWeapon
)

var itemNames = map[Type]string{
	Weapon:               "weapon",
	SizeBasedWeapon:      "sizeBasedWeapon",
	AdjacencyBasedWeapon: "adjacencyBasedWeapon",
	Reactive:             "reactive",
	VenomousBasedWeapon:  "venomousBasedWeapon",
	SoloBasedWeapon:      "soloBasedWeapon",
}

var itemTypes = map[string]Type{
	"weapon":               Weapon,
	"sizeBasedWeapon":      SizeBasedWeapon,
	"adjacencyBasedWeapon": AdjacencyBasedWeapon,
	"reactive":             Reactive,
	"venomousBasedWeapon":  VenomousBasedWeapon,
	"soloBasedWeapon":      SoloBasedWeapon,
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

func (it Type) ToBehaviorFunc() func(*Fish, *Fish, int, []*Fish) bool {
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
	case SoloBasedWeapon:
		return SoloAttackingBehavior
	default:
		return AttackingBehavior
	}
}
