package item

type ItemAttackedEvent struct {
	Source *Item
	Target *Item
	Damage int
}
