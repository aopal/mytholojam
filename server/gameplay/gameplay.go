package gameplay

import (
	"mytholojam/server/resources"
	"mytholojam/server/types"
	"sync"
)

var gameList map[string]*types.Game
var glLock sync.RWMutex
var bufferSize int

// var moveList map[string]*types.Move
// var equipList map[string]*types.EquipmentTemplate
// var spiritList map[string]*types.SpiritTemplate
var moveTypes []string

func Init() {
	gameList = make(map[string]*types.Game)
	bufferSize = 10
	// moveList = make(map[string]*types.Move)
	// equipList = make(map[string]*types.EquipmentTemplate)
	// spiritList = make(map[string]*types.SpiritTemplate)
	// moveTypes = []string{"STRN", "FLAM", "WEAR", "NAME"}

	// moveList = resources.MoveList
	// moveList["switch"] = resources.Switch
	// moveList["strong"] = resources.Switch
	// moveList["weak"] = resources.Switch
	// moveList["fast"] = resources.Switch

	// moveF, _ := ioutil.ReadFile("server/resources/moves.json")
	// equipF, _ := ioutil.ReadFile("server/resources/equipment.json")
	// spiritF, _ := ioutil.ReadFile("server/resources/spirits.json")

	// _ = json.Unmarshal([]byte(moveF), &moveList)
	// _ = json.Unmarshal([]byte(equipF), &equipList)
	// _ = json.Unmarshal([]byte(spiritF), &spiritList)
}

func initializeDummyPlayer1(p *types.Player) {
	e1 := resources.EquipList["Sword"].NewEquipment()
	e2 := resources.EquipList["Shield"].NewEquipment()
	// e3 := resources.EquipList["bow"].NewEquipment()

	s1 := resources.SpiritList["Warrior"].NewSpirit()
	s2 := resources.SpiritList["Thief"].NewSpirit()

	// e1.ID = "dd6575b1-1b5a-4488-b939-d92185b6256c"
	// e2.ID = "ab71ae00-f1b9-43b0-b9c8-534222c0633d"
	// e3.ID = "cf8485fc-859a-498c-b7e5-5012c9cff328"

	// s1.ID = "8bfd5081-7dd0-4d76-a726-c7b60497967d"
	// s2.ID = "6d3853d8-fffb-492c-98ac-b849b447abc4"

	p.Equipment[e1.ID] = e1
	p.Equipment[e2.ID] = e2
	// p.Equipment[e3.ID] = e3

	p.Spirits[s1.ID] = s1
	p.Spirits[s2.ID] = s2

	s1.Inhabit(e1)
	s2.Inhabit(e2)
}

func initializeDummyPlayer2(p *types.Player) {
	e1 := resources.EquipList["Shield"].NewEquipment()
	e2 := resources.EquipList["Sword"].NewEquipment()

	s1 := resources.SpiritList["Thief"].NewSpirit()
	s2 := resources.SpiritList["Warrior"].NewSpirit()

	p.Equipment[e1.ID] = e1
	p.Equipment[e2.ID] = e2

	p.Spirits[s1.ID] = s1
	p.Spirits[s2.ID] = s2

	s1.Inhabit(e2)
	s2.Inhabit(e1)
}
