package resources

import (
	"mytholojam/server/types"
)

var empty types.CallbackArray = types.CallbackArray{}
var defaultOnHitArr types.CallbackArray = types.CallbackArray{defaultOnHit}

var switchOnHit types.Callback = func(user *types.Spirit, target types.Damageable, _ *types.Move, _ int) {
	user.Inhabit(target)
}

var defaultOnHit types.Callback = func(_ *types.Spirit, target types.Damageable, _ *types.Move, damage int) {
	target.TakeDamage(damage)
}

var defaultRecoil types.Callback = func(user *types.Spirit, _ types.Damageable, _ *types.Move, _ int) {
	user.TakeDamage(RecoilDamage)
}
