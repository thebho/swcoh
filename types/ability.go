package types

// Ability is a action or passive stat that belongs to a Toon, and can
// be promoted
type Ability struct {
	abilityBase
	abilityType AbilityType
	level       int
	cooldown    int
	unlocked    bool
}

// NewAbility returns a new
func NewAbility(abilityName, abilityToon string, aType AbilityType, level, cooldown int, unlocked bool) Ability {
	return Ability{
		abilityBase: abilityBase{name: abilityName, toon: abilityToon},
		abilityType: aType,
		level:       level,
		cooldown:    cooldown,
		unlocked:    unlocked,
	}
}

// AbilityType is a
type AbilityType int

const (
	//AbilityBasic is an ability every toon has
	AbilityBasic AbilityType = iota
	//AbilityLeader ability used in leader slot
	AbilityLeader
	//AbilitySpecial is a secondary ability with cooldown
	AbilitySpecial
	//AbilityUnique is a passive ability
	AbilityUnique
)

// AbilityBase is the static information about an Ability
type abilityBase struct {
	name string
	toon string
	// enhancement []Enhancement
}

// AbilityType stringifies ability types
func (a Ability) AbilityType() string {
	switch a.abilityType {
	case AbilityBasic:
		return "basic"
	case AbilityLeader:
		return "leader"
	case AbilitySpecial:
		return "special"
	case AbilityUnique:
		return "unique"
	default:
		return ""
	}
}
