package types

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Equipment struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	HP            int            `json:"hp"`
	MaxHP         int            `json:"maxHP"`
	ATK           int            `json:"atk"`
	Defs          map[string]int `json:"defenses"`
	Weight        int            `json:"weight"`
	Moves         []*Move        `json:"moves"`
	Inhabited     bool           `json:"inhabited"`
	InhabitedBy   *Spirit        `json:"inhabitedBy"`
	InhabitedById string         `json:"inhabitedById"`
	onHit         Callback
	onMiss        Callback
	onDbl         Callback
}

type EquipmentTemplate struct {
	Name   string         `json:"name"`
	MaxHP  int            `json:"maxHP"`
	ATK    int            `json:"atk"`
	Defs   map[string]int `json:"defenses"`
	Weight int            `json:"weight"`
	Moves  []*Move        `json:"moves"`
	OnHit  Callback       `json:"-"`
	OnMiss Callback       `json:"-"`
	OnDbl  Callback       `json:"-"`
}

func (e *Equipment) GetID() string {
	return e.ID
}

func (e *Equipment) GetName() string {
	return e.Name
}

func (e *Equipment) GetDef(dmgType string) int {
	return e.Defs[dmgType]
}

func (e *Equipment) TakeDamage(dmg int) {
	e.HP -= dmg
}

func (e *Equipment) OnHit(user *Spirit, target *Equipment, move *Move) {
	e.onHit(user, target, move)
}

func (e *Equipment) OnMiss(user *Spirit, target *Equipment, move *Move) {
	e.onMiss(user, target, move)
}

func (e *Equipment) OnDbl(user *Spirit, target *Equipment, move *Move) {
	e.onDbl(user, target, move)
}

func (e *Equipment) MarshalJSON() ([]byte, error) {
	inhabitedById := ""
	if e.InhabitedBy != nil {
		inhabitedById = e.InhabitedBy.ID
	}

	// fmt.Printf("%+v\n", e)
	// return json.Marshal("asdf")

	return json.Marshal(&struct {
		ID            string         `json:"id"`
		Name          string         `json:"name"`
		HP            int            `json:"hp"`
		MaxHP         int            `json:"maxHP"`
		ATK           int            `json:"atk"`
		Defs          map[string]int `json:"defenses"`
		Weight        int            `json:"weight"`
		Moves         []*Move        `json:"moves"`
		Inhabited     bool           `json:"inhabited"`
		InhabitedById string         `json:"inhabitedBy"`
	}{
		ID:            e.ID,
		Name:          e.Name,
		HP:            e.HP,
		MaxHP:         e.MaxHP,
		ATK:           e.ATK,
		Defs:          e.Defs,
		Weight:        e.Weight,
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

	e.onHit = et.OnHit
	e.onMiss = et.OnMiss
	e.onDbl = et.OnDbl

	id, _ := uuid.NewRandom()
	e.ID = id.String()

	for _, move := range et.Moves {
		e.Moves = append(e.Moves, move)
	}

	return e
}
