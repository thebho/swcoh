package types

import (
	"reflect"
	"testing"
)

func TestAddShards(t *testing.T) {
	toon := Toon{StarLevel: 1, UnusedShards: 2}
	toon.AddShards(14)

	if toon.UnusedShards != 16 {
		t.Error("Shard total didn't add up")
	}

	if toon.PromotionStatus != true {
		t.Error("Went from 1 to 11 shards and didn't allow permission")
	}

}

func TestBadPromotion(t *testing.T) {
	toon := Toon{PromotionStatus: false}
	err := toon.Promote()

	if err == nil {
		t.Errorf("Expected Promotion to fail")
	}
}

func TestPromotion(t *testing.T) {
	promotionSlice := []Toon{
		Toon{PromotionStatus: true, UnusedShards: 16, StarLevel: 1},
		Toon{PromotionStatus: true, UnusedShards: 26, StarLevel: 2},
		Toon{PromotionStatus: true, UnusedShards: 31, StarLevel: 3},
		Toon{PromotionStatus: true, UnusedShards: 66, StarLevel: 4},
		Toon{PromotionStatus: true, UnusedShards: 86, StarLevel: 5},
		Toon{PromotionStatus: true, UnusedShards: 101, StarLevel: 6},
	}
	// toon := Toon{promotionStatus: true, unusedShards: 12, starLevel: 1}
	for _, toon := range promotionSlice {
		startLevel := toon.StarLevel
		err := toon.Promote()

		if err != nil {
			t.Errorf("Expected promotion to succeed")
		}

		if toon.UnusedShards != 1 {
			t.Errorf("Unusued shards not accurate after promotion to %d stars", toon.StarLevel)
		}

		if toon.StarLevel != startLevel+1 {
			t.Errorf("Expected a promotion from: %d stars", toon.StarLevel)
		}
	}
}

func TestUnlockingPreviouslyUnlocked(t *testing.T) {
	toon := Toon{Unlocked: true}
	err := toon.Unlock()

	if err == nil {
		t.Errorf("Expecting unlocked toon to through an error")
	}
}

func TestUnlock(t *testing.T) {
	promotionSlice := []Toon{
		Toon{UnusedShards: 10, StarsToUnlock: 1},
		Toon{UnusedShards: 25, StarsToUnlock: 2},
		Toon{UnusedShards: 50, StarsToUnlock: 3},
		Toon{UnusedShards: 80, StarsToUnlock: 4},
		Toon{UnusedShards: 145, StarsToUnlock: 5},
	}

	for _, toon := range promotionSlice {
		if !toon.CanUnlock() {
			t.Errorf("CanUnlocks failed: %d, %d", toon.UnusedShards, toon.StarsToUnlock)

		}
		err := toon.Unlock()
		if err != nil {
			t.Errorf("Failed to unlock: %d, %d, %v", toon.UnusedShards, toon.StarsToUnlock, err)
		}
	}
}

func TestBadUnlock(t *testing.T) {
	promotionSlice := []Toon{
		Toon{UnusedShards: 9, StarsToUnlock: 1},
		Toon{UnusedShards: 20, StarsToUnlock: 2},
		Toon{UnusedShards: 45, StarsToUnlock: 3},
		Toon{UnusedShards: 79, StarsToUnlock: 4},
		Toon{UnusedShards: 143, StarsToUnlock: 5},
	}

	for _, toon := range promotionSlice {
		if toon.CanUnlock() {
			t.Errorf("CanUnlocks failed: %d, %d", toon.UnusedShards, toon.StarsToUnlock)

		}
		err := toon.Unlock()
		if err == nil {
			t.Errorf("Expceted failure to unlock: %d, %d, %v", toon.UnusedShards, toon.StarsToUnlock, err)
		}
	}
}

func TestUnlockingUnlocked(t *testing.T) {
	toon := Toon{Unlocked: true}
	if toon.CanUnlock() {
		t.Errorf("Shouldn't be able to unlock %v", toon.Unlocked)
	}
}

func TestBasicStatPower(t *testing.T) {
	level := 10
	starLevel := 2
	toon := Toon{
		ToonBasicStatsMultiplier: ToonBasicStatsMultiplier{
			AgilityMultiplier:  2,
			StrengthMultiplier: 3,
			TacticsMultiplier:  4},
		Level:     10,
		StarLevel: 2}

	statAssert := Stats{
		Agility:  2 * level * starLevel,
		Strength: 3 * level * starLevel,
		Tactics:  4 * level * starLevel,
	}

	if !reflect.DeepEqual(statAssert, toon.GetBasicStatPower()) {
		t.Error("Basic Stat Math not accurate")
	}

	level = 20
	starLevel++
	toon.Level += 10
	toon.StarLevel = 3

	statAssert = Stats{
		Agility:  2 * level * starLevel,
		Strength: 3 * level * starLevel,
		Tactics:  4 * level * starLevel,
	}

	if !reflect.DeepEqual(statAssert, toon.GetBasicStatPower()) {
		t.Error("Basic Stat Math not accurate")
	}

}
