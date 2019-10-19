package types

type Player struct {
	Equipment   map[string]*Equipment `json:"equipment"`
	Spirits     map[string]*Spirit    `json:"spirits"`
	ID          string                `json:"-"`
	Opponent    *Player               `json:"-"`
	NextActions []*Action             `json:"-"`
}
