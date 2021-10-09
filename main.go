package main

import (
	"battler/battler"
	"fmt"
	"os"
	"strconv"
)

const (
	BATTLEFIELD_WIDTH = 25
	BATTLEFIELD_HEIGHT = 10

	SIMULATION_TOTAL_TICKS_VISUAL    = 10000
	SIMULATION_TOTAL_TICKS_STATISTIC = 1000
	SIMULATION_TOTAL_TICKS_BALANCE  = 1000
	SIMULATION_CYCLES_BALANCE = 100

	TICK_LENGTH_MS = 10
)

func main() {
	args := os.Args[1:]

	if len(args) <= 1 {
		fmt.Printf(
			"Arguments: \n" +
				" -v <faction1> <faction2> <budget> \n" +
				" -s <faction1> <faction2> <budget> <tries> \n")
		return
	}

	//team1units := map[string]int{
	//	"TANK": 6,
	//	"ARTY": 3,
	//	"MISSILE": 2,
	//}
	//team2units := map[string]int{
	//	"TANK": 4,
	//	"TRIKE": 4,
	//	"JUGGERNAUT": 1,
	//}
	faction1 := ""
	if len(args) > 2 {
		faction1 = args[1]
	}
	faction2 := ""
	if len(args) > 2 {
		faction2 = args[2]
	}
	budget := 500
	if len(args) > 3 {
		budget, _ = strconv.Atoi(args[3])
	}

	battler.InitUnitsData(1)
	if args[0] == "-b" {
		findMostBalancedBudget(faction1, faction2)
	}
	if args[0] == "-g" {
		battler.InitUnitsData(5)
		for k, _ := range battler.UNITS_DATA {
			if battler.UNITS_DATA[k].Factions == "generated" {
				ustrings := battler.CreateUnit(0, 0, k, &battler.Team{}).ExportStringStatsData()
				for s := range ustrings {
					fmt.Println(ustrings[s])
				}
				fmt.Println("---")
			}
		}
	}
	if args[0] == "-v" {
		doVisualMode(budget, faction1, faction2)
	}
	if args[0] == "-s" {
		cycles := 100
		doStatistic(budget, cycles, faction1, faction2)
	}
}

