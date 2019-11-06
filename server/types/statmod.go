package types

type StatMod struct {
	ID        string
	AtkMod    int
	SpeedMod  int
	WeightMod int
	DefMods   map[string]int
}

func collect(mods []*StatMod, f func(*StatMod, int) int) int {
	result := 0
	for _, item := range mods {
		result = f(item, result)
	}
	return result
}

func cumAtkMod(mods []*StatMod) int {
	return collect(mods, func(s *StatMod, i int) int {
		return s.AtkMod + i
	})
}

func cumSpeedMod(mods []*StatMod) int {
	return collect(mods, func(s *StatMod, i int) int {
		return s.SpeedMod + i
	})
}

func cumWeightMod(mods []*StatMod) int {
	return collect(mods, func(s *StatMod, i int) int {
		return s.WeightMod + i
	})
}

func cumDefMod(mods []*StatMod, defType string) int {
	return collect(mods, func(s *StatMod, i int) int {
		return s.DefMods[defType] + i
	})
}
