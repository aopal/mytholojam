package resources

import "mytholojam/server/types"

var SpiritList map[string]*types.SpiritTemplate = map[string]*types.SpiritTemplate{
	Warrior.Name: &Warrior,
	Thief.Name:   &Thief,
	Flame.Name:   &Flame,
	Hive.Name:    &Hive,
}

// SAMPLE SPIRITS
var Warrior types.SpiritTemplate = types.SpiritTemplate{
	Name:   "Warrior",
	MaxHP:  highHP,
	ATK:    avgATK,
	Speed:  avgSPD,
	OnHit:  empty,
	OnMiss: empty,
	OnDbl:  empty,
	Defs: map[string]int{
		spiritTypes[0]: lowDEF,
		spiritTypes[1]: avgDEF,
		spiritTypes[2]: avgDEF,
		spiritTypes[3]: lowDEF,
	},
	Moves: map[string]*types.Move{
		Switch.Name: &Switch,
		Strong.Name: &Strong,
		Weak.Name:   &Weak,
	},
}

var Thief types.SpiritTemplate = types.SpiritTemplate{
	Name:   "Thief",
	MaxHP:  lowHP,
	ATK:    avgATK,
	Speed:  highSPD,
	OnHit:  empty,
	OnMiss: empty,
	OnDbl:  empty,
	Defs: map[string]int{
		spiritTypes[0]: lowDEF,
		spiritTypes[1]: avgDEF,
		spiritTypes[2]: lowDEF,
		spiritTypes[3]: avgDEF,
	},
	Moves: map[string]*types.Move{
		Switch.Name: &Switch,
		Weak.Name:   &Weak,
		Fast.Name:   &Fast,
	},
}

// FIRST PASS SPIRITS
var Flame types.SpiritTemplate = types.SpiritTemplate{
	Name:   "Flame",
	MaxHP:  highHP,
	ATK:    avgATK,
	Speed:  avgSPD,
	OnHit:  empty,
	OnMiss: empty,
	OnDbl:  empty,
	Defs: map[string]int{
		spiritTypes[0]: midloDEF,
		spiritTypes[1]: infDEF,
		spiritTypes[2]: lowDEF,
		spiritTypes[3]: lowDEF,
	},
	Moves: map[string]*types.Move{
		Switch.Name: &Switch,
		Strong.Name: &Strong,
		Weak.Name:   &Weak,
		Fast.Name:   &Fast,
	},
}

var Hive types.SpiritTemplate = types.SpiritTemplate{
	Name:   "Hive",
	MaxHP:  highHP,
	ATK:    avgATK,
	Speed:  avgSPD,
	OnHit:  empty,
	OnMiss: empty,
	OnDbl:  empty,
	Defs: map[string]int{
		spiritTypes[0]: midhiDEF,
		spiritTypes[1]: midloDEF,
		spiritTypes[2]: highDEF,
		spiritTypes[3]: lowDEF,
	},
	Moves: map[string]*types.Move{
		Switch.Name: &Switch,
		Strong.Name: &Strong,
		Weak.Name:   &Weak,
		Fast.Name:   &Fast,
	},
}
