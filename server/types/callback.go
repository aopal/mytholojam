package types

type Callback func(user *Spirit, target Damageable, move *Move, damage int)

type CallbackArray []Callback
