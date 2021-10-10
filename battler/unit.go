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
		hitpoints:          UNITS_DATA[code].maxHitpoints,
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
	target.ReceiveDamage(damage, attacker.Data.Weapon.attackType)
}

func (u *Unit) ReceiveDamage(damage, damageType int) {
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
					u.hitpoints = u.Data.maxHitpoints
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
	str = append(str, fmt.Sprintf("%s, %d/%d hp", u.Data.Name, u.hitpoints, u.Data.maxHitpoints))
	str = append(str, fmt.Sprintf("Weapon: AR %d toHit %d%% RNG %d CD %d Type %d",
		u.Data.Weapon.attackRating, u.Data.Weapon.percentToHit, u.Data.Weapon.attackRange,
		u.Data.Weapon.attackCooldown, u.Data.Weapon.attackType))
	str = append(str, fmt.Sprintf("Armor %v toBlock %d%%", u.Data.Armor.values, u.Data.Armor.percentToBlock))
	str = append(str, fmt.Sprintf("Movement %d", u.Data.movementCooldown))
	return str
}
