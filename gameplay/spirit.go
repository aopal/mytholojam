package gameplay

import (
	"encoding/json"

	"github.com/google/uuid"
)

type spirit struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	HP         int            `json:"hp"`
	MaxHP      int            `json:"maxHP"`
	ATK        int            `json:"atk"`
	Defs       map[string]int `json:"defenses"`
	Speed      int            `json:"speed"`
	Moves      []*move        `json:"moves"`
	Inhabiting *equipment     `json:"inhabiting"`
}

type spiritTemplate struct {
	Name  string         `json:"name"`
	MaxHP int            `json:"maxHP"`
	ATK   int            `json:"atk"`
	Defs  map[string]int `json:"defenses"`
	Speed int            `json:"speed"`
	Moves []string       `json:"moves"`
}

func (s *spirit) getID() string {
	return s.ID
}

func (s *spirit) getName() string {
	return s.Name
}

func (s *spirit) getDef(dmgType string) int {
	return s.Defs[dmgType]
}

func (s *spirit) takeDamage(dmg int) {
	s.HP -= dmg
}

func (s *spirit) Inhabit(e *equipment) bool {
	if e.Inhabited {
		return false
	}

	if s.Inhabiting != nil {
		s.Inhabiting.Inhabited = false
		s.Inhabiting.InhabitedBy = nil
	}

	e.Inhabited = true
	e.InhabitedBy = s
	s.Inhabiting = e

	return true
}

func (s *spirit) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID         string         `json:"id"`
		HP         int            `json:"hp"`
		Name       string         `json:"name"`
		MaxHP      int            `json:"maxHP"`
		ATK        int            `json:"atk"`
		Defs       map[string]int `json:"defenses"`
		Speed      int            `json:"Speed"`
		Moves      []*move        `json:"moves"`
		Inhabiting string         `json:"inhabiting"`
	}{
		ID:         s.ID,
		HP:         s.HP,
		Name:       s.Name,
		MaxHP:      s.MaxHP,
		ATK:        s.ATK,
		Defs:       s.Defs,
		Speed:      s.Speed,
		Moves:      s.Moves,
		Inhabiting: s.Inhabiting.ID,
	})
}

func (st *spiritTemplate) NewSpirit() *spirit {
	s := new(spirit)

	s.Name = st.Name
	s.MaxHP = st.MaxHP
	s.HP = st.MaxHP
	s.ATK = st.ATK
	s.Defs = st.Defs
	s.Speed = st.Speed

	id, _ := uuid.NewRandom()
	s.ID = id.String()

	for _, move := range st.Moves {
		s.Moves = append(s.Moves, moveList[move])
	}

	return s
}
