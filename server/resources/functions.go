package resources

import (
	"mytholojam/server/types"
)

var noop types.Callback = func(_ *types.Spirit, _ types.Damageable, _ *types.Move, _ int) {}

var switchOnHit types.Callback = func(user *types.Spirit, target types.Damageable, move *types.Move, damage int) {
	user.Inhabit(target)
}

// default for equip/spirit, not moves
var defaultOnHit types.Callback = func(user *types.Spirit, target types.Damageable, move *types.Move, damage int) {
	target.TakeDamage(damage)
}
