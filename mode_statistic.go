package main

import (
	"battler/battler"
	"battler/fibrandom"
	"fmt"
)

func doStatistic(budget, cycles int, faction1, faction2 string) {
	rnd := fibrandom.FibRandom{}
	rnd.InitDefault()

	army1wins := 0
	army1moneylosses := 0
	army2wins := 0
	army2moneylosses := 0

	for cycle := 0; cycle < cycles; cycle++ {
		army1 := gatherArmy(budget, 10, faction1)
		army2 :=  gatherArmy(budget, 10, faction2)
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
	fmt.Printf("WINS: %s %d - %d %s\n", faction1, army1wins, army2wins, faction2)
	fmt.Printf("MEAN MONEY LOST: %s %d - %d %s\n", faction1, army1moneylosses/cycles, army2moneylosses/cycles, faction2)
}
