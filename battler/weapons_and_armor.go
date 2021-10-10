package battler

type weapon struct {
	attackRating                int
	attackType                  int
	percentToHit                int
	attackRange, attackCooldown int
}

type armor struct {
	values         map[int]int
	percentToBlock int
}

func (w *weapon) rollDamage() int {
	toHit := 30
	if w.percentToHit > 0 {
		toHit = w.percentToHit
	}
	damage := 0
	for i := 0; i < w.attackRating; i++ {
		dice := rnd.RandInRange(1, 100)
		if dice <= toHit {
			damage++
		}
	}
	return damage
}

func (a *armor) rollArmor(damageType int) int {
	toBlock := 30
	if a.percentToBlock > 0 {
		toBlock = a.percentToBlock
	}
	blocks := 0
	for i := 0; i < a.values[damageType]; i++ {
		if rnd.RandInRange(1, 100) <= toBlock {
			blocks++
		}
	}
	return blocks
}
