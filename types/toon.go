package types

import "fmt"

// Toon is a character model for toons
type Toon struct {
	ToonBase
	key             string
	starLevel       int //max 7
	level           int //max 80
	gearLevel       int // max 11
	unusedShards    int
	promotionStatus bool //enough shards to promote?
	starsToUnlock   int
	// Multipliers //intelligence, strength, agility by star level
	// Stats //
	abilities []Ability
	// mods []Mod
	// gear []Gear will be gear info
}

// ToonBase is the unique character information that doesn't change
type ToonBase struct {
	name     string
	toonType string //attacker, support, tank
	// alliances []Alliance
}

// AddShards increments the shards for a toon and alters the promotionStatus
// if a toon is ready to promote
func (t *Toon) AddShards(shards int) {
	t.unusedShards += shards
	if t.unusedShards > shardMath(t.starLevel) {
		t.promotionStatus = true
	}
}

// Promote bumps a star level if a promotion is available
// Converts unusedshards to reflect new total
// TODO: Pass the user interface to spend credits on promotion
func (t *Toon) Promote() error {
	if t.promotionStatus != true {
		return fmt.Errorf("Cannot promote toon: %s, not enough shards", t.key)
	}

	t.unusedShards = t.unusedShards - shardMath(t.starLevel)

	t.starLevel++

	return nil
}

func shardMath(currentLevel int) int {
	switch currentLevel {
	case 1:
		return 10
	case 2:
		return 25
	case 3:
		return 30
	case 4:
		return 65
	case 5:
		return 85
	case 6:
		return 100
	default:
		return -1
	}
}
