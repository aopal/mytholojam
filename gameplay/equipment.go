package gameplay

import (
	"encoding/json"

	"github.com/google/uuid"
)

type equipment struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	HP          int     `json:"hp"`
	MaxHP       int     `json:"maxHP"`
	ATK         int     `json:"atk"`
	DEF         int     `json:"def"`
	Moves       []*move `json:"moves"`
	Inhabited   bool    `json:"inhabited"`
	InhabitedBy *spirit `json:"inhabitedBy"`
}

type equipmentTemplate struct {
	Name  string   `json:"name"`
	MaxHP int      `json:"maxHP"`
	ATK   int      `json:"atk"`
	DEF   int      `json:"def"`
	Moves []string `json:"moves"`
}

func (e *equipment) getID() string {
	return e.ID
}

func (e *equipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		HP        int     `json:"hp"`
		MaxHP     int     `json:"maxHP"`
		ATK       int     `json:"atk"`
		DEF       int     `json:"def"`
		Moves     []*move `json:"moves"`
		Inhabited bool    `json:"inhabited"`
		// InhabitedBy string  `json:"inhabitedBy"`
	}{
		ID:        e.ID,
		Name:      e.Name,
		HP:        e.HP,
		MaxHP:     e.MaxHP,
		ATK:       e.ATK,
		DEF:       e.DEF,
		Moves:     e.Moves,
		Inhabited: e.Inhabited,
		// InhabitedBy: e.InhabitedBy.ID,
	})
}

func (et *equipmentTemplate) NewEquipment() *equipment {
	e := new(equipment)
	e.Name = et.Name
	e.MaxHP = et.MaxHP
	e.HP = et.MaxHP
	e.ATK = et.ATK
	e.DEF = et.DEF

	id, _ := uuid.NewRandom()
	e.ID = id.String()

	for _, move := range et.Moves {
		e.Moves = append(e.Moves, moveList[move])
	}

	return e
}
