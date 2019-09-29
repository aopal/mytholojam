package gameplay

import (
	"encoding/json"

	"github.com/google/uuid"
)

type equipment struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	HP          int            `json:"hp"`
	MaxHP       int            `json:"maxHP"`
	ATK         int            `json:"atk"`
	Defs        map[string]int `json:"defenses"`
	Weight      int            `json:"weight"`
	Moves       []*move        `json:"moves"`
	Inhabited   bool           `json:"inhabited"`
	InhabitedBy *spirit        `json:"inhabitedBy"`
}

type equipmentTemplate struct {
	Name   string         `json:"name"`
	MaxHP  int            `json:"maxHP"`
	ATK    int            `json:"atk"`
	Defs   map[string]int `json:"defenses"`
	Weight int            `json:"weight"`
	Moves  []string       `json:"moves"`
}

func (e *equipment) getID() string {
	return e.ID
}

func (e *equipment) getName() string {
	return e.Name
}

func (e *equipment) getDef(dmgType string) int {
	return e.Defs[dmgType]
}

func (e *equipment) takeDamage(dmg int) {
	e.HP -= dmg
}

func (e *equipment) MarshalJSON() ([]byte, error) {
	inhabitedBy := ""
	if e.InhabitedBy != nil {
		inhabitedBy = e.InhabitedBy.ID
	}

	return json.Marshal(&struct {
		ID          string         `json:"id"`
		Name        string         `json:"name"`
		HP          int            `json:"hp"`
		MaxHP       int            `json:"maxHP"`
		ATK         int            `json:"atk"`
		Defs        map[string]int `json:"defenses"`
		Moves       []*move        `json:"moves"`
		Inhabited   bool           `json:"inhabited"`
		InhabitedBy string         `json:"inhabitedBy"`
	}{
		ID:          e.ID,
		Name:        e.Name,
		HP:          e.HP,
		MaxHP:       e.MaxHP,
		ATK:         e.ATK,
		Defs:        e.Defs,
		Moves:       e.Moves,
		Inhabited:   e.Inhabited,
		InhabitedBy: inhabitedBy,
	})
}

func (et *equipmentTemplate) NewEquipment() *equipment {
	e := new(equipment)
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
