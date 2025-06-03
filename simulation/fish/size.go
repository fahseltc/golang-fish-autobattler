package fish

type Size int

const (
	SizeTiny Size = iota
	SizeSmall
	SizeMedium
	SizeLarge
	SizeHuge
	SizeUnknown
)

var sizeNames = map[Size]string{
	SizeTiny:    "tiny",
	SizeSmall:   "small",
	SizeMedium:  "medium",
	SizeLarge:   "large",
	SizeHuge:    "huge",
	SizeUnknown: "unknown",
}

var sizeTypes = map[string]Size{
	"tiny":    SizeTiny,
	"small":   SizeSmall,
	"medium":  SizeMedium,
	"large":   SizeLarge,
	"huge":    SizeHuge,
	"unknown": SizeUnknown,
}

func (s Size) String() string {
	return sizeNames[s]
}

func SizeFromString(s string) Size {
	if sz, ok := sizeTypes[s]; ok {
		return sz
	} else {
		return SizeUnknown
	}
}
