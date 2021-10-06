package main

import (
	"battler/battler"
	"battler/console_wrapper"
	"battler/fibrandom"
	"time"
)

func doVisualMode(budget int) {
	console_wrapper.Init_console()
	defer console_wrapper.Close_console()
	rnd := fibrandom.FibRandom{}
	rnd.InitDefault()
	b := battler.InitBattlefield(rnd, 20, 10, "BLUES", "REDS", gatherArmy(budget, 10), gatherArmy(budget, 10))
	for tick := 0; tick < 1000; tick++ {
		renderBattlefield(b)
		time.Sleep(25 * time.Millisecond)
		b.SimulateTick()
		if console_wrapper.ReadKeyAsync() == "ESCAPE" || b.LeftTeam.CurrentTotalCost == 0 || b.RightTeam.CurrentTotalCost == 0 {
			break
		}
	}
	renderBattlefield(b)
	console_wrapper.ReadKey()
}
