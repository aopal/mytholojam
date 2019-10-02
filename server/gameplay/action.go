package gameplay

type action struct {
	User    *spirit      `json:"user"`
	Targets []*equipment `json:"targets"`
	Move    *move        `json:"move"` // name of attacking move, or the special 'swap' move
}

type actionPayload struct {
	Token   string    `json:"token"`
	Actions []*action `json:"actions"`
}

type damageable interface {
	getID() string
	getName() string
	getDef(string) int
	takeDamage(int)
}
