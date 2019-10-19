package types

type Damageable interface {
	GetID() string
	GetName() string
	GetDef(string) int
	GetHP() int
	TakeDamage(int)
	GetEquipment() *Equipment
	OnHit(user *Spirit, target Damageable, move *Move, damage int)
	OnMiss(user *Spirit, target Damageable, move *Move, damage int)
	OnDbl(user *Spirit, target Damageable, move *Move, damage int)
}
