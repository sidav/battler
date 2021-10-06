package main

import (
	"battler/battler"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf(
			"Arguments: \n" +
				" -v <budget> \n" +
				" -s <budget> <tries> \n")
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
	battler.InitUnitsData()
	if args[0] == "-v" {
		budget, _ := strconv.Atoi(args[1])
		faction1 := args[2]
		faction2 := args[3]
		doVisualMode(budget, faction1, faction2)
	}
	if args[0] == "-s" {
		budget, _ := strconv.Atoi(args[1])
		cycles := 100
		faction1 := args[2]
		faction2 := args[3]
		doStatistic(budget, cycles, faction1, faction2)
	}
}

