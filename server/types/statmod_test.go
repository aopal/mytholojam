package types

import (
	"testing"

	"gotest.tools/assert"
)

func TestStatmods(t *testing.T) {
	var s Spirit

	s.ATK = 5
	s.Speed = 10
	s.Defs = map[string]int{
		"type1": 1,
		"type2": 2,
	}

	s.StatMods = map[string]*StatMod{
		"id1": &StatMod{
			AtkMod:   4,
			SpeedMod: 3,
			DefMods: map[string]int{
				"type1": 1,
				"type2": 0,
			},
		},
		"id2": &StatMod{
			AtkMod:   -2,
			SpeedMod: -3,
			DefMods: map[string]int{
				"type1": 0,
				"type2": -1,
			},
		},
	}

	assert.Assert(t, s.GetAtk() == s.ATK+2)
	assert.Assert(t, s.GetSpeed() == s.Speed)
	assert.Assert(t, s.GetDef("type1") == s.Defs["type1"]+1)
	assert.Assert(t, s.GetDef("type2") == s.Defs["type2"]-1)
}
