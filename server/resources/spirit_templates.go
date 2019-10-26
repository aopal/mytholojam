package resources

import "mytholojam/server/types"

var SpiritList map[string]*types.SpiritTemplate = map[string]*types.SpiritTemplate{
	Warrior.Name: &Warrior,
	Thief.Name:   &Thief,
}

// SAMPLE SPIRITS
var Warrior types.SpiritTemplate = types.SpiritTemplate{
	Name:   "Warrior",
	MaxHP:  highHP,
	ATK:    avgATK,
	Speed:  avgSPD,
	OnHit:  defaultOnHit,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		spiritTypes[0]: lowDEF,
		spiritTypes[1]: avgDEF,
		spiritTypes[2]: avgDEF,
		spiritTypes[3]: lowDEF,
		switchType:     switchDEF,
	},
	Moves: []*types.Move{
		&Switch,
		&Strong,
		&Weak,
	},
}

var Thief types.SpiritTemplate = types.SpiritTemplate{
	Name:   "Thief",
	MaxHP:  lowHP,
	ATK:    avgATK,
	Speed:  highSPD,
	OnHit:  defaultOnHit,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		spiritTypes[0]: lowDEF,
		spiritTypes[1]: avgDEF,
		spiritTypes[2]: lowDEF,
		spiritTypes[3]: avgDEF,
		switchType:     switchDEF,
	},
	Moves: []*types.Move{
		&Switch,
		&Weak,
		&Fast,
	},
}

// FIRST PASS SPIRITS
var Flame types.SpiritTemplate = types.SpiritTemplate{
	Name:   "Flame",
	MaxHP:  highHP,
	ATK:    avgATK,
	Speed:  avgSPD,
	OnHit:  defaultOnHit,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		spiritTypes[0]: midloDEF,
		spiritTypes[1]: infDEF,
		spiritTypes[2]: lowDEF,
		spiritTypes[3]: lowDEF,
		switchType:     switchDEF,
	},
	Moves: []*types.Move{
		&Switch,
	},
}

var Hive types.SpiritTemplate = types.SpiritTemplate{
	Name:   "Hive",
	MaxHP:  highHP,
	ATK:    avgATK,
	Speed:  avgSPD,
	OnHit:  defaultOnHit,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		spiritTypes[0]: midhiDEF,
		spiritTypes[1]: midloDEF,
		spiritTypes[2]: highDEF,
		spiritTypes[3]: lowDEF,
		switchType:     switchDEF,
	},
	Moves: []*types.Move{
		&Switch,
	},
}
