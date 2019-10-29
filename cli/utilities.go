package main

import (
	"fmt"
	"strings"

	"github.com/aopal/mytholojam/server/types"
)

func findEntities(user, move, target string) (*types.Spirit, *types.Move, *types.Equipment, error) {
	var u *types.Spirit = nil
	var m *types.Move = nil
	var t *types.Equipment = nil
	var multipleMatches bool = false

	for _, v := range currentPlayer().Spirits {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(user)) {
			if u != nil {
				u = nil
				multipleMatches = true
				break
			}
			u = v
		}
	}

	if u == nil {
		return nil, nil, nil, fmt.Errorf("Invalid user, could not find spirit for search string " + user)
	} else if multipleMatches {
		return nil, nil, nil, fmt.Errorf("Invalid user, multiple spirits match search string " + user)
	}

	for _, v := range u.Moves {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(move)) {
			if m != nil {
				m = nil
				multipleMatches = true
				break
			}
			m = v
		}
	}

	for _, v := range u.Inhabiting.Moves {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(move)) {
			if m != nil {
				m = nil
				multipleMatches = true
				break
			}
			m = v
		}
	}

	if m == nil {
		return nil, nil, nil, fmt.Errorf("Invalid move, could not find move matching search string " + move + " for user " + u.Name)
	} else if multipleMatches {
		return nil, nil, nil, fmt.Errorf("Invalid move, multiple moves match search string " + move + " for user " + u.Name)
	}

	var teamTargeted map[string]*types.Equipment

	if m.TeamTargetable == "self" {
		teamTargeted = currentPlayer().Equipment
	} else if m.TeamTargetable == "other" {
		teamTargeted = opponent().Equipment
	}

	for _, v := range teamTargeted {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(target)) {
			if t != nil {
				t = nil
				multipleMatches = true
				break
			}
			t = v
		}
	}

	if t == nil {
		return nil, nil, nil, fmt.Errorf("Invalid target, could not find target matching search string " + target)
	} else if multipleMatches {
		return nil, nil, nil, fmt.Errorf("Invalid move, multiple targets match search string " + target)
	}

	return u, m, t, nil
}
