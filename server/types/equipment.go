package types

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Equipment struct {
	ID            string           `json:"id"`
	Name          string           `json:"name"`
	HP            int              `json:"hp"`
	MaxHP         int              `json:"maxHP"`
	ATK           int              `json:"atk"`
	AtkMod        int              `json:"atkMod"`
	Defs          map[string]int   `json:"defenses"`
	DefMods       map[string]int   `json:"defMods"`
	Weight        int              `json:"weight"`
	WeightMod     int              `json:"weightMod"`
	Moves         map[string]*Move `json:"moves"`
	Inhabited     bool             `json:"inhabited"`
	InhabitedBy   *Spirit          `json:"inhabitedBy"`
	InhabitedById string           `json:"inhabitedById"`
	onHit         CallbackArray
	onMiss        CallbackArray
	onDbl         CallbackArray
}

type EquipmentTemplate struct {
	Name   string           `json:"name"`
	MaxHP  int              `json:"maxHP"`
	ATK    int              `json:"atk"`
	Defs   map[string]int   `json:"defenses"`
	Weight int              `json:"weight"`
	Moves  map[string]*Move `json:"moves"`
	OnHit  CallbackArray    `json:"-"`
	OnMiss CallbackArray    `json:"-"`
	OnDbl  CallbackArray    `json:"-"`
}

func (e *Equipment) GetID() string             { return e.ID }
func (e *Equipment) GetName() string           { return e.Name }
func (e *Equipment) GetDef(dmgType string) int { return e.Defs[dmgType] + e.DefMods[dmgType] }
func (e *Equipment) GetAtk() int               { return e.ATK + e.AtkMod }
func (e *Equipment) GetWeight() int            { return e.Weight + e.WeightMod }
func (e *Equipment) GetHP() int                { return e.HP }
func (e *Equipment) TakeDamage(dmg int)        { e.HP -= dmg }
func (e *Equipment) GetEquipment() *Equipment  { return e }

func (e *Equipment) OnHit(user *Spirit, target Damageable, move *Move, damage int) {
	for _, f := range e.onHit {
		f(user, target, move, damage)
	}
}

func (e *Equipment) OnMiss(user *Spirit, target Damageable, move *Move, damage int) {
	for _, f := range e.onMiss {
		f(user, target, move, damage)
	}
}

func (e *Equipment) OnDbl(user *Spirit, target Damageable, move *Move, damage int) {
	for _, f := range e.onDbl {
		f(user, target, move, damage)
	}
}

func (e *Equipment) MarshalJSON() ([]byte, error) {
	inhabitedById := ""
	if e.InhabitedBy != nil {
		inhabitedById = e.InhabitedBy.ID
	}

	return json.Marshal(&struct {
		ID            string           `json:"id"`
		Name          string           `json:"name"`
		HP            int              `json:"hp"`
		MaxHP         int              `json:"maxHP"`
		ATK           int              `json:"atk"`
		AtkMod        int              `json:"atkMod"`
		Defs          map[string]int   `json:"defenses"`
		DefMods       map[string]int   `json:"defMods"`
		Weight        int              `json:"weight"`
		WeightMod     int              `json:"weightMod"`
		Moves         map[string]*Move `json:"moves"`
		Inhabited     bool             `json:"inhabited"`
		InhabitedById string           `json:"inhabitedBy"`
	}{
		ID:            e.ID,
		Name:          e.Name,
		HP:            e.HP,
		MaxHP:         e.MaxHP,
		ATK:           e.ATK,
		AtkMod:        e.AtkMod,
		Defs:          e.Defs,
		DefMods:       e.DefMods,
		Weight:        e.Weight,
		WeightMod:     e.WeightMod,
		Moves:         e.Moves,
		Inhabited:     e.Inhabited,
		InhabitedById: inhabitedById,
	})
}

func (et *EquipmentTemplate) NewEquipment() *Equipment {
	e := new(Equipment)
	e.Name = et.Name
	e.MaxHP = et.MaxHP
	e.HP = et.MaxHP
	e.ATK = et.ATK
	e.Defs = et.Defs
	e.Weight = et.Weight

	e.DefMods = make(map[string]int)

	e.onHit = et.OnHit
	e.onMiss = et.OnMiss
	e.onDbl = et.OnDbl

	id, _ := uuid.NewRandom()
	e.ID = id.String()

	e.Moves = make(map[string]*Move)
	for key, move := range et.Moves {
		e.Moves[key] = move
	}

	return e
}
