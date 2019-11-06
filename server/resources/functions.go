package resources

import (
	"github.com/aopal/mytholojam/server/types"
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

var defaultLowerDefenses types.Callback = func(_ *types.Spirit, target types.Damageable, _ *types.Move, _ int) {
	statMod := &types.StatMod{
		DefMods: map[string]int{
			spiritTypes[0]: -1,
			spiritTypes[1]: -1,
			spiritTypes[2]: -1,
			spiritTypes[3]: -1,
		},
	}
	target.ApplyStatMod(statMod)
}

var defaultLowerAtk types.Callback = func(_ *types.Spirit, target types.Damageable, _ *types.Move, _ int) {
	statMod := &types.StatMod{
		AtkMod: -1,
	}
	target.ApplyStatMod(statMod)
}

var defaultLowerSpeed types.Callback = func(_ *types.Spirit, target types.Damageable, _ *types.Move, _ int) {
	statMod := &types.StatMod{
		SpeedMod:  -1,
		WeightMod: -1,
	}
	target.ApplyStatMod(statMod)
}
