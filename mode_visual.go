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
	army1 := gatherArmy(budget, 10, faction1)
	army2 :=  gatherArmy(budget, 10, faction2)
	b := battler.InitBattlefield(rnd, BATTLEFIELD_WIDTH, BATTLEFIELD_HEIGHT, faction1, faction2, army1, army2)
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
