package main

import (
	"battler/battler"
	"battler/fibrandom"
	"strings"
)

func gatherArmy(budget, maxClass int, faction string) map[string]int {
	rnd := fibrandom.FibRandom{}
	rnd.InitDefault()
	currClasses := map[int]int {
		battler.CLASS_ASSAULT: 0,
		battler.CLASS_LONGRANGE: 0,
		battler.CLASS_SUPPORT: 0,
	}
	var army = make(map[string]int, 0)
	selectCycles := 0
	for budget > 0 && selectCycles < 10000 {
		selectCycles++
		affordablesList := getListOfUnitCodesCheaperThan(budget, faction)
		if len(affordablesList) == 0 {
			break
		}
		code := affordablesList[rnd.Rand(len(affordablesList))]
		class := battler.UNITS_DATA[code].Class
		if currClasses[class] == maxClass {
			continue
		}
		if _, ok := army[code]; ok {
			army[code] += 1
		} else {
			army[code] = 1
		}
		currClasses[class]++
		budget -= battler.UNITS_DATA[code].Cost
	}
	return army
}

func getListOfUnitCodesCheaperThan(value int, faction string) []string {
	var list []string
	for k, v := range battler.UNITS_DATA {
		if v.Cost <= value && (faction == "" || v.Factions == "" || strings.Contains(v.Factions, faction)) {
			list = append(list, k)
		}
	}
	return list
}
