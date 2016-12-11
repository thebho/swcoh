package types

import (
	"testing"
)

func TestAddShards(t *testing.T) {
	toon := Toon{starLevel: 1, unusedShards: 2}
	toon.AddShards(10)

	if toon.unusedShards != 12 {
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
		Toon{promotionStatus: true, unusedShards: 11, starLevel: 1},
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
