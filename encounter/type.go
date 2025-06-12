package encounter

type Type int

const (
	EncounterTypeInitial Type = iota
	EncounterTypeShop
	EncounterTypeChoice
	EncounterTypeBattle
	EncounterTypeUnknown
)

var encNames = map[Type]string{
	EncounterTypeInitial: "initial",
	EncounterTypeShop:    "shop",
	EncounterTypeChoice:  "choice",
	EncounterTypeBattle:  "battle",
	EncounterTypeUnknown: "unknown",
}

var encTypes = map[string]Type{
	"initial": EncounterTypeInitial,
	"shop":    EncounterTypeShop,
	"choice":  EncounterTypeChoice,
	"battle":  EncounterTypeBattle,
}

func (it Type) String() string {
	return encNames[it]
}

func TypeFromString(s string) Type {
	if t, ok := encTypes[s]; ok {
		return t
	} else {
		return EncounterTypeUnknown
	}
}
