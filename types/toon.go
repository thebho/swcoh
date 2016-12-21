package types

import "fmt"

// Toon is a character model for toons
type Toon struct {
	ToonBase
	Key             string
	Unlocked        bool
	StarLevel       int //max 7
	Level           int //max 80
	GearLevel       int // max 11
	UnusedShards    int
	PromotionStatus bool //enough shards to promote?
	StarsToUnlock   int
	ToonBasicStatsMultiplier
	// Stats //
	Abilities []Ability
	// Mods []Mod
	// Gear []Gear will be gear info
}

// ToonBase is the unique character information that doesn't change
type ToonBase struct {
	Name     string
	ToonType string //attacker, support, tank
	// alliances []Alliance
}

// ToonBasicStatsMultiplier are static values that multiple by level and starLevel to display power
type ToonBasicStatsMultiplier struct {
	StrengthMultiplier int // Health and Armor
	AgilityMultiplier  int // Physical Damage
	TacticsMultiplier  int // Special Damage and Resistance
}

// AddShards increments the shards for a toon and alters the promotionStatus
// if a toon is ready to promote
func (t *Toon) AddShards(shards int) {
	t.UnusedShards += shards
	if t.UnusedShards > shardMath(t.StarLevel) {
		t.PromotionStatus = true
	}
}

// Promote bumps a star level if a promotion is available
// Converts unusedshards to reflect new total
// TODO: Pass the user interface to spend credits on promotion
func (t *Toon) Promote() error {
	if t.PromotionStatus != true {
		return fmt.Errorf("Cannot promote toon: %s, not enough shards", t.Key)
	}

	t.UnusedShards = t.UnusedShards - shardMath(t.StarLevel)

	t.StarLevel++

	return nil
}

// CanUnlock returns a true if enough shards have been collected an a toon
// has not been previously Unlocked
func (t *Toon) CanUnlock() bool {
	if t.Unlocked {
		return false
	}

	if t.UnusedShards >= unlockMath(t.StarsToUnlock) {
		return true
	}

	return false
}

// Unlock converts UnusedShards to Unlock a toon if the toon has reached the threshhold
func (t *Toon) Unlock() error {
	if t.Unlocked {
		return fmt.Errorf("Attempting to unlock and alread unlocked toon")
	}

	unlockShards := unlockMath(t.StarsToUnlock)
	if t.UnusedShards >= unlockMath(t.StarsToUnlock) {
		t.Unlocked = true
		t.UnusedShards = t.UnusedShards - unlockShards
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

// Stats is dynamically created by multiplying level * starLevel * multiplier
type Stats struct {
	Strength int
	Agility  int
	Tactics  int
}

// GetBasicStatPower returns the current power of str, agi, tac
func (t *Toon) GetBasicStatPower() Stats {
	fmt.Println("Test")
	starLevelLevel := t.StarLevel * t.Level
	str := t.StrengthMultiplier * starLevelLevel
	agi := t.AgilityMultiplier * starLevelLevel
	tac := t.TacticsMultiplier * starLevelLevel
	return Stats{str, agi, tac}
}
