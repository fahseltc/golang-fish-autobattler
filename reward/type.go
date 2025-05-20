package reward

type Type int

const (
	Item Type = iota
	Currency
)

var itemNames = map[Type]string{
	Item:     "item",
	Currency: "currency",
}

var itemTypes = map[string]Type{
	"item":     Item,
	"currency": Currency,
}

func (it Type) String() string {
	return itemNames[it]
}

func TypeFromString(s string) Type {
	if t, ok := itemTypes[s]; ok {
		return t
	} else {
		return Currency
	}
}
