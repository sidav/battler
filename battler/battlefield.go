package battler

import "battler/fibrandom"

var rnd fibrandom.FibRandom

type Battlefield struct {
	Sizex, Sizey int
	CurrentTick  int
	UnitsMap     []*Unit
	LeftTeam     *Team
	RightTeam    *Team
}

func (b *Battlefield) GetAllUnitsInFrontOf(u *Unit, horizontalRange int, ownTeam, enemyTeam bool) []*Unit {
	var lookFrom, lookTo int
	if u.Team == b.LeftTeam {
		direction := 1
		lookFrom = u.X + direction
		lookTo = u.X + direction * horizontalRange
	}
	if u.Team == b.RightTeam {
		direction := -1
		lookFrom = u.X + direction * horizontalRange
		lookTo = u.X + direction
	}
	units := make([]*Unit, 0)
	for _, currUnit := range b.UnitsMap {
		if u == currUnit {
			continue
		}
		if !ownTeam && u.Team == currUnit.Team {
			continue
		}
		if !enemyTeam && u.Team != currUnit.Team {
			continue
		}
		if currUnit.X >= lookFrom && currUnit.X <= lookTo {
			units = append(units, currUnit)
		}
	}
	return units
}

func (b *Battlefield) GetUnitsOfAnotherClassInFrontOf(u *Unit, horizontalRange int, ownTeam, enemyTeam bool) []*Unit {
	var lookFrom, lookTo int
	if u.Team == b.LeftTeam {
		direction := 1
		lookFrom = u.X + direction
		lookTo = u.X + direction * horizontalRange
	}
	if u.Team == b.RightTeam {
		direction := -1
		lookFrom = u.X + direction * horizontalRange
		lookTo = u.X + direction
	}
	units := make([]*Unit, 0)
	for _, currUnit := range b.UnitsMap {
		if u == currUnit {
			continue
		}
		if currUnit.Data.Class == u.Data.Class {
			continue
		}
		if !ownTeam && u.Team == currUnit.Team {
			continue
		}
		if !enemyTeam && u.Team != currUnit.Team {
			continue
		}
		if currUnit.X >= lookFrom && currUnit.X <= lookTo {
			units = append(units, currUnit)
		}
	}
	return units
}

func (b *Battlefield) getClosestUnitInFrontOf(u *Unit, horizontalRange int, ownTeam, enemyTeam bool) *Unit {
	var direction int
	// var lookFrom, lookTo int
	if u.Team == b.LeftTeam {
		direction = 1
		//lookFrom = u.X + direction
		//lookTo = u.X + direction * horizontalRange
	}
	if u.Team == b.RightTeam {
		direction = -1
		//lookFrom = u.X + direction * horizontalRange
		//lookTo = u.X + direction
	}
	for h := 1; h <= horizontalRange; h++ {
		for v := 0; v < b.Sizey; v++ {
			currX := u.X + h*direction
			currY1 := u.Y + v
			currY2 := u.Y - v
			for _, currUnit := range b.UnitsMap {
				if !ownTeam && u.Team == currUnit.Team {
					continue
				}
				if !enemyTeam && u.Team != currUnit.Team {
					continue
				}
				if currUnit.X == currX && (currUnit.Y == currY1 || currUnit.Y == currY2) {
					return currUnit
				}
			}
		}
	}
	return nil
}
