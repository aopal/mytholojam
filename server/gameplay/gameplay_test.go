package gameplay

import (
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
