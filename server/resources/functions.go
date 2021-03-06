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
	statMod := types.NewStatMod()
	statMod.DefMods[spiritTypes[0]] = -1
	statMod.DefMods[spiritTypes[1]] = -1
	statMod.DefMods[spiritTypes[2]] = -1
	statMod.DefMods[spiritTypes[3]] = -1
	target.ApplyStatMod(statMod)
}

var defaultLowerAtk types.Callback = func(_ *types.Spirit, target types.Damageable, _ *types.Move, _ int) {
	statMod := types.NewStatMod()
	statMod.AtkMod = -1
	target.ApplyStatMod(statMod)
}

var defaultLowerSpeed types.Callback = func(_ *types.Spirit, target types.Damageable, _ *types.Move, _ int) {
	statMod := types.NewStatMod()
	statMod.SpeedMod = -1
	statMod.WeightMod = -1
	target.ApplyStatMod(statMod)
}
