package main

import (
	"battler/battler"
	"battler/fibrandom"
	"fmt"
)

func doStatistic(budget, cycles int) {
	rnd := fibrandom.FibRandom{}
	rnd.InitDefault()
	army1 := gatherArmy(budget, 10)
	army2 :=  gatherArmy(budget, 10)

	army1wins := 0
	army1moneylosses := 0
	army2wins := 0
	army2moneylosses := 0

	for cycle := 0; cycle < cycles; cycle++ {
		b := battler.InitBattlefield(rnd, 20, 10, "BLUES", "REDS", army1, army2)
		for tick := 0; tick < 1000; tick++ {
			b.SimulateTick()
		}
		if b.LeftTeam.CurrentTotalCost > b.RightTeam.CurrentTotalCost {
			army1wins++
		}
		if b.LeftTeam.CurrentTotalCost < b.RightTeam.CurrentTotalCost {
			army2wins++
		}
		army1moneylosses += budget - b.LeftTeam.CurrentTotalCost
		army2moneylosses += budget - b.RightTeam.CurrentTotalCost
	}
	fmt.Printf("WINS: BLUES %d - %d REDS\n", army1wins, army2wins)
	fmt.Printf("MEAN MONEY LOST: BLUES %d - %d REDS\n", army1moneylosses/cycles, army2moneylosses/cycles)
}
