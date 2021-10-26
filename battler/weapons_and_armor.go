package battler

type Weapon struct {
	AttackRating                int
	AttackType                  string
	PercentToHit                int
	AttackRange, AttackCooldown int
}

type Armor struct {
	Values         map[string]int
	PercentToBlock int
}

func (w *Weapon) rollDamage() int {
	toHit := 30
	if w.PercentToHit > 0 {
		toHit = w.PercentToHit
	}
	damage := 0
	for i := 0; i < w.AttackRating; i++ {
		dice := rnd.RandInRange(1, 100)
		if dice <= toHit {
			damage++
		}
	}
	return damage
}

func (a *Armor) rollArmor(damageType string) int {
	toBlock := 30
	if a.PercentToBlock > 0 {
		toBlock = a.PercentToBlock
	}
	blocks := 0
	for i := 0; i < a.Values[damageType]; i++ {
		if rnd.RandInRange(1, 100) <= toBlock {
			blocks++
		}
	}
	return blocks
}
