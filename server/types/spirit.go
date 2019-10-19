package types

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Spirit struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	HP           int            `json:"hp"`
	MaxHP        int            `json:"maxHP"`
	ATK          int            `json:"atk"`
	Defs         map[string]int `json:"defenses"`
	Speed        int            `json:"speed"`
	Moves        []*Move        `json:"moves"`
	Inhabiting   *Equipment     `json:"inhabiting"`
	InhabitingId string         `json:"inhabitingId"`
	onHit        Callback
	onMiss       Callback
	onDbl        Callback
}

type SpiritTemplate struct {
	Name   string         `json:"name"`
	MaxHP  int            `json:"maxHP"`
	ATK    int            `json:"atk"`
	Defs   map[string]int `json:"defenses"`
	Speed  int            `json:"speed"`
	Moves  []*Move        `json:"moves"`
	OnHit  Callback       `json:"-"`
	OnMiss Callback       `json:"-"`
	OnDbl  Callback       `json:"-"`
}

func (s *Spirit) GetID() string {
	return s.ID
}

func (s *Spirit) GetName() string {
	return s.Name
}

func (s *Spirit) GetDef(dmgType string) int {
	return s.Defs[dmgType]
}

func (s *Spirit) TakeDamage(dmg int) {
	s.HP -= dmg
}

func (s *Spirit) OnHit(user *Spirit, target *Equipment, move *Move) {
	s.onHit(user, target, move)
}

func (s *Spirit) OnMiss(user *Spirit, target *Equipment, move *Move) {
	s.onMiss(user, target, move)
}

func (s *Spirit) OnDbl(user *Spirit, target *Equipment, move *Move) {
	s.onDbl(user, target, move)
}

func (s *Spirit) Inhabit(e *Equipment) bool {
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

func (s *Spirit) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID           string         `json:"id"`
		HP           int            `json:"hp"`
		Name         string         `json:"name"`
		MaxHP        int            `json:"maxHP"`
		ATK          int            `json:"atk"`
		Defs         map[string]int `json:"defenses"`
		Speed        int            `json:"Speed"`
		Moves        []*Move        `json:"moves"`
		Inhabiting   string         `json:"inhabiting"`
		InhabitingId string         `json:"inhabitingId"`
	}{
		ID:           s.ID,
		HP:           s.HP,
		Name:         s.Name,
		MaxHP:        s.MaxHP,
		ATK:          s.ATK,
		Defs:         s.Defs,
		Speed:        s.Speed,
		Moves:        s.Moves,
		InhabitingId: s.Inhabiting.ID,
	})
}

func (st *SpiritTemplate) NewSpirit() *Spirit {
	s := new(Spirit)

	s.Name = st.Name
	s.MaxHP = st.MaxHP
	s.HP = st.MaxHP
	s.ATK = st.ATK
	s.Defs = st.Defs
	s.Speed = st.Speed

	s.onHit = st.OnHit
	s.onMiss = st.OnMiss
	s.onDbl = st.OnDbl

	id, _ := uuid.NewRandom()
	s.ID = id.String()

	for _, move := range st.Moves {
		s.Moves = append(s.Moves, move)
	}

	return s
}
