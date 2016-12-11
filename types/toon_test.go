package types

import (
	"testing"
)

func TestAddShards(t *testing.T) {
	toon := Toon{starLevel: 1, unusedShards: 2}
	toon.AddShards(14)

	if toon.unusedShards != 16 {
		t.Error("Shard total didn't add up")
	}

	if toon.promotionStatus != true {
		t.Error("Went from 1 to 11 shards and didn't allow permission")
	}

}

func TestBadPromotion(t *testing.T) {
	toon := Toon{promotionStatus: false}
	err := toon.Promote()

	if err == nil {
		t.Errorf("Expected promotion to fail")
	}
}

func TestPromotion(t *testing.T) {
	promotionSlice := []Toon{
		Toon{promotionStatus: true, unusedShards: 16, starLevel: 1},
		Toon{promotionStatus: true, unusedShards: 26, starLevel: 2},
		Toon{promotionStatus: true, unusedShards: 31, starLevel: 3},
		Toon{promotionStatus: true, unusedShards: 66, starLevel: 4},
		Toon{promotionStatus: true, unusedShards: 86, starLevel: 5},
		Toon{promotionStatus: true, unusedShards: 101, starLevel: 6},
	}
	// toon := Toon{promotionStatus: true, unusedShards: 12, starLevel: 1}
	for _, toon := range promotionSlice {
		startLevel := toon.starLevel
		err := toon.Promote()

		if err != nil {
			t.Errorf("Expected promotion to succeed")
		}

		if toon.unusedShards != 1 {
			t.Errorf("Unusued shards not accurate after promotion to %d stars", toon.starLevel)
		}

		if toon.starLevel != startLevel+1 {
			t.Errorf("Expected a promotion from: %d stars", toon.starLevel)
		}
	}
}

func TestUnlockingPreviouslyUnlocked(t *testing.T) {
	toon := Toon{unlocked: true}
	err := toon.Unlock()

	if err == nil {
		t.Errorf("Expecting unlocked toon to through an error")
	}
}

func TestUnlock(t *testing.T) {
	promotionSlice := []Toon{
		Toon{unusedShards: 10, starsToUnlock: 1},
		Toon{unusedShards: 25, starsToUnlock: 2},
		Toon{unusedShards: 50, starsToUnlock: 3},
		Toon{unusedShards: 80, starsToUnlock: 4},
		Toon{unusedShards: 145, starsToUnlock: 5},
	}

	for _, toon := range promotionSlice {
		if !toon.CanUnlock() {
			t.Errorf("CanUnlocks failed: %d, %d", toon.unusedShards, toon.starsToUnlock)

		}
		err := toon.Unlock()
		if err != nil {
			t.Errorf("Failed to unlock: %d, %d, %v", toon.unusedShards, toon.starsToUnlock, err)
		}
	}
}

func TestBadUnlock(t *testing.T) {
	promotionSlice := []Toon{
		Toon{unusedShards: 9, starsToUnlock: 1},
		Toon{unusedShards: 20, starsToUnlock: 2},
		Toon{unusedShards: 45, starsToUnlock: 3},
		Toon{unusedShards: 79, starsToUnlock: 4},
		Toon{unusedShards: 143, starsToUnlock: 5},
	}

	for _, toon := range promotionSlice {
		if toon.CanUnlock() {
			t.Errorf("CanUnlocks failed: %d, %d", toon.unusedShards, toon.starsToUnlock)

		}
		err := toon.Unlock()
		if err == nil {
			t.Errorf("Expceted failure to unlock: %d, %d, %v", toon.unusedShards, toon.starsToUnlock, err)
		}
	}
}

func TestUnlockingUnlocked(t *testing.T) {
	toon := Toon{unlocked: true}
	if toon.CanUnlock() {
		t.Errorf("Shouldn't be able to unlock %v", toon.unlocked)

	}
}
