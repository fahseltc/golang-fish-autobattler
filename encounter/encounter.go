package encounter

import (
	"fishgame/simulation/collection"
	"fishgame/simulation/fish"
	"log"
)

type Encounter struct {
	Title   string
	Type    Type
	Fish    *collection.Collection
	Rewards []*Reward
}

func NewEncounter(encData jsonEncounter, registry *fish.FishStatsRegistry) EncounterInterface {
	enc := Encounter{
		Title:   encData.Title,
		Type:    TypeFromString(encData.Type),
		Fish:    collection.NewCollection(ENV),
		Rewards: make([]*Reward, 0),
	}

	for _, enemyJson := range encData.Enemies {
		f, err := registry.GetFish(enemyJson.Name) // todo: add index
		if err != nil {
			log.Fatalf("could not load fish with name:%v", enemyJson.Name)
		}
		enc.Fish.AddFish(f, enemyJson.SlotIndex)
	}

	for _, rewardData := range encData.Rewards {
		r := NewReward(rewardData)
		enc.Rewards = append(enc.Rewards, r)
	}

	return enc
}
func (enc Encounter) GetCollection() *collection.Collection {
	return enc.Fish
}
func (enc Encounter) GetRewards() []*Reward {
	return enc.Rewards
}
func (enc Encounter) GetType() Type {
	return enc.Type
}
func (enc Encounter) IsDone() bool {
	return false
}
func (enc Encounter) IsStarted() bool {
	return false
}
func (enc Encounter) IsGameOver() bool {
	return false
}
