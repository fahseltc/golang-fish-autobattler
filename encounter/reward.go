package encounter

// type RewardType int

// const (
// 	FishReward Type = iota
// 	CurrencyReward
// )

type Reward struct {
	Fish     []string
	Currency int
}

func NewReward(json jsonReward) *Reward {
	return &Reward{
		Fish:     json.Fish,
		Currency: json.Currency,
	}
}
