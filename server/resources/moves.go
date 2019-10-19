package resources

import "mytholojam/server/types"

// SAMPLE MOVES
var Switch types.Move = types.Move{
	Name:           "Switch",
	Power:          switchPower,
	Type:           switchType,
	Priority:       switchPri,
	MultiTarget:    false,
	TeamTargetable: selfTarget,
}

var Strong types.Move = types.Move{
	Name:           "Strong",
	Power:          highPWR,
	Type:           moveTypes[0],
	Priority:       avgPri,
	MultiTarget:    false,
	TeamTargetable: opTarget,
}

var Weak types.Move = types.Move{
	Name:           "Weak",
	Power:          lowPWR,
	Type:           moveTypes[1],
	Priority:       avgPri,
	MultiTarget:    true,
	TeamTargetable: opTarget,
}

var Fast types.Move = types.Move{
	Name:           "Fast",
	Power:          avgPWR,
	Type:           moveTypes[2],
	Priority:       highPri,
	MultiTarget:    false,
	TeamTargetable: opTarget,
}

var MoveList map[string]*types.Move = map[string]*types.Move{
	Switch.Name: &Switch,
	Strong.Name: &Strong,
	Weak.Name:   &Weak,
	Fast.Name:   &Fast,
}
