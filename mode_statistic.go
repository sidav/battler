package main

import (
	"battler/battler"
	"battler/fibrandom"
	"fmt"
)

func findMostBalancedBudget(faction1, faction2 string) {
	currBudget := 10000
	//a1wins, a2wins := 0, 0
	stepSize := currBudget
	bestA1WinsValue, bestBudget := SIMULATION_CYCLES_BALANCE, 0
	for stepSize > 50 {
		rnd := &fibrandom.FibRandom{}
		rnd.InitDefault()
		fmt.Printf("Testing %d budget... ", currBudget)
		thisBudgetA1Wins := 0
		for cycle := 0; cycle < SIMULATION_CYCLES_BALANCE; cycle++ {
			army1won, _, _ := doStatisticSimulation(rnd, currBudget, faction1, faction2)
			if army1won {
				thisBudgetA1Wins++
			}
		}
		fmt.Printf("%d wins (%.1f%%)\n", thisBudgetA1Wins, calc50BalancePercent(thisBudgetA1Wins, SIMULATION_CYCLES_BALANCE))
		stepSize = currBudget/2
		if abs(SIMULATION_CYCLES_BALANCE/2 - thisBudgetA1Wins) <= abs(SIMULATION_CYCLES_BALANCE/2 - bestA1WinsValue) {
			bestA1WinsValue = thisBudgetA1Wins
			bestBudget = currBudget
			currBudget = currBudget/2
		} else {
			currBudget = currBudget/2 + currBudget/4
		}
	}
	fmt.Printf("Best budget is %d (%d-%d wins, %.1f%%)\n", bestBudget, bestA1WinsValue, SIMULATION_CYCLES_BALANCE-bestA1WinsValue,
		calc50BalancePercent(bestA1WinsValue, SIMULATION_CYCLES_BALANCE))
}

func doStatistic(budget, cycles int, faction1, faction2 string) {
	rnd := &fibrandom.FibRandom{}
	rnd.InitDefault()

	army1wins := 0
	army1moneylosses := 0
	army2wins := 0
	army2moneylosses := 0

	for cycle := 0; cycle < cycles; cycle++ {
		army1won, a1losses, a2losses := doStatisticSimulation(rnd, budget, faction1, faction2)
		if army1won {
			army1wins += 1
			army1moneylosses += a1losses
		} else {
			army2wins += 1
			army2moneylosses += a2losses
		}

	}
	fmt.Printf("WINS: %s %d - %d %s\n", faction1, army1wins, army2wins, faction2)
	a1LossesString := "n/a"
	if army1wins != 0 {
		a1LossesString = fmt.Sprintf("%d", army1moneylosses/army1wins)
	}
	a2LossesString := "n/a"
	if army2wins != 0 {
		a2LossesString = fmt.Sprintf("%d", army2moneylosses/army2wins)
	}
	fmt.Printf("MEAN MONEY LOST: %s %s - %s %s\n", faction1, a1LossesString, a2LossesString, faction2)
}

func doStatisticSimulation(rnd *fibrandom.FibRandom, budget int, faction1, faction2 string) (bool, int, int) {
	army1won := false
	army1 := gatherArmy(budget, BATTLEFIELD_HEIGHT, faction1)
	army2 :=  gatherArmy(budget, BATTLEFIELD_HEIGHT, faction2)
	b := battler.InitBattlefield(*rnd, BATTLEFIELD_WIDTH, BATTLEFIELD_HEIGHT, "BLUES", "REDS", army1, army2)
	for tick := 0; tick < SIMULATION_TOTAL_TICKS_STATISTIC; tick++ {
		b.SimulateTick()
	}
	if b.LeftTeam.CurrentTotalCost > b.RightTeam.CurrentTotalCost {
		army1won = true
	}
	if b.LeftTeam.CurrentTotalCost < b.RightTeam.CurrentTotalCost {
		army1won = false
	}
	army1moneylosses := budget - b.LeftTeam.CurrentTotalCost
	army2moneylosses := budget - b.RightTeam.CurrentTotalCost
	return army1won, army1moneylosses, army2moneylosses
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calc50BalancePercent(current, twiceIdeal int) float64 {
	if current > twiceIdeal/2 {
		current = twiceIdeal - current
	}
	return 100 * float64(current) / float64(twiceIdeal/2)
}
