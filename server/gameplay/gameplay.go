package gameplay

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

var gameList map[string]*game
var glLock sync.RWMutex
var bufferSize int
var moveList map[string]*move
var equipList map[string]*equipmentTemplate
var spiritList map[string]*spiritTemplate
var types []string

func Init() {
	gameList = make(map[string]*game)
	bufferSize = 10
	moveList = make(map[string]*move)
	equipList = make(map[string]*equipmentTemplate)
	spiritList = make(map[string]*spiritTemplate)
	types = []string{"type1", "type2", "type3", "type4"}

	moveF, _ := ioutil.ReadFile("resources/moves.json")
	equipF, _ := ioutil.ReadFile("resources/equipment.json")
	spiritF, _ := ioutil.ReadFile("resources/spirits.json")

	_ = json.Unmarshal([]byte(moveF), &moveList)
	_ = json.Unmarshal([]byte(equipF), &equipList)
	_ = json.Unmarshal([]byte(spiritF), &spiritList)
}

func initializeDummyPlayer1(p *player) {
	e1 := equipList["sword"].NewEquipment()
	e2 := equipList["shield"].NewEquipment()
	e3 := equipList["bow"].NewEquipment()

	s1 := spiritList["warrior"].NewSpirit()
	s2 := spiritList["cleric"].NewSpirit()

	e1.ID = "dd6575b1-1b5a-4488-b939-d92185b6256c"
	e2.ID = "ab71ae00-f1b9-43b0-b9c8-534222c0633d"
	e3.ID = "cf8485fc-859a-498c-b7e5-5012c9cff328"

	s1.ID = "8bfd5081-7dd0-4d76-a726-c7b60497967d"
	s2.ID = "6d3853d8-fffb-492c-98ac-b849b447abc4"

	p.Equipment[e1.ID] = e1
	p.Equipment[e2.ID] = e2
	p.Equipment[e3.ID] = e3

	p.Spirits[s1.ID] = s1
	p.Spirits[s2.ID] = s2

	s1.Inhabit(e1)
	s2.Inhabit(e2)
}

func initializeDummyPlayer2(p *player) {
	e1 := equipList["helmet"].NewEquipment()
	e2 := equipList["breastplate"].NewEquipment()
	e3 := equipList["axe"].NewEquipment()

	s1 := spiritList["thief"].NewSpirit()
	s2 := spiritList["mage"].NewSpirit()

	e1.ID = "3a2c5b71-db2f-4b5a-bd4c-52a5f3250ea1"
	e2.ID = "e2a8f993-9b2b-427e-b6ff-e34f148801d4"
	e3.ID = "f714d4bd-fedb-468a-be86-f24546dc15c1"

	s1.ID = "e7dde1bf-0f4c-4b13-8c30-1c0e18c0458c"
	s2.ID = "ddb8c6cc-bd62-4033-b583-14f4a6d1fe28"

	p.Equipment[e1.ID] = e1
	p.Equipment[e2.ID] = e2
	p.Equipment[e3.ID] = e3

	p.Spirits[s1.ID] = s1
	p.Spirits[s2.ID] = s2

	s1.Inhabit(e3)
	s2.Inhabit(e1)
}
