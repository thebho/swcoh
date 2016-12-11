package types

// Ability is a action or passive stat that belongs to a Toon, and can
// be promoted
type Ability struct {
	AbilityBase
	level int
}

// AbilityBase is the static information about an Ability
type AbilityBase struct {
	name string
	toon string
	// enhancement []Enhancement
}
