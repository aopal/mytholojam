package types

type Damageable interface {
	GetID() string
	GetName() string
	GetDef(string) int
	TakeDamage(int)
	OnHit(user *Spirit, target *Equipment, move *Move)
	OnMiss(user *Spirit, target *Equipment, move *Move)
	OnDbl(user *Spirit, target *Equipment, move *Move)
}
