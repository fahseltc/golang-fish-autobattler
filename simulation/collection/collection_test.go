package collection

import (
	"fishgame/shared/environment"
	"fishgame/simulation/fish"
	"testing"
)

var ENV *environment.Env

func setupCollection() *Collection {
	ENV = environment.NewEnv(nil, nil)
	return NewCollection(ENV)
}

// [UnitOfWork_StateUnderTest_ExpectedBehaviour]
// https://osherove.com/blog/2005/4/3/naming-standards-for-unit-tests.html
func Test_NewCollection_Constructor_InitializedNil(t *testing.T) {
	coll := setupCollection()
	for key, val := range coll.fishSlotMap {
		if val != nil {
			t.Errorf("NewCollection initialized non-nil - key: %v, val:%v", key, val)
		}
	}
}

func Test_GetRandomFish_WithOneFish_IsReturned(t *testing.T) {
	coll := setupCollection()
	coll.AddFish(fish.NewFish(ENV, "test", "test-fish", nil), 0)
	if coll.GetRandomFish() == nil {
		t.Error("GetRandomFish in collection with 1 fish returned nil!")
	}
}
func Test_GetRandomFish_WithNoFish_ReturnsNil(t *testing.T) {
	coll := setupCollection()
	if coll.GetRandomFish() != nil {
		t.Error("GetRandomFish returned non-nil when initialized nil")
	}
}

func Test_AddFish_OccupiedSlot_ReturnsFalse(t *testing.T) {
	coll := setupCollection()
	if !coll.AddFish(fish.NewFish(ENV, "test1", "test-fish1", nil), 0) {
		t.Error("Couldnt add fish to empty slot!")
	}
	retVal := coll.AddFish(fish.NewFish(ENV, "test2", "test-fish2", nil), 0)
	if retVal {
		t.Error("Fish added to the same slot twice!")
	}
}
func Test_AddFish_IndexOutOfBounds_ReturnsFalse(t *testing.T) {
	coll := setupCollection()
	if coll.AddFish(fish.NewFish(ENV, "t", "t", nil), 999) {
		t.Error("Fish cannot be added to index 999")
	}
}

func Test_GetAllFish_WithOneFish_ReturnsFishSlice(t *testing.T) {
	coll := setupCollection()
	coll.AddFish(fish.NewFish(ENV, "test", "test-fish", nil), 0)
	retVal := coll.GetAllFish()
	if len(retVal) != 5 {
		t.Errorf("Slice returned was:%v length and should be 5", len(retVal))
	}
	if retVal[0].Name != "test" && retVal[0].Description != "test-fish" {
		t.Errorf("Wrong fish returned")
	}
}
func Test_GetAllFish_WithTwoFish_ReturnsFishSlice(t *testing.T) {
	coll := setupCollection()
	added1 := coll.AddFish(fish.NewFish(ENV, "test1", "test-fish1", nil), 0)
	added2 := coll.AddFish(fish.NewFish(ENV, "test2", "test-fish2", nil), 4)
	if !added1 || !added2 {
		t.Error("Unable to add initial 2 fish successfully")
	}
	retVal := coll.GetAllFish()
	if len(retVal) != 5 {
		t.Errorf("Slice returned was:%v length and should be 5", len(retVal))
	}
	if retVal[0].Name != "test1" && retVal[0].Description != "test-fish1" {
		t.Error("Wrong fish1 returned")
	}
	if retVal[4].Name != "test2" && retVal[4].Description != "test-fish2" {
		t.Error("Wrong fish2 returned")
	}
}
func Test_GetAllFish_WithNoFish_ReturnsAllNil(t *testing.T) {
	coll := setupCollection()
	retVal := coll.GetAllFish()
	if len(retVal) != 5 {
		t.Errorf("Slice returned was:%v length and should be 5", len(retVal))
	}
	for _, fish := range retVal {
		if fish != nil {
			t.Error("returned fish was non-nil and should be nil")
		}
	}
}

func Test_IndexEmpty_WithNoFishAndAllIndexes_ReturnsTrue(t *testing.T) {
	coll := setupCollection()
	for i := 0; i < 5; i++ {
		if !coll.IndexEmpty(i) {
			t.Errorf("Cannot add fish to index:%v and it should be able to.", i)
		}
	}
}
func Test_IndexEmpty_WithFiveFishAndAllIndexes_ReturnsFalse(t *testing.T) {
	coll := setupCollection()
	coll.AddFish(fish.NewFish(ENV, "test", "test-fish", nil), 0)
	coll.AddFish(fish.NewFish(ENV, "test", "test-fish", nil), 1)
	coll.AddFish(fish.NewFish(ENV, "test", "test-fish", nil), 2)
	coll.AddFish(fish.NewFish(ENV, "test", "test-fish", nil), 3)
	coll.AddFish(fish.NewFish(ENV, "test", "test-fish", nil), 4)
	for i := 0; i < 5; i++ {
		if coll.IndexEmpty(i) {
			t.Errorf("Can add fish to index:%v and it should be full already.", i)
		}
	}
}

