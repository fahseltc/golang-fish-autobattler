package collection

import (
	"fishgame-sim/environment"
	"fishgame-sim/fish"
	"math/rand/v2"
)

type Collection struct {
	fishSlotMap    map[int]*fish.Fish
	preventChanges bool
}

func NewCollection(env *environment.Env) *Collection {
	coll := &Collection{
		fishSlotMap:    make(map[int]*fish.Fish),
		preventChanges: false,
	}
	coll.fishSlotMap[0] = nil
	coll.fishSlotMap[1] = nil
	coll.fishSlotMap[2] = nil
	coll.fishSlotMap[3] = nil
	coll.fishSlotMap[4] = nil

	env.EventBus.Subscribe("StartSimulationEvent", coll.startSimulationEventHandler)
	env.EventBus.Subscribe("StopSimulationEvent", coll.stopSimulationEventHandler)
	return coll
}

func (coll *Collection) Update(dt float64, enemyColl *Collection) {
	for _, fish := range coll.fishSlotMap {
		if fish != nil && fish.IsAlive() {
			target := enemyColl.GetRandomFish()
			if target != nil {
				fish.Update(dt, target)
			}
		}
	}
}

func (coll *Collection) GetAllFish() []*fish.Fish {
	arr := make([]*fish.Fish, 5)
	for i := 0; i < 5; i++ {
		arr[i] = coll.fishSlotMap[i]
	}
	return arr
}

func (coll *Collection) GetRandomFish() *fish.Fish {
	// Collect non-nil fish pointers
	nonNilFish := []*fish.Fish{}
	for _, f := range coll.fishSlotMap {
		if f != nil {
			nonNilFish = append(nonNilFish, f)
		}
	}
	if len(nonNilFish) == 0 {
		return nil
	}
	return nonNilFish[rand.IntN(len(nonNilFish))]
}

func (coll *Collection) IndexEmpty(index int) bool {
	if index > 4 || index < 0 {
		return false
	}
	return coll.fishSlotMap[index] == nil
}

func (coll *Collection) AddFish(fish *fish.Fish, index int) bool {
	if coll.preventChanges {
		return false
	}
	if !coll.IndexEmpty(index) {
		return false
	}
	coll.fishSlotMap[index] = fish
	return true
}
func (coll *Collection) RemoveFish(id string) bool {
	if coll.preventChanges {
		return false
	}
	for index, f := range coll.fishSlotMap {
		if f != nil && f.Id.String() == id {
			coll.fishSlotMap[index] = nil
			return true
		}
	}
	return false
}
func (coll *Collection) MoveFish(sourceIndex int, targetIndex int) bool {
	if coll.preventChanges {
		return false
	}
	if sourceIndex > 4 ||
		sourceIndex < 0 ||
		targetIndex > 4 ||
		targetIndex < 0 ||
		sourceIndex == targetIndex {
		return false
	}
	sourceFish := coll.fishSlotMap[sourceIndex]
	if sourceFish == nil {
		return false
	} else {
		targetFish := coll.fishSlotMap[targetIndex]
		if targetFish != nil {
			// swap them
			coll.fishSlotMap[targetIndex] = sourceFish
			coll.fishSlotMap[sourceIndex] = targetFish
		} else {
			// targer is empty, just move it
			coll.fishSlotMap[targetIndex] = sourceFish // move fish
			coll.fishSlotMap[sourceIndex] = nil        // clear its old slot
		}

		return true
	}
}
func (coll *Collection) DisableChanges() {
	coll.preventChanges = true
}
func (coll *Collection) EnableChanges() {
	coll.preventChanges = false
}
func (coll *Collection) IsChangeable() bool {
	return !coll.preventChanges
}
func (coll *Collection) AllFishDead() bool {
	for _, fish := range coll.fishSlotMap {
		if fish != nil && fish.IsAlive() {
			return false
		}
	}
	return true
}

// Event handlers
func (coll *Collection) startSimulationEventHandler(event environment.Event) {
	coll.DisableChanges()
}

func (coll *Collection) stopSimulationEventHandler(event environment.Event) {
	coll.EnableChanges()
}
