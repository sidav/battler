package main

import (
	"battler/battler"
	"battler/console_wrapper"
	"fmt"
	"strconv"
)

func renderBattlefield(b *battler.Battlefield) {
	console_wrapper.Clear_console()
	renderField(b)
	renderUnits(b)
	console_wrapper.PutString(fmt.Sprintf("TICK %d", b.CurrentTick), 0, console_wrapper.CONSOLE_HEIGHT-1)
	console_wrapper.PutString(fmt.Sprintf("%s: %d", b.LeftTeam.Name, b.LeftTeam.CurrentTotalCost), 0, b.Sizey+1)
	putStringOnRight(fmt.Sprintf("%s: %d", b.RightTeam.Name, b.RightTeam.CurrentTotalCost), b.Sizey+1)
	console_wrapper.Flush_console()
}

func renderField(b *battler.Battlefield) {
	for x := 0; x < b.Sizex; x++ {
		for y:= 0; y < b.Sizey; y++ {
			console_wrapper.SetFgColor(console_wrapper.DARK_GREEN)
			console_wrapper.PutChar('.', x, y)
		}
	}
}

func renderUnits(b *battler.Battlefield) {
	color := console_wrapper.WHITE
	renderStats := true
	for _, u := range b.UnitsMap {
		if u.Team == b.LeftTeam {
			color = console_wrapper.DARK_BLUE
		} else {
			color = console_wrapper.DARK_RED
			if renderStats {
				stats := u.ExportStringStatsData()
				for y, s := range stats {
					putStringOnRight(s, y)
				}
				renderStats = false
			}
		}
		if u.NextTickToAct <= b.CurrentTick && false {
			console_wrapper.SetColor(console_wrapper.BLACK, color)
		} else {
			console_wrapper.SetColor(color, console_wrapper.BLACK)
		}
		charToRender := u.Data.DisplayedChar
		if u.Data.NumInSquad > 1 {
			charToRender = rune(strconv.Itoa(u.RemainingSquadSize)[0])
		}
		console_wrapper.PutChar(charToRender, u.X, u.Y)
	}
	console_wrapper.SetBgColor(console_wrapper.BLACK)
}

func putStringOnRight(str string, y int) {
	xcoord := console_wrapper.CONSOLE_WIDTH - len(str)
	console_wrapper.PutString(str, xcoord, y)
}
