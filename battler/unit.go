package battler

import "fmt"

type Unit struct {
	X, Y               int
	hitpoints          int
	NextTickToAct      int
	Data               *UnitData
	Team               *Team
	RemainingSquadSize int
}

func CreateUnit(x, y int, code string, team *Team) *Unit {
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

func (u *Unit) ReceiveDamage(damage int) {
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

func (u *Unit) ExportStringStatsData() []string {
	str := make([]string, 0)
	str = append(str, fmt.Sprintf("%s, %d/%d hp", u.Data.Name, u.hitpoints, u.Data.maxHitpoints))
	str = append(str, fmt.Sprintf("Weapon: AR %d RNG %d CD %d Type %d",
		u.Data.Weapon.attackRating, u.Data.Weapon.attackRange, u.Data.Weapon.attackCooldown, u.Data.Weapon.attackType))
	str = append(str, fmt.Sprintf("Armor %v", u.Data.Armor.values))
	str = append(str, fmt.Sprintf("Movement %d", u.Data.movementCooldown))
	return str
}
