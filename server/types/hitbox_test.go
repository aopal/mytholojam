package types

import (
	"testing"

	"gotest.tools/assert"
)

func TestDraw(t *testing.T) {
	g := Grid{Width: 50, Height: 12, Entities: make(map[string]*Hitbox)}

	g.Entities["hitbox1"] = &Hitbox{
		ID:       "hitbox1",
		Height:   len(spearShape),
		Width:    len(spearShape[0]),
		Zindex:   0,
		Shape:    spearShape,
		Position: coordinate{y: 2, x: 3},
	}

	g.Entities["hitbox2"] = &Hitbox{
		ID:       "hitbox2",
		Height:   len(needleAttackShape),
		Width:    len(needleAttackShape[0]),
		Zindex:   0,
		Shape:    needleAttackShape,
		Position: coordinate{y: 6, x: 10},
	}

	g.Entities["hitbox3"] = &Hitbox{
		ID:       "hitbox3",
		Height:   len(spearShape),
		Width:    len(spearShape[0]),
		Zindex:   0,
		Shape:    spearShape,
		Position: coordinate{y: 0, x: 25},
	}

	g.Entities["hitbox4"] = &Hitbox{
		ID:       "hitbox4",
		Height:   len(needleAttackShape),
		Width:    len(needleAttackShape[0]),
		Zindex:   0,
		Shape:    needleAttackShape,
		Position: coordinate{y: 1, x: 40},
	}

	g.Draw()
}

func TestCollistion(t *testing.T) {
	g := Grid{Width: 20, Height: 12, Entities: make(map[string]*Hitbox)}

	g.Entities["hitbox1"] = &Hitbox{
		ID:       "hitbox1",
		Height:   len(spearShape),
		Width:    len(spearShape[0]),
		Zindex:   0,
		Shape:    spearShape,
		Position: coordinate{y: 2, x: 3},
		Object:   &Equipment{},
	}

	g.Entities["hitbox2"] = &Hitbox{
		ID:       "hitbox2",
		Height:   len(spearShape),
		Width:    len(spearShape[0]),
		Zindex:   0,
		Shape:    spearShape,
		Position: coordinate{y: 2, x: 5},
		Object:   &Equipment{},
	}

	g.Entities["hitbox3"] = &Hitbox{
		ID:       "hitbox3",
		Height:   len(spearShape),
		Width:    len(spearShape[0]),
		Zindex:   0,
		Shape:    spearShape,
		Position: coordinate{y: 2, x: 7},
		Object:   &Equipment{},
	}

	g.Draw()

	a := Hitbox{
		ID:       "attacker",
		Height:   len(needleAttackShape),
		Width:    len(needleAttackShape[0]),
		Zindex:   0,
		Shape:    needleAttackShape,
		Position: coordinate{y: 6, x: 4},
	}

	collisions := g.DetectCollisions(&a)

	assert.Assert(t, len(collisions) == 2)
	assert.Assert(t, collisions["hitbox1"] != nil)
	assert.Assert(t, collisions["hitbox2"] != nil)
	assert.Assert(t, collisions["hitbox3"] == nil)

	a2 := Hitbox{
		ID:       "attacker",
		Height:   len(needleAttackShape),
		Width:    len(needleAttackShape[0]),
		Zindex:   0,
		Shape:    needleAttackShape,
		Position: coordinate{y: 4, x: 4},
	}

	collisions = g.DetectCollisions(&a2)

	assert.Assert(t, len(collisions) == 3)
	assert.Assert(t, collisions["hitbox1"] != nil)
	assert.Assert(t, collisions["hitbox2"] != nil)
	assert.Assert(t, collisions["hitbox3"] != nil)
}

var spearShape [][]rune = [][]rune{
	[]rune{'s', ' '},
	[]rune{'s', 's'},
	[]rune{'s', 's'},
	[]rune{'s', 's'},
	[]rune{' ', 's'},
	[]rune{' ', 's'},
	[]rune{' ', 's'},
	[]rune{' ', 's'},
}

var needleAttackShape [][]rune = [][]rune{
	[]rune{'n', ' ', 'n', ' '},
	[]rune{' ', 'n', ' ', 'n'},
	[]rune{'n', ' ', 'n', ' '},
	[]rune{' ', 'n', ' ', 'n'},
	[]rune{'n', ' ', 'n', ' '},
}
