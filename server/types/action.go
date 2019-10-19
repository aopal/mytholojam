package types

type Action struct {
	User    *Spirit      `json:"user"`
	Targets []*Equipment `json:"targets"`
	Move    *Move        `json:"move"`
}

type ActionPayload struct {
	Token   string    `json:"token"`
	Actions []*Action `json:"actions"`
}
