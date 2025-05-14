package ui

type CollisionType int

const (
	CollisionTopHalf CollisionType = iota
	CollisionBottomHalf
	CollisionNone
)

type Collision struct {
	Type     CollisionType
	Collides bool
}

func (ct CollisionType) ToBool() bool {
	if ct == CollisionBottomHalf || ct == CollisionTopHalf {
		return true
	} else {
		return false
	}
}
