package types

import (
	"log"

	"github.com/google/uuid"
)

type StatMod struct {
	ID        string
	AtkMod    int
	SpeedMod  int
	WeightMod int
	DefMods   map[string]int
}

func collect(mods map[string]*StatMod, f func(*StatMod, int) int) int {
	result := 0
	for _, item := range mods {
		result = f(item, result)
	}
	return result
}

func cumAtkMod(mods map[string]*StatMod) int {
	return collect(mods, func(s *StatMod, i int) int {
		return s.AtkMod + i
	})
}

func cumSpeedMod(mods map[string]*StatMod) int {
	return collect(mods, func(s *StatMod, i int) int {
		return s.SpeedMod + i
	})
}

func cumWeightMod(mods map[string]*StatMod) int {
	return collect(mods, func(s *StatMod, i int) int {
		return s.WeightMod + i
	})
}

func cumDefMod(mods map[string]*StatMod, defType string) int {
	return collect(mods, func(s *StatMod, i int) int {
		return s.DefMods[defType] + i
	})
}

func NewStatMod() *StatMod {
	statmod := new(StatMod)
	statModID, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	statmod.ID = statModID.String()
	return statmod
}
