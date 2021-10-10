package battler

import "strings"

func GenerateRandomUnitData() *UnitData {
	chars := []rune{'!', '@', '#', '$', '%', '&', '*', '+'}
	return &UnitData{
		DisplayedChar: chars[rnd.Rand(len(chars))],
		Cost:          -1, // rnd.Rand(200) * 10,
		Name:          generateRandomTitle() + generateRandomName(),
		Class:         CLASS_ASSAULT,
		Factions:      "generated",
		Weapon: weapon{
			attackRating:   rnd.RandInRange(5, 15),
			percentToHit:   rnd.RandInRange(1, 10) * 10,
			attackType:     rnd.RandInRange(0, 2),
			attackRange:    rnd.RandInRange(1, 20),
			attackCooldown: rnd.RandInRange(1, 10) * 5,
		},
		Armor: armor{
			percentToBlock: rnd.RandInRange(1, 9) * 10,
			values: map[int]int{
				KINETIC:   rnd.RandInRange(1, 20),
				ENERGY:    rnd.RandInRange(1, 20),
				EXPLOSION: rnd.RandInRange(1, 20),
			},
		},
		maxHitpoints:     rnd.RandInRange(1, 100) * 10,
		NumInSquad:       0,
		movementCooldown: rnd.RandInRange(1, 20) * 5,
	}
}

func generateRandomTitle() string {
	firsts := []string{
		"Forgotten",
		"Ancient",
		"Eldritch",
		"Monstrous",
		"Fearsome",
		"Terrifying",
		"Merciless",
		"Nightmarish",
		"Apocalyptic",
	}
	lasts := []string{
		"beast",
		"horror",
		"abomination",
		"spawn",
		"behemoth",
		"monstrosity",
		"devourer",
		"nightmare",
		"harbringer",
	}
	return firsts[rnd.Rand(len(firsts))] + " " + lasts[rnd.Rand(len(lasts))] + " "
}

func generateRandomName() string {
	glas := []string{"a", "o", "y", "i", "u", "hu"}
	soglas := []string{"x", "z", "q", "v", "j", "xl", "h"}
	length := rnd.RandInRange(2, 5)
	randomName := ""
	if rnd.OneChanceFrom(2) {
		randomName += strings.ToUpper(soglas[rnd.Rand(len(soglas))])
	} else {
		randomName += strings.ToUpper(glas[rnd.Rand(len(glas))])
	}

	for i := 0; i < length; i++ {
		randomName += glas[rnd.Rand(len(glas))]
		randomName += soglas[rnd.Rand(len(soglas))]
	}
	if rnd.OneChanceFrom(2) {
		randomName += glas[rnd.Rand(len(glas))]
	}

	return randomName
}
