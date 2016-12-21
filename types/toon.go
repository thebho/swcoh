package types

import "fmt"

// Toon is a character model for toons
type Toon struct {
	toonBase
	key             string
	unlocked        bool
	starLevel       int //max 7
	level           int //max 80
	gearLevel       int // max 11
	unusedShards    int
	promotionStatus bool //enough shards to promote?
	starsToUnlock   int
	toonBasicStats
	// Stats //
	abilities []Ability
	// mods []Mod
	// gear []Gear will be gear info
}

// ToonBase is the unique character information that doesn't change
type toonBase struct {
	name     string
	toonType string //attacker, support, tank
	// alliances []Alliance
}

type toonBasicStats struct {
	strength int // Health and Armor
	agility  int // Physical Damage
	tactics  int // Special Damage and Resistance
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

// CanUnlock returns a true if enough shards have been collected an a toon
// has not been previously unlocked
func (t *Toon) CanUnlock() bool {
	if t.unlocked {
		return false
	}

	if t.unusedShards >= unlockMath(t.starsToUnlock) {
		return true
	}

	return false
}

// Unlock converts unusedShards to Unlock a toon if the toon has reached the threshhold
func (t *Toon) Unlock() error {
	if t.unlocked {
		return fmt.Errorf("Attempting to unlock and alread unlocked toon")
	}

	unlockShards := unlockMath(t.starsToUnlock)
	if t.unusedShards >= unlockMath(t.starsToUnlock) {
		t.unlocked = true
		t.unusedShards = t.unusedShards - unlockShards
		return nil
	}
	return fmt.Errorf("Unable to unlock")
}

func unlockMath(unlockLevel int) int {
	switch unlockLevel {
	case 1:
		return 10
	case 2:
		return 25
	case 3:
		return 50
	case 4:
		return 80
	case 5:
		return 145
	}
	return -1
}

func shardMath(currentLevel int) int {
	switch currentLevel {
	case 1:
		return 15
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
