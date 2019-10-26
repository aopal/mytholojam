package types

type Move struct { // the swap move is unique
	Name           string        `json:"name"`
	Power          int           `json:"power"`
	Type           string        `json:"type"`
	Priority       int           `json:"priority"`
	MultiTarget    bool          `json:"multiTarget"`
	TeamTargetable string        `json:"teamTargetable"`
	OnHitFuncs     CallbackArray `json:"-"`
	OnMissFuncs    CallbackArray `json:"-"`
	OnDblFuncs     CallbackArray `json:"-"`
}

func (m *Move) OnHit(user *Spirit, target Damageable, move *Move, damage int) {
	for _, f := range m.OnHitFuncs {
		f(user, target, move, damage)
	}
}

func (m *Move) OnMiss(user *Spirit, target Damageable, move *Move, damage int) {
	for _, f := range m.OnMissFuncs {
		f(user, target, move, damage)
	}
}

func (m *Move) OnDbl(user *Spirit, target Damageable, move *Move, damage int) {
	for _, f := range m.OnDblFuncs {
		f(user, target, move, damage)
	}
}
