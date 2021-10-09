package battler

func GenerateRandomUnitData() *UnitData {
	return &UnitData{
		DisplayedChar:    '&',
		Cost:             -1, // rnd.Rand(200) * 10,
		Name:             "Forgotten Beast",
		Class:            CLASS_ASSAULT,
		Factions:         "generated",
		Weapon:           weapon{
			attackRating:   rnd.RandInRange(1, 10),
			attackType:     rnd.RandInRange(0, 2),
			attackRange:    rnd.RandInRange(1, 20),
			attackCooldown: rnd.RandInRange(1, 10) * 5,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   rnd.RandInRange(1, 12),
				ENERGY:   rnd.RandInRange(1, 12),
				EXPLOSION:   rnd.RandInRange(1, 12),
			},
		},
		maxHitpoints:     rnd.RandInRange(1, 50) * 10,
		NumInSquad:       0,
		movementCooldown: rnd.RandInRange(1, 20) * 2,
	}
}

//func GenerateRandomUnit(x, y int, team *Team) *Unit {
//
//	return &Unit{
//		X:             x,
//		Y:             y,
//		hitpoints:     data.maxHitpoints,
//		RemainingSquadSize: data.NumInSquad,
//		NextTickToAct: 0,
//		Data:          data,
//		Team:          team,
//	}
//}
