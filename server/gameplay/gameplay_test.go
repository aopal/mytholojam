package gameplay

import (
	"mytholojam/server/resources"
	"mytholojam/server/types"
	"testing"

	"gotest.tools/assert"
)

func TestSetup(t *testing.T) {
	gameID := "TestSetup"
	expectedNumEqup := 5
	expectedNumSpirits := 2

	Init()

	// create game
	g, err, _ := createGame(gameID)
	assert.Assert(t, err == nil)

	// add second player
	err, _ = joinGame(g)
	assert.Assert(t, err == nil)

	// assert players got loaded correctly
	assert.Assert(t, g.Player1.ID != "")
	assert.Assert(t, len(g.Player1.Equipment) == expectedNumEqup)
	assert.Assert(t, len(g.Player1.Spirits) == expectedNumSpirits)

	assert.Assert(t, g.Player2.ID != "")
	assert.Assert(t, len(g.Player1.Equipment) == expectedNumEqup)
	assert.Assert(t, len(g.Player1.Spirits) == expectedNumSpirits)
}

func TestActions(t *testing.T) {
	gameID := "TestActions"

	Init()

	g, p1, p2 := initializeTest(t, gameID)

	p1act := &types.ActionPayload{
		Token: p1.ID,
		Actions: []*types.Action{
			&types.Action{
				User: p1.Spirits["Warrior"],
				Targets: []*types.Equipment{
					p2.Equipment["Sword"],
				},
				Move: resources.MoveList["Strong"],
				Turn: 1,
			},
			&types.Action{
				User: p1.Spirits["Thief"],
				Targets: []*types.Equipment{
					p2.Equipment["Shield"],
				},
				Move: resources.MoveList["Weak"],
				Turn: 1,
			},
		},
	}

	p2act := &types.ActionPayload{
		Token: p2.ID,
		Actions: []*types.Action{
			&types.Action{
				User: p2.Spirits["Warrior"],
				Targets: []*types.Equipment{
					p1.Equipment["Shield"],
				},
				Move: resources.MoveList["Strong"],
				Turn: 1,
			},
			&types.Action{
				User: p2.Spirits["Thief"],
				Targets: []*types.Equipment{
					p2.Equipment["Bow"],
				},
				Move: resources.MoveList["Switch"],
				Turn: 1,
			},
		},
	}

	err, _ := takeAction(g, p1act)
	assert.Assert(t, err == nil)

	err, _ = takeAction(g, p2act)
	assert.Assert(t, err == nil)

	assert.Assert(t, p1.Equipment["Sword"].InhabitedBy == p1.Spirits["Warrior"])
	assert.Assert(t, p1.Equipment["Shield"].InhabitedBy == p1.Spirits["Thief"])
	assert.Assert(t, p1.Equipment["Bow"].InhabitedBy == nil)

	assert.Assert(t, p2.Equipment["Sword"].InhabitedBy == p2.Spirits["Warrior"])
	assert.Assert(t, p2.Equipment["Shield"].InhabitedBy == nil)
	assert.Assert(t, p2.Equipment["Bow"].InhabitedBy == p2.Spirits["Thief"])

	p1a1Damage := calculateDamage(
		p1act.Actions[0].User,
		p1act.Actions[0].Targets[0].InhabitedBy,
		p1act.Actions[0].Move,
	)
	p1a2Damage := calculateDamage(
		p1act.Actions[1].User,
		p1act.Actions[1].Targets[0],
		p1act.Actions[1].Move,
	)
	p2a1Damage := calculateDamage(
		p2act.Actions[0].User,
		p2act.Actions[0].Targets[0].InhabitedBy,
		p2act.Actions[0].Move,
	)

	assert.Assert(t, p1.Spirits["Warrior"].HP == p1.Spirits["Warrior"].MaxHP-resources.RecoilDamage)
	assert.Assert(t, p1.Spirits["Thief"].HP == p1.Spirits["Thief"].MaxHP-p2a1Damage)

	assert.Assert(t, p2.Spirits["Warrior"].HP == p2.Spirits["Warrior"].MaxHP-p1a1Damage-resources.RecoilDamage)
	assert.Assert(t, p2.Spirits["Thief"].HP == p2.Spirits["Thief"].MaxHP)

	assert.Assert(t, p1.Equipment["Sword"].HP == p1.Equipment["Sword"].MaxHP)
	assert.Assert(t, p1.Equipment["Shield"].HP == p1.Equipment["Shield"].MaxHP)
	assert.Assert(t, p1.Equipment["Bow"].HP == p1.Equipment["Bow"].MaxHP)

	assert.Assert(t, p2.Equipment["Sword"].HP == p2.Equipment["Sword"].MaxHP)
	assert.Assert(t, p2.Equipment["Shield"].HP == p2.Equipment["Shield"].MaxHP-p1a2Damage)
	assert.Assert(t, p2.Equipment["Bow"].HP == p2.Equipment["Bow"].MaxHP)
}

func initializeTest(t *testing.T, gameID string) (*types.Game, *types.Player, *types.Player) {
	g := types.NewGame(gameID)
	p1 := createPlayerForTesting()
	p2 := createPlayerForTesting()

	g.Players[p1.ID] = p1
	g.Player1 = p1

	g.Players[p2.ID] = p2
	g.Player2 = p2

	p2.Opponent = g.Player1
	g.Player1.Opponent = p2

	assert.Assert(t, len(p1.Equipment) == 3)
	assert.Assert(t, len(p1.Spirits) == 2)

	assert.Assert(t, len(p2.Equipment) == 3)
	assert.Assert(t, len(p2.Spirits) == 2)

	return g, p1, p2
}

func createPlayerForTesting() *types.Player {
	p, _ := types.NewPlayer()

	e1 := resources.EquipList["Sword"].NewEquipment()
	e2 := resources.EquipList["Shield"].NewEquipment()
	e3 := resources.EquipList["Bow"].NewEquipment()
	e1.ID = "Sword"
	e2.ID = "Shield"
	e3.ID = "Bow"

	s1 := resources.SpiritList["Warrior"].NewSpirit()
	s2 := resources.SpiritList["Thief"].NewSpirit()
	s1.ID = "Warrior"
	s2.ID = "Thief"

	p.Equipment[e1.ID] = e1
	p.Equipment[e2.ID] = e2
	p.Equipment[e3.ID] = e3

	p.Spirits[s1.ID] = s1
	p.Spirits[s2.ID] = s2

	s1.Inhabit(e1)
	s2.Inhabit(e2)

	return p
}
