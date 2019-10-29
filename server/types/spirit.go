package types

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Spirit struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	HP           int              `json:"hp"`
	MaxHP        int              `json:"maxHP"`
	ATK          int              `json:"atk"`
	AtkMod       int              `json:"atkMod"`
	Defs         map[string]int   `json:"defenses"`
	DefMods      map[string]int   `json:"defMods"`
	Speed        int              `json:"speed"`
	SpeedMod     int              `json:"speedMod"`
	Moves        map[string]*Move `json:"moves"`
	Inhabiting   *Equipment       `json:"inhabiting"`
	InhabitingId string           `json:"inhabitingId"`
	onHit        CallbackArray
	onMiss       CallbackArray
	onDbl        CallbackArray
}

type SpiritTemplate struct {
	Name   string           `json:"name"`
	MaxHP  int              `json:"maxHP"`
	ATK    int              `json:"atk"`
	Defs   map[string]int   `json:"defenses"`
	Speed  int              `json:"speed"`
	Moves  map[string]*Move `json:"moves"`
	OnHit  CallbackArray    `json:"-"`
	OnMiss CallbackArray    `json:"-"`
	OnDbl  CallbackArray    `json:"-"`
}

func (s *Spirit) GetID() string             { return s.ID }
func (s *Spirit) GetName() string           { return s.Name }
func (s *Spirit) GetDef(dmgType string) int { return s.Defs[dmgType] + s.DefMods[dmgType] }
func (s *Spirit) GetAtk() int               { return s.ATK + s.AtkMod }
func (s *Spirit) GetSpeed() int             { return s.Speed + s.SpeedMod }
func (s *Spirit) GetHP() int                { return s.HP }
func (s *Spirit) TakeDamage(dmg int)        { s.HP -= dmg }
func (s *Spirit) GetEquipment() *Equipment  { return s.Inhabiting }

func (s *Spirit) GetMove(moveName string) *Move {
	if m, ok := s.Moves[moveName]; ok {
		return m
	} else if m, ok := s.Inhabiting.Moves[moveName]; ok {
		return m
	}

	return nil
}

func (s *Spirit) OnHit(user *Spirit, target Damageable, move *Move, damage int) {
	for _, f := range s.onHit {
		f(user, target, move, damage)
	}
}

func (s *Spirit) OnMiss(user *Spirit, target Damageable, move *Move, damage int) {
	for _, f := range s.onMiss {
		f(user, target, move, damage)
	}
}

func (s *Spirit) OnDbl(user *Spirit, target Damageable, move *Move, damage int) {
	for _, f := range s.onDbl {
		f(user, target, move, damage)
	}
}

func (s *Spirit) Inhabit(t Damageable) bool {
	e := t.GetEquipment()

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
		ID            string           `json:"id"`
		HP            int              `json:"hp"`
		Name          string           `json:"name"`
		MaxHP         int              `json:"maxHP"`
		ATK           int              `json:"atk"`
		AtkMod        int              `json:"atkMod"`
		Defs          map[string]int   `json:"defenses"`
		DefMods       map[string]int   `json:"defMods"`
		Speed         int              `json:"speed"`
		SpeedMod      int              `json:"speedMod"`
		StatModifiers map[string]int   `json:"statMods"`
		Moves         map[string]*Move `json:"moves"`
		Inhabiting    string           `json:"inhabiting"`
		InhabitingId  string           `json:"inhabitingId"`
	}{
		ID:           s.ID,
		HP:           s.HP,
		Name:         s.Name,
		MaxHP:        s.MaxHP,
		ATK:          s.ATK,
		AtkMod:       s.AtkMod,
		Defs:         s.Defs,
		DefMods:      s.DefMods,
		Speed:        s.Speed,
		SpeedMod:     s.SpeedMod,
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

	s.DefMods = make(map[string]int)

	s.onHit = st.OnHit
	s.onMiss = st.OnMiss
	s.onDbl = st.OnDbl

	id, _ := uuid.NewRandom()
	s.ID = id.String()

	s.Moves = make(map[string]*Move)
	for key, move := range st.Moves {
		s.Moves[key] = move
	}

	return s
}
