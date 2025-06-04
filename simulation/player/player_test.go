package player

import (
	"fishgame/shared/environment"
	"testing"
)

// [UnitOfWork_StateUnderTest_ExpectedBehaviour]
// https://osherove.com/blog/2005/4/3/naming-standards-for-unit-tests.html

func setupTestPlayer() *Player {
	env := environment.NewEnv(nil, nil)
	return NewPlayer(env, "test-player")
}
func Test_NewPlayer_Default_InitializesNilorEmpty(t *testing.T) {
	p := setupTestPlayer()

	if p.currency != 0 {
		t.Error("Initial player currency should be zero!")
	}

	if len(p.Fish.GetAllFish()) != 5 {
		t.Error("Initial Player Fish count should be 5")
	}
	for _, fish := range p.Fish.GetAllFish() {
		if fish != nil {
			t.Error("Initial Player Fish should all be nil!")
		}
	}
}

func Test_SpendCurrency_MoreThanOwned_DoesntAllowNegativeCurrency(t *testing.T) {
	p := setupTestPlayer()
	p.AddCurrency(90)

	if p.SpendCurrency(100) {
		t.Error("Cannot spend more currency than owned!")
	}
	if p.GetCurrencyAmount() != 90 {
		t.Error("Currency number should not have changed!")
	}
}