func Test_MoveFish_SwapFishAndEmptySlot_ReturnsTrue(t *testing.T) {
	coll := setupCollection()
	added := coll.AddFish(fish.NewFish(ENV, "test1", "test-fish1", nil), 3)
	if !added {
		t.Error("Failed to add fish to collection")
	}
	initialFishState := coll.GetAllFish()
	if initialFishState[0] != nil {
		t.Error("empty index was non-empty")
	}
	if initialFishState[3].Name != "test1" {
		t.Error("fish in index is incorrect")
	}
	success := coll.MoveFish(3, 2)
	if !success {
		t.Error("Fish failed to swap locations")
	}
	afterFishState := coll.GetAllFish()
	if afterFishState[3] != nil {
		t.Error("source fish was not removed")
	}
	if afterFishState[2].Description != "test-fish1" {
		t.Error("Wrong fish in target index")
	}

}
func Test_MoveFish_SwapTwoFish_ReturnsTrue(t *testing.T) {
	coll := setupCollection()
	added1 := coll.AddFish(fish.NewFish(ENV, "eel", "he long", nil), 2)
	added2 := coll.AddFish(fish.NewFish(ENV, "goldfish", "he golden", nil), 4)
	if !added1 || !added2 {
		t.Error("unable to add initial 2 fish")
	}
	success := coll.MoveFish(2, 4)
	if !success {
		t.Error("Unable to swap fish!")
	}
	afterFishState := coll.GetAllFish()
	if afterFishState[2].Name != "goldfish" || afterFishState[2].Description != "he golden" {
		t.Error("target fish was not moved to source slot")
	}
	if afterFishState[4].Name != "eel" || afterFishState[4].Description != "he long" {
		t.Error("source fish was not moved to target slot")
	}

}
func Test_MoveFish_WithNilSourceFish_ReturnsFalse(t *testing.T) {
	coll := setupCollection()
	success := coll.MoveFish(0, 2)
	if success {
		t.Error("source fish was empty and false should have been returned")
	}
}

func Test_DisableChanges_Default_PreventsChanges(t *testing.T) {
	coll := setupCollection()
	ret := coll.AddFish(fish.NewFish(ENV, "test", "test-fish", nil), 4)
	if !ret {
		t.Error("unable to add initial fish")
	}
	coll.DisableChanges()
	ret = coll.AddFish(fish.NewFish(ENV, "testfish2", "test-fish2", nil), 2)
	if ret {
		t.Error("Fish should not have been added to collection but it was.")
	}
}

func Test_EnableChanges_Default_EnablesChanges(t *testing.T) {
	coll := setupCollection()
	coll.DisableChanges()
	ret := coll.AddFish(fish.NewFish(ENV, "test", "test-fish", nil), 4)
	if ret {
		t.Error("Fish should not have been added to collection but it was.")
	}
	coll.EnableChanges()
	ret = coll.AddFish(fish.NewFish(ENV, "test", "test-fish", nil), 4)
	if !ret {
		t.Error("unable to add fish and it should have been added")
	}
}

func Test_AllFishDead_WithAllDeadFish_ReturnsTrue(t *testing.T) {
	coll := setupCollection()
	for i := 0; i < 5; i++ {
		stats := fish.NewWeaponStats(1, 1, 1)
		fish := fish.NewFish(ENV, "test", "test-fish", &stats)
		fish.TakeDamage(1)
		coll.AddFish(fish, i)
	}
	if !coll.AllFishDead() {
		t.Error("all fish should be dead!")
	}
}
func Test_AllFishDead_WithAllAliveFish_ReturnsFalse(t *testing.T) {
	coll := setupCollection()
	for i := 0; i < 5; i++ {
		stats := fish.NewWeaponStats(1, 1, 1)
		fish := fish.NewFish(ENV, "test", "test-fish", &stats)
		coll.AddFish(fish, i)
	}
	if coll.AllFishDead() {
		t.Error("all fish should be alive!")
	}
}

func Test_ById_WithValidString_ReturnsFishAndIndex(t *testing.T) {
	coll := setupCollection()
	stats := fish.NewWeaponStats(1, 1, 1)
	fish := fish.NewFish(ENV, "asdfgh", "test-fish", &stats)
	coll.AddFish(fish, 2)

	foundIndex, foundFish := coll.ById(fish.Id.String())
	if foundFish.Name != "asdfgh" {
		t.Error("found fish should have same name")
	}
	if foundIndex != 2 {
		t.Error("found fish should have same index")
	}
}
func Test_ById_WithInvalidString_ReturnsNilAnd99(t *testing.T) {
	coll := setupCollection()
	stats := fish.NewWeaponStats(1, 1, 1)
	fish := fish.NewFish(ENV, "zxcvbn", "test-fish", &stats)
	coll.AddFish(fish, 2)

	foundIndex, foundFish := coll.ById("test-fish")
	if foundFish != nil {
		t.Error("fish should not have been found from invalid ID")
	}
	if foundIndex != 99 {
		t.Error("fish index should be 999 for invalid")
	}
}
