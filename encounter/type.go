package encounter

type Type int

const (
	Initial Type = iota
	Shop
	Choice
	Battle
)
