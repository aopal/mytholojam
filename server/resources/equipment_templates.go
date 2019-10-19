package resources

import (
	"mytholojam/server/types"
)

// SAMPLE EQUPIMENT
var Sword types.EquipmentTemplate = types.EquipmentTemplate{
	Name:   "Sword",
	MaxHP:  avgHP,
	ATK:    highATK,
	Weight: avgWeight,
	OnHit:  noop,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		moveTypes[0]: lowDEF,
		moveTypes[1]: lowDEF,
		moveTypes[2]: lowDEF,
		moveTypes[3]: lowDEF,
	},
	Moves: []*types.Move{},
}

var Shield types.EquipmentTemplate = types.EquipmentTemplate{
	Name:   "Shield",
	MaxHP:  highHP,
	ATK:    lowATK,
	Weight: highWeight,
	OnHit:  noop,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		moveTypes[0]: highDEF,
		moveTypes[1]: highDEF,
		moveTypes[2]: highDEF,
		moveTypes[3]: lowDEF,
	},
	Moves: []*types.Move{},
}

var EquipList map[string]*types.EquipmentTemplate = map[string]*types.EquipmentTemplate{
	Sword.Name:  &Sword,
	Shield.Name: &Shield,
}
