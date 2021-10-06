package battler

type Unit struct {
	X, Y               int
	hitpoints          int
	NextTickToAct      int
	Data               UnitData
	Team               *Team
	RemainingSquadSize int
}

func createUnit(x, y int, code string, team *Team) *Unit {
	return &Unit{
		X:             x,
		Y:             y,
		hitpoints:     UNITS_DATA[code].maxHitpoints,
		RemainingSquadSize: UNITS_DATA[code].NumInSquad,
		NextTickToAct: 0,
		Data:          UNITS_DATA[code],
		Team:          team,
	}
}

func (u *Unit) receiveDamage(damage int) {
	u.hitpoints -= damage
	if u.hitpoints <= 0 && u.Data.NumInSquad > 1 {
		for u.RemainingSquadSize > 0 && u.hitpoints <= 0 {
			u.RemainingSquadSize -= 1
			u.hitpoints+=u.Data.maxHitpoints
		}
		if u.RemainingSquadSize == 0 {
			u.hitpoints = 0
		}
	}
}
