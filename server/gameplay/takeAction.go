package gameplay

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/aopal/mytholojam/server/types"

	"github.com/gorilla/mux"
)

func ActionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]

	if _, ok := gameList[gameID]; !ok {
		w.WriteHeader(400)
		w.Write([]byte("No game with that ID exists.\n"))
		return
	}

	g := gameList[gameID]
	g.Lock.Lock()
	defer g.Lock.Unlock()

	print(g, "Take action request received")

	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var payload types.ActionPayload
	json.Unmarshal(b, &payload)

	err, code := takeAction(g, &payload)

	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(""))
	}
}

func takeAction(g *types.Game, payload *types.ActionPayload) (error, int) {
	if payload.Token != g.Player1.ID && payload.Token != g.Player2.ID {
		return errors.New("Invalid token.\n"), 401
	}

	p := g.Players[payload.Token]

	// debug("validating actions")
	err, code := validateActions(g, p, payload.Actions)

	if err != nil {
		// debug(err)
		return err, code
	}

	// debug("Actions are valid")
	p.NextActions = payload.Actions

	if g.Player1.NextActions != nil && g.Player2.NextActions != nil {
		// debug("Actions submitted, calculating order")
		calculateActionOrder(g)
		g.TurnCount++
	}

	return nil, 200
}

func validateActions(g *types.Game, p *types.Player, actions []*types.Action) (error, int) {
	if len(actions) != len(p.Spirits) { // # spirits should only ever be 1 or 2
		return errors.New("Wrong number of actions.\n"), 400
	}

	if p.NextActions != nil {
		return errors.New("Actions already submitted.\n"), 400
	}

	for _, a := range actions {
		err, code := validateSingleAction(g, p, a)
		if err != nil {
			return err, code
		}
	}

	return nil, 200
}

func validateSingleAction(g *types.Game, p *types.Player, a *types.Action) (error, int) {
	if _, ok := p.Spirits[a.User.ID]; !ok {
		return errors.New("Invalid action1.\n"), 400
	}

	var teamTargeted map[string]*types.Equipment
	op := p.Opponent
	a.User = p.Spirits[a.User.ID]        // ensure client can't submit fake spirit stats
	a.Move = a.User.GetMove(a.Move.Name) // ensure client can't submit fake moves

	if a.User == nil {
		return errors.New("No spirit with the given id exists on the player.\n"), 400
	}

	if a.Move == nil {
		return errors.New("User doesn't have access to that move.\n"), 400
	}

	if len(a.Targets) == 0 || (len(a.Targets) > 1 && !a.Move.MultiTarget) {
		return errors.New("Invalid action2.\n"), 400
	}

	for _, t := range a.Targets {
		if a.Move.TeamTargetable == "self" {
			teamTargeted = p.Equipment
		} else if a.Move.TeamTargetable == "other" {
			teamTargeted = op.Equipment
		} else {
			return errors.New("Invalid action3.\n"), 400
		}

		if _, ok := teamTargeted[t.ID]; !ok {
			return errors.New("Invalid action4.\n"), 400
		}
	}

	// The target object in the action will just be a clone of the actual equipment with that id
	// so lookup the object in the proper map by id, and use it instead
	for i, _ := range a.Targets {
		a.Targets[i] = teamTargeted[a.Targets[i].ID]
	}

	return nil, 200
}

func calculateActionOrder(g *types.Game) {
	actions := append(
		g.Player1.NextActions,
		g.Player2.NextActions...,
	)

	sort.Slice(actions, func(i, j int) bool {
		if actions[i].Move.Priority == actions[j].Move.Priority {
			iEffSpd := actions[i].User.GetSpeed() - actions[i].User.Inhabiting.GetWeight()
			jEffSpd := actions[j].User.GetSpeed() - actions[j].User.Inhabiting.GetWeight()
			return iEffSpd > jEffSpd
		} else {
			return actions[i].Move.Priority > actions[j].Move.Priority
		}
	})

	for _, a := range actions {
		applyEffect(g, a)
	}

	g.ActionOrder = append(g.ActionOrder, actions...)
	g.NumActions += len(actions)

	g.Player1.NextActions = nil
	g.Player2.NextActions = nil
}

func applyEffect(g *types.Game, a *types.Action) {
	for _, t := range a.Targets {
		var target types.Damageable

		if t.Inhabited {
			target = t.InhabitedBy
		} else {
			target = t
		}

		damage := CalculateDamage(a.User, target, a.Move)
		damageDone := applyCallbacks(a.User, target, a.Move, damage)

		str := fmt.Sprintf("%v uses %v on %v for %v damage\n", a.User.Name, a.Move.Name, target.GetName(), damageDone)
		a.ActionText += str
		print(g, str)
	}
}

func applyCallbacks(user *types.Spirit, target types.Damageable, move *types.Move, damage int) int {
	hpBefore := target.GetHP()

	move.OnHit(user, target, move, damage)
	target.OnHit(user, target, move, damage)
	user.OnHit(user, target, move, damage)

	hpAfter := target.GetHP()

	return hpBefore - hpAfter
}

func CalculateDamage(user *types.Spirit, target types.Damageable, move *types.Move) int {
	effAtk := user.GetAtk() + move.Power
	damage := effAtk - target.GetDef(move.Type)

	if damage < 0 {
		damage = 0
	}

	return damage
}
