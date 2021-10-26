package battler

func (b *Battlefield) SimulateTick() {
	for _, u := range b.UnitsMap {
		b.ActForUnit(u)
	}

	for unitIndex := len(b.UnitsMap)-1; unitIndex >= 0; unitIndex-- {
		u := b.UnitsMap[unitIndex]
		if u.hitpoints <= 0 {
			b.UnitsMap = append(b.UnitsMap[:unitIndex], b.UnitsMap[unitIndex+1:]...)
		}
	}
	b.CurrentTick++
	b.gatherCurrentStatus()
}

func (b *Battlefield) gatherCurrentStatus() {
	b.LeftTeam.CurrentTotalCost = 0
	b.RightTeam.CurrentTotalCost = 0
	for _, u := range b.UnitsMap {
		u.Team.CurrentTotalCost += u.Data.Cost
	}
}

func (b *Battlefield) ActForUnit(u *Unit) {
	if u.NextTickToAct > b.CurrentTick {
		return
	}
	if b.tryAttackClosest(u) {
		return
	}
	b.tryMoveUnitForward(u)
}

func (b *Battlefield) tryAttackClosest(attacker *Unit) bool {
	target := b.getClosestUnitInFrontOf(attacker, attacker.Data.Weapon.AttackRange, false, true)
	if target == nil {
		return false
	}
	attacker.NextTickToAct = b.CurrentTick + attacker.Data.Weapon.AttackCooldown
	attacker.AttackTarget(target)
	return true
}

func (b *Battlefield) tryMoveUnitForward(u *Unit) bool {
	if len(b.GetUnitsOfAnotherClassInFrontOf(u, 1, true, false)) != 0 {
		return false
	}
	if len(b.GetAllUnitsInFrontOf(u, 1, false, true)) == 0 {
		var direction int
		if u.Team == b.LeftTeam {
			direction = 1
		}
		if u.Team == b.RightTeam {
			direction = -1
		}
		u.X += direction
		u.NextTickToAct = b.CurrentTick + u.Data.MovementCooldown
		return true
	}
	return false
}
