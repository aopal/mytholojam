package types

type Action struct {
	User    *Spirit      `json:"user"`
	Targets []*Equipment `json:"targets"`
	Move    *Move        `json:"move"` // name of attacking move, or the special 'swap' move
}

type ActionPayload struct {
	Token   string    `json:"token"`
	Actions []*Action `json:"actions"`
}
