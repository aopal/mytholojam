package resources

import "github.com/aopal/mytholojam/server/types"

var MoveList map[string]*types.Move = map[string]*types.Move{
	Switch.Name: &Switch,
	Strong.Name: &Strong,
	Weak.Name:   &Weak,
	Fast.Name:   &Fast,
}

// SAMPLE MOVES
var Switch types.Move = types.Move{
	Name:           "Switch",
	Power:          switchPower,
	Type:           switchType,
	Priority:       switchPri,
	MultiTarget:    false,
	TeamTargetable: selfTarget,
	OnHitFuncs: types.CallbackArray{
		switchOnHit,
	},
	OnMissFuncs: empty,
	OnDblFuncs:  empty,
}

var Strong types.Move = types.Move{
	Name:           "Strong",
	Power:          highPWR,
	Type:           spiritTypes[0],
	Priority:       avgPri,
	MultiTarget:    false,
	TeamTargetable: opTarget,
	OnHitFuncs: types.CallbackArray{
		defaultOnHit,
		defaultRecoil,
	},
	OnMissFuncs: empty,
	OnDblFuncs:  empty,
}

var Weak types.Move = types.Move{
	Name:           "Weak",
	Power:          lowPWR,
	Type:           spiritTypes[1],
	Priority:       avgPri,
	MultiTarget:    true,
	TeamTargetable: opTarget,
	OnHitFuncs:     defaultOnHitArr,
	OnMissFuncs:    empty,
	OnDblFuncs:     empty,
}

var Fast types.Move = types.Move{
	Name:           "Fast",
	Power:          avgPWR,
	Type:           spiritTypes[2],
	Priority:       highPri,
	MultiTarget:    false,
	TeamTargetable: opTarget,
	OnHitFuncs:     defaultOnHitArr,
	OnMissFuncs:    empty,
	OnDblFuncs:     empty,
}

var StatChange types.Move = types.Move{
	Name:           "Lower Defs",
	Power:          avgPWR,
	Type:           spiritTypes[2],
	Priority:       avgPri,
	MultiTarget:    false,
	TeamTargetable: opTarget,
	OnHitFuncs: types.CallbackArray{
		defaultLowerDefenses,
	},
	OnMissFuncs: empty,
	OnDblFuncs: types.CallbackArray{
		defaultLowerDefenses,
		defaultLowerDefenses,
	},
}

// DESIGN FIRST PASS MOVES

var SpearThrust types.Move = types.Move{
	Name:           "Thrust",
	Power:          avgPWR,
	Type:           spiritTypes[0],
	Priority:       avgPri,
	MultiTarget:    false,
	TeamTargetable: opTarget,
	OnHitFuncs: types.CallbackArray{
		doDamage,
	},
	OnMissFuncs: types.CallbackArray{
		defaultRecoilEquipment,
	},
	OnDblFuncs: types.CallbackArray{
		doDamage, // TODO: PIERCE to attack equipment inhabitated/target behind
	},
}

var RodCrush types.Move = types.Move{
	Name:           "Crush",
	Power:          highPWR,
	Type:           spiritTypes[0],
	Priority:       lowPri,
	MultiTarget:    false,
	TeamTargetable: opTarget,
	OnHitFuncs: types.CallbackArray{
		lowerSelfEquipmentStrn,
		lowerSelfEquipmentWear,
	},
	OnMissFuncs: types.CallbackArray{
		lowerName,
		lowerFlam,
	},
	OnDblFuncs: types.CallbackArray{
		doDamage,
		lowerName,
		lowerFlam,
	},
}

var Cremate types.Move = types.Move{
	Name:           "Cremate",
	Power:          midloPWR,
	Type:           spiritTypes[1], // TODO: Ability to target min or max of 2 defenses
	Priority:       avgPri,
	MultiTarget:    false,
	TeamTargetable: opTarget,
	OnHitFuncs: types.CallbackArray{ // TODO: Recoil not on a "miss" vs. target but as a general-case/targetless callback after use
		defaultRecoilEquipment,
		lowerWear,
		lowerName,
	},
	OnMissFuncs: types.CallbackArray{
		defaultRecoilEquipment,
	},
	OnDblFuncs: types.CallbackArray{
		defaultRecoilEquipment,
		doDamage,
		lowerWear,
		lowerName,
	},
}

var Needle types.Move = types.Move{
	Name:           "Needle",
	Power:          lowPWR, // TODO: Power multiplicative by squares of target hit
	Type:           spiritTypes[2],
	Priority:       avgPri,
	MultiTarget:    false,
	TeamTargetable: opTarget,
	OnHitFuncs: types.CallbackArray{
		doDamage,
	},
	OnMissFuncs: empty,
	OnDblFuncs: types.CallbackArray{
		doDamage,
	},
}

var Incinerate types.Move = types.Move{
	Name:           "Incinerate",
	Power:          highPWR,
	Type:           spiritTypes[1],
	Priority:       avgPri,
	MultiTarget:    false,
	TeamTargetable: opTarget,
	OnHitFuncs: types.CallbackArray{
		defaultRecoilEquipment,
		lowerFlam,
	},
	OnMissFuncs: types.CallbackArray{
		defaultRecoil,
		defaultRecoilEquipment,
	},
	OnDblFuncs: types.CallbackArray{
		doDamage,
		doDamage,
	},
}
