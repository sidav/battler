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
		X:                  x,
		Y:                  y,
		hitpoints:          UNITS_DATA[code].MaxHitpoints,
		RemainingSquadSize: UNITS_DATA[code].NumInSquad,
		NextTickToAct:      0,
		Data:               UNITS_DATA[code],
		Team:               team,
	}
}

func (u *Unit) isSquad() bool {
	return u.Data.NumInSquad > 1 && u.RemainingSquadSize > 1
}

func (attacker *Unit) AttackTarget(target *Unit) {
	damage := 0
	if attacker.isSquad() {
		for currInSquad := 0; currInSquad < attacker.RemainingSquadSize; currInSquad++ {
			damage += attacker.Data.Weapon.rollDamage()
		}
	} else {
		damage = attacker.Data.Weapon.rollDamage()
	}
	target.ReceiveDamage(damage, attacker.Data.Weapon.AttackType)
}

func (u *Unit) ReceiveDamage(damage int, damageType string) {
	if u.isSquad() {
		for damage > 0 && u.RemainingSquadSize > 0 {
			blockedDamage := u.Data.Armor.rollArmor(damageType)
			damage -= blockedDamage
			if damage > 0 {
				if damage > u.hitpoints {
					damage -= u.hitpoints
					u.hitpoints = 0
				} else {
					u.hitpoints -= damage
					damage = 0
				}
			}
			if u.hitpoints <= 0 {
				u.RemainingSquadSize--
				if u.RemainingSquadSize > 0 {
					u.hitpoints = u.Data.MaxHitpoints
				}
			}
		}
	} else {
		blockedDamage := u.Data.Armor.rollArmor(damageType)
		damage -= blockedDamage
		if damage > 0 {
			u.hitpoints -= damage
		}
	}
}

func (u *Unit) ExportStringStatsData() []string {
	str := make([]string, 0)
	str = append(str, fmt.Sprintf("%s, %d/%d hp", u.Data.Name, u.hitpoints, u.Data.MaxHitpoints))
	str = append(str, fmt.Sprintf("Weapon: %s AR %d toHit %d%% RNG %d CD %d",
		u.Data.Weapon.AttackType, u.Data.Weapon.AttackRating, u.Data.Weapon.PercentToHit, u.Data.Weapon.AttackRange,
		u.Data.Weapon.AttackCooldown))
	str = append(str, fmt.Sprintf("Armor %v toBlock %d%%", u.Data.Armor.Values, u.Data.Armor.PercentToBlock))
	str = append(str, fmt.Sprintf("Movement %d", u.Data.MovementCooldown))
	return str
}
