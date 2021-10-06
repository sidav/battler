package main

import (
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
	budget, _ := strconv.Atoi(args[1])
	cycles, _ := strconv.Atoi(args[2])
	if args[0] == "-v" {
		doVisualMode(budget)
	}
	if args[0] == "-s" {
		doStatistic(budget, cycles)
	}
}

