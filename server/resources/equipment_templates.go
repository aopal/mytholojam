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

// FIRST PASS EQUIPMENT
var AngloSaxonSpear types.EquipmentTemplate = types.EquipmentTemplate{
	Name:   "AngloSaxonSpear",
	MaxHP:  avgHP,
	ATK:    avgATK,
	Weight: avgWeight,
	OnHit:  defaultOnHit,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		spiritTypes[0]: midhiDEF,
		spiritTypes[1]: midloDEF,
		spiritTypes[2]: avgDEF,
		spiritTypes[3]: avgDEF,
		switchType:     switchDEF,
	},
	Moves: []*types.Move{},
}

var DanishNeedleSet types.EquipmentTemplate = types.EquipmentTemplate{
	Name:   "DanishNeedleSet",
	MaxHP:  avgHP,
	ATK:    avgATK,
	Weight: avgWeight,
	OnHit:  defaultOnHit,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		spiritTypes[0]: midloDEF,
		spiritTypes[1]: highDEF,
		spiritTypes[2]: highDEF,
		spiritTypes[3]: avgDEF,
		switchType:     switchDEF,
	},
	Moves: []*types.Move{},
}

var ArthiRod types.EquipmentTemplate = types.EquipmentTemplate{
	Name:   "ArthiRod",
	MaxHP:  avgHP,
	ATK:    avgATK,
	Weight: avgWeight,
	OnHit:  defaultOnHit,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		spiritTypes[0]: avgDEF,
		spiritTypes[1]: lowDEF,
		spiritTypes[2]: midloDEF,
		spiritTypes[3]: midhiDEF,
		switchType:     switchDEF,
	},
	Moves: []*types.Move{},
}

var IberianGoatSkull types.EquipmentTemplate = types.EquipmentTemplate{
	Name:   "IberianGoatSkull",
	MaxHP:  avgHP,
	ATK:    avgATK,
	Weight: avgWeight,
	OnHit:  defaultOnHit,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		spiritTypes[0]: midloDEF,
		spiritTypes[1]: highDEF,
		spiritTypes[2]: highDEF,
		spiritTypes[3]: lowDEF,
		switchType:     switchDEF,
	},
	Moves: []*types.Move{},
}

var AttisCrown types.EquipmentTemplate = types.EquipmentTemplate{
	Name:   "AttisCrown",
	MaxHP:  avgHP,
	ATK:    avgATK,
	Weight: avgWeight,
	OnHit:  defaultOnHit,
	OnMiss: noop,
	OnDbl:  noop,
	Defs: map[string]int{
		spiritTypes[0]: avgDEF,
		spiritTypes[1]: lowDEF,
		spiritTypes[2]: avgDEF,
		spiritTypes[3]: highDEF,
		switchType:     switchDEF,
	},
	Moves: []*types.Move{},
}
