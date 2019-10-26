package types

import (
	"errors"

	"github.com/google/uuid"
)

type Player struct {
	Equipment   map[string]*Equipment `json:"equipment"`
	Spirits     map[string]*Spirit    `json:"spirits"`
	ID          string                `json:"-"`
	Opponent    *Player               `json:"-"`
	NextActions []*Action             `json:"-"`
}

func NewPlayer() (*Player, error) {
	playerToken, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("Could not create player.\n")
	}

	p := Player{
		Equipment:   make(map[string]*Equipment),
		Spirits:     make(map[string]*Spirit),
		ID:          playerToken.String(),
		NextActions: nil,
	}

	return &p, nil
}
