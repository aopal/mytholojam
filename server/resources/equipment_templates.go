package resources

import (
	"mytholojam/server/types"
)

var EquipList map[string]*types.EquipmentTemplate = map[string]*types.EquipmentTemplate{
	Sword.Name:  &Sword,
	Shield.Name: &Shield,
	Bow.Name:    &Bow,
}

// SAMPLE EQUPIMENT
var Sword types.EquipmentTemplate = types.EquipmentTemplate{
	Name:   "Sword",
	MaxHP:  avgHP,
	ATK:    highATK,
	Weight: avgWeight,
	OnHit:  defaultOnHit,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		spiritTypes[0]: lowDEF,
		spiritTypes[1]: lowDEF,
		spiritTypes[2]: lowDEF,
		spiritTypes[3]: lowDEF,
		switchType:     switchDEF,
	},
	Moves: []*types.Move{},
}

var Shield types.EquipmentTemplate = types.EquipmentTemplate{
	Name:   "Shield",
	MaxHP:  highHP,
	ATK:    lowATK,
	Weight: highWeight,
	OnHit:  defaultOnHit,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		spiritTypes[0]: highDEF,
		spiritTypes[1]: highDEF,
		spiritTypes[2]: highDEF,
		spiritTypes[3]: lowDEF,
		switchType:     switchDEF,
	},
	Moves: []*types.Move{},
}

var Bow types.EquipmentTemplate = types.EquipmentTemplate{
	Name:   "Bow",
	MaxHP:  avgHP,
	ATK:    highATK,
	Weight: lowWeight,
	OnHit:  defaultOnHit,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		spiritTypes[0]: avgDEF,
		spiritTypes[1]: lowDEF,
		spiritTypes[2]: lowDEF,
		spiritTypes[3]: avgDEF,
		switchType:     switchDEF,
	},
	Moves: []*types.Move{},
}
