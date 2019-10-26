package gameplay

import (
	"fmt"
	"log"
	"mytholojam/server/resources"
	"mytholojam/server/types"
	"sync"
)

var gameList map[string]*types.Game
var glLock sync.RWMutex

func Init() {
	gameList = make(map[string]*types.Game)
}

func initializeDummyPlayer(p *types.Player) {
	e1 := resources.EquipList["AngloSaxonSpear"].NewEquipment()
	e2 := resources.EquipList["DanishNeedleSet"].NewEquipment()
	e3 := resources.EquipList["ArthiRod"].NewEquipment()
	e4 := resources.EquipList["IberianGoatSkull"].NewEquipment()
	e5 := resources.EquipList["AttisCrown"].NewEquipment()

	s1 := resources.SpiritList["Flame"].NewSpirit()
	s2 := resources.SpiritList["Hive"].NewSpirit()
	p.Equipment[e1.ID] = e1
	p.Equipment[e2.ID] = e2
	p.Equipment[e3.ID] = e3
	p.Equipment[e4.ID] = e4
	p.Equipment[e5.ID] = e5

	p.Spirits[s1.ID] = s1
	p.Spirits[s2.ID] = s2

	s1.Inhabit(e1)
	s2.Inhabit(e2)
}

func debugf(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}

func debug(a interface{}) {
	debugf("%v", a)
}

func print(g *types.Game, format string, a ...interface{}) {
	log.Printf("[game: "+g.GameID+"] "+format, a...)
}

// func printline()
