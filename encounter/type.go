package encounter

type Type int

const (
	EncounterTypeInitial Type = iota
	EncounterTypeShop
	EncounterTypeChoice
	EncounterTypeBattle
)
