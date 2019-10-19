package gameplay

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mytholojam/server/resources"
	"mytholojam/server/types"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
)

func ActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Take action request received")

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

	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err, code := takeAction(g, b)

	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(""))
	}
}

func takeAction(g *types.Game, body []byte) (error, int) {
	var payload types.ActionPayload

	json.Unmarshal(body, &payload)

	if payload.Token != g.Player1.ID && payload.Token != g.Player2.ID {
		return errors.New("Invalid token.\n"), 401
	}

	p := g.Players[payload.Token]

	err, code := validateActions(g, p, payload.Actions)

	if err != nil {
		return err, code
	}

	p.NextActions = payload.Actions

	if g.Player1.NextActions != nil && g.Player2.NextActions != nil {
		calculateActionOrder(g)
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
	a.User = p.Spirits[a.User.ID]            // ensure client can't submit fake spirit stats
	a.Move = resources.MoveList[a.Move.Name] // ensure client can't submit fake moves

	if len(a.Targets) == 0 || len(a.Targets) > 1 && !a.Move.MultiTarget {
		return errors.New("Invalid action2.\n"), 400
	}

	// fmt.Println(a.Targets)
	// fmt.Println(*a.Move)

	for _, t := range a.Targets {
		if a.Move.TeamTargetable == "self" {
			teamTargeted = p.Equipment
		} else if a.Move.TeamTargetable == "other" {
			teamTargeted = op.Equipment
		} else {
			return errors.New("Invalid action3.\n"), 400
		}

		// fmt.Println(teamTargeted)
		// fmt.Println(t.ID)
		// fmt.Println(teamTargeted[t.ID])

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
			iEffSpd := actions[i].User.Speed - actions[i].User.Inhabiting.Weight
			jEffSpd := actions[j].User.Speed - actions[j].User.Inhabiting.Weight
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
}

func applyEffect(g *types.Game, a *types.Action) {
	if a.Move.Name == "switch" {
		fmt.Println(a.User.Name, "switches from", a.User.Inhabiting.Name, "to", a.Targets[0].Name)
		a.User.Inhabit(a.Targets[0])
	} else {
		effAtk := a.User.ATK + a.Move.Power

		for _, t := range a.Targets {
			var target types.Damageable

			if t.Inhabited {
				target = t.InhabitedBy
			} else {
				target = t
			}

			damage := effAtk - target.GetDef(a.Move.Type)
			if damage < 0 {
				damage = 0
			}
			target.TakeDamage(damage)

			fmt.Println(a.User.Name, "uses", a.Move.Name, "on", target.GetName(), "for", damage, "damage.")
		}
	}
}
