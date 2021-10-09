package battler

import (
	"battler/fibrandom"
	"fmt"
)

func InitUnitsData(generatedUnitsNumber int) {
	rnd = fibrandom.FibRandom{}
	rnd.InitDefault()
	UNITS_DATA = make(map[string]*UnitData, 0)
	for k, v := range other_units {
		UNITS_DATA[k] = v
		if UNITS_DATA[k].Name == "" {
			UNITS_DATA[k].Name = k
		}
	}
	for k, v := range guerillas_units {
		UNITS_DATA[k] = v
		if UNITS_DATA[k].Name == "" {
			UNITS_DATA[k].Name = k
		}
	}
	for k, v := range elites_units {
		UNITS_DATA[k] = v
		if UNITS_DATA[k].Name == "" {
			UNITS_DATA[k].Name = k
		}
	}
	for k, v := range zerg_units {
		UNITS_DATA[k] = v
		if UNITS_DATA[k].Name == "" {
			UNITS_DATA[k].Name = k
		}
	}
	for k, v := range arm_units {
		UNITS_DATA[k] = v
		if UNITS_DATA[k].Name == "" {
			UNITS_DATA[k].Name = k
		}
	}
	for i := 0; i < generatedUnitsNumber; i++ {
		key := fmt.Sprintf("Forgotten%d", i+1)
		UNITS_DATA[key] = GenerateRandomUnitData()
	}
}

func InitBattlefield(random fibrandom.FibRandom, sizex, sizey int, leftTeamName, rightTeamName string, team1unitcodes, team2unitcodes map[string]int) *Battlefield {
	rnd = random
	b := Battlefield{
		Sizex: sizex,
		Sizey: sizey,
		CurrentTick: 1,
	}
	b.LeftTeam = &Team{
		Name:  leftTeamName,
	}
	b.RightTeam = &Team{
		Name:  rightTeamName,
	}

	var unitsInRows [3]int // 0 - last, 1 - second, 2 - first
	for code, number := range team1unitcodes {
		for y := 0; y < number; y++ {
			xcoord := 0
			row := 0
			switch UNITS_DATA[code].Class {
			case CLASS_ASSAULT:
				xcoord = 2
				row = 2
			case CLASS_SUPPORT:
				xcoord = 1
				row = 1
			case CLASS_LONGRANGE:
				xcoord = 0
				row = 0
			}
			unitsInRows[row]++
			ycoord := b.Sizey/2 - unitsInRows[row]/2
			if unitsInRows[row] % 2 == 0 {
				ycoord = unitsInRows[row]/2 + b.Sizey/2
			}
			b.UnitsMap = append(b.UnitsMap, CreateUnit(xcoord, ycoord-1, code, b.LeftTeam))
		}
	}
	unitsInRows[0] = 0
	unitsInRows[1] = 0
	unitsInRows[2] = 0
	for code, number := range team2unitcodes {
		for y := 0; y < number; y++ {
			xcoord := 0
			row := 0
			switch UNITS_DATA[code].Class {
			case CLASS_ASSAULT:
				xcoord = sizex - 3
				row = 2
			case CLASS_SUPPORT:
				xcoord = sizex - 2
				row = 1
			case CLASS_LONGRANGE:
				xcoord = sizex - 1
				row = 0
			}
			unitsInRows[row]++
			ycoord := b.Sizey/2 - unitsInRows[row]/2
			if unitsInRows[row] % 2 == 0 {
				ycoord = unitsInRows[row]/2 + b.Sizey/2
			}
			b.UnitsMap = append(b.UnitsMap, CreateUnit(xcoord, ycoord-1, code, b.RightTeam))
		}
	}

	return &b
}
