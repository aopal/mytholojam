package resources

import "mytholojam/server/types"

// SAMPLE SPIRITS
var Warrior types.SpiritTemplate = types.SpiritTemplate{
	Name:   "Warrior",
	MaxHP:  highHP,
	ATK:    highATK,
	Speed:  avgSPD,
	OnHit:  noop,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		moveTypes[0]: lowDEF,
		moveTypes[1]: avgDEF,
		moveTypes[2]: avgDEF,
		moveTypes[3]: lowDEF,
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
	ATK:    highATK,
	Speed:  highSPD,
	OnHit:  noop,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		moveTypes[0]: lowDEF,
		moveTypes[1]: avgDEF,
		moveTypes[2]: lowDEF,
		moveTypes[3]: avgDEF,
	},
	Moves: []*types.Move{
		&Switch,
		&Weak,
		&Fast,
	},
}

var SpiritList map[string]*types.SpiritTemplate = map[string]*types.SpiritTemplate{
	Warrior.Name: &Warrior,
	Thief.Name:   &Thief,
}
