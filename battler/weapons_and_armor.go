package battler

type weapon struct {
	attackRating                int
	attackType                  int
	attackRange, attackCooldown int
}

type armor struct {
	values map[int]int
}

func (u *Unit) RollDamageOnArmor(armor *armor) int {
	w := u.Data.Weapon
	if armor.values[w.attackType] >= 10 {
		return 0
	}
	damage := 0
	if u.Data.NumInSquad > 1 {
		for currInSquad := 0; currInSquad < u.RemainingSquadSize; currInSquad++ {
			for i := 0; i < w.attackRating; i++ {
				dice := rnd.RandInRange(1, 10)
				if dice > armor.values[w.attackType] {
					damage++
				}
			}
		}
	} else {
		for i := 0; i < w.attackRating; i++ {
			dice := rnd.RandInRange(1, 10)
			if dice > armor.values[w.attackType] {
				damage++
			}
		}
	}
	return damage
}
