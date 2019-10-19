package resources

import "mytholojam/server/types"

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
	OnHit:          switchOnHit,
	OnMiss:         noop,
	OnDbl:          noop,
}

var Strong types.Move = types.Move{
	Name:           "Strong",
	Power:          highPWR,
	Type:           spiritTypes[0],
	Priority:       avgPri,
	MultiTarget:    false,
	TeamTargetable: opTarget,
	OnHit:          noop,
	OnMiss:         noop,
	OnDbl:          noop,
}

var Weak types.Move = types.Move{
	Name:           "Weak",
	Power:          lowPWR,
	Type:           spiritTypes[1],
	Priority:       avgPri,
	MultiTarget:    true,
	TeamTargetable: opTarget,
	OnHit:          noop,
	OnMiss:         noop,
	OnDbl:          noop,
}

var Fast types.Move = types.Move{
	Name:           "Fast",
	Power:          avgPWR,
	Type:           spiritTypes[2],
	Priority:       highPri,
	MultiTarget:    false,
	TeamTargetable: opTarget,
	OnHit:          noop,
	OnMiss:         noop,
	OnDbl:          noop,
}
