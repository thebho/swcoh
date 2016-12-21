package types

import "testing"

func TestNewAbility(t *testing.T) {
	ability := Ability{
		abilityBase: abilityBase{name: "Basic", toon: "Rey"},
		abilityType: AbilityBasic,
		level:       1,
		cooldown:    0,
		unlocked:    false,
	}
	abilityCopy := NewAbility("Basic", "Rey", AbilityBasic, 1, 0, false)
	if ability != abilityCopy {
		t.Error("Abilities don't match")
	}
}

func TestAbilityToString(t *testing.T) {
	ability := Ability{abilityType: AbilityBasic}
	if ability.AbilityType() != "basic" {
		t.Error("AbilityBasic mismatch")
	}
	ability = Ability{abilityType: AbilityLeader}
	if ability.AbilityType() != "leader" {
		t.Error("AbilityLeader mismatch")
	}
	ability = Ability{abilityType: AbilitySpecial}
	if ability.AbilityType() != "special" {
		t.Error("AbilitySpecial mismatch")
	}
	ability = Ability{abilityType: AbilityUnique}
	if ability.AbilityType() != "unique" {
		t.Error("AbilityUnique mismatch")
	}
	ability = Ability{abilityType: 5}
	if ability.AbilityType() != "" {
		t.Error("Ability Type mismatch")
	}
}
