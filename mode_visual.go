package main

import (
	"battler/battler"
	"battler/console_wrapper"
	"battler/fibrandom"
	"time"
)

func doVisualMode(budget int, faction1, faction2 string) {
	console_wrapper.Init_console()
	defer console_wrapper.Close_console()
	rnd := fibrandom.FibRandom{}
	rnd.InitDefault()
	army1 := gatherArmy(budget, BATTLEFIELD_HEIGHT, faction1)
	army2 :=  gatherArmy(budget, BATTLEFIELD_HEIGHT, faction2)
	b := battler.InitBattlefield(rnd, BATTLEFIELD_WIDTH, BATTLEFIELD_HEIGHT, faction1, faction2, army1, army2)
	for tick := 0; tick < SIMULATION_TOTAL_TICKS_VISUAL; tick++ {
		renderBattlefield(b)
		time.Sleep(TICK_LENGTH_MS * time.Millisecond)
		b.SimulateTick()
		if console_wrapper.ReadKeyAsync() == "ESCAPE" || b.LeftTeam.CurrentTotalCost == 0 || b.RightTeam.CurrentTotalCost == 0 {
			break
		}
	}
	renderBattlefield(b)
	console_wrapper.ReadKey()
}
