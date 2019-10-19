package gameplay

import (
	"mytholojam/server/resources"
	"mytholojam/server/types"
	"sync"
)

var gameList map[string]*types.Game
var glLock sync.RWMutex
var spiritTypes []string

func Init() {
	gameList = make(map[string]*types.Game)
}

func initializeDummyPlayer1(p *types.Player) {
	e1 := resources.EquipList["Sword"].NewEquipment()
	e2 := resources.EquipList["Shield"].NewEquipment()
	e3 := resources.EquipList["Bow"].NewEquipment()

	s1 := resources.SpiritList["Warrior"].NewSpirit()
	s2 := resources.SpiritList["Thief"].NewSpirit()

	p.Equipment[e1.ID] = e1
	p.Equipment[e2.ID] = e2
	p.Equipment[e3.ID] = e3

	p.Spirits[s1.ID] = s1
	p.Spirits[s2.ID] = s2

	s1.Inhabit(e1)
	s2.Inhabit(e2)
}

func initializeDummyPlayer2(p *types.Player) {
	e1 := resources.EquipList["Shield"].NewEquipment()
	e2 := resources.EquipList["Sword"].NewEquipment()
	e3 := resources.EquipList["Bow"].NewEquipment()

	s1 := resources.SpiritList["Thief"].NewSpirit()
	s2 := resources.SpiritList["Warrior"].NewSpirit()

	p.Equipment[e1.ID] = e1
	p.Equipment[e2.ID] = e2
	p.Equipment[e3.ID] = e3

	p.Spirits[s1.ID] = s1
	p.Spirits[s2.ID] = s2

	s1.Inhabit(e2)
	s2.Inhabit(e1)
}
