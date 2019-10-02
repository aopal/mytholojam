package gameplay

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
}

type equipmentTemplate struct {
	Name   string         `json:"name"`
	MaxHP  int            `json:"maxHP"`
	ATK    int            `json:"atk"`
	Defs   map[string]int `json:"defenses"`
	Weight int            `json:"weight"`
	Moves  []string       `json:"moves"`
}

func (e *Equipment) getID() string {
	return e.ID
}

func (e *Equipment) getName() string {
	return e.Name
}

func (e *Equipment) getDef(dmgType string) int {
	return e.Defs[dmgType]
}

func (e *Equipment) takeDamage(dmg int) {
	e.HP -= dmg
}

func (e *Equipment) MarshalJSON() ([]byte, error) {
	inhabitedById := ""
	if e.InhabitedBy != nil {
		inhabitedById = e.InhabitedBy.ID
	}

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

func (et *equipmentTemplate) NewEquipment() *Equipment {
	e := new(Equipment)
	e.Name = et.Name
	e.MaxHP = et.MaxHP
	e.HP = et.MaxHP
	e.ATK = et.ATK
	e.Defs = et.Defs
	e.Weight = et.Weight

	id, _ := uuid.NewRandom()
	e.ID = id.String()

	for _, move := range et.Moves {
		e.Moves = append(e.Moves, moveList[move])
	}

	return e
}
