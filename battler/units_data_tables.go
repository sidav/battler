package battler

var other_units = map[string]UnitData {
	// Assault
	"TANK": {
		Cost:          10,
		DisplayedChar: 'T',
		Class:         CLASS_ASSAULT,
		Factions: "elites,guerillas",
		maxHitpoints:  10,
		Weapon: weapon{
			attackRating:   3,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   3,
				ENERGY:    3,
				EXPLOSION: 3,
			},
		},
		movementCooldown: 20,
	},
	"JUGGERNAUT": {
		Cost:          50,
		DisplayedChar: 'J',
		Class:         CLASS_ASSAULT,
		Factions: "elites,guerillas",
		maxHitpoints:  100,
		Weapon: weapon{
			attackRating:   5,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   2,
				ENERGY:    2,
				EXPLOSION: 2,
			},
		},
		movementCooldown: 25,
	},
	"CHUCKNORRIS": {
		Cost:          1000,
		DisplayedChar: '&',
		Class:         CLASS_ASSAULT,
		Factions: "chuck",
		maxHitpoints:  1,
		Weapon: weapon{
			attackRating:   100,
			attackType:     ENERGY,
			attackRange:    1,
			attackCooldown: 25,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   10,
				ENERGY:    10,
				EXPLOSION: 10,
			},
		},
		movementCooldown: 25,
	},
	// Support
	"ARTY": {
		Cost:          10,
		DisplayedChar: 'a',
		Class:         CLASS_SUPPORT,
		Factions: "elites,guerillas",
		maxHitpoints:  5,
		Weapon: weapon{
			attackRating:   5,
			attackType:     EXPLOSION,
			attackRange:    5,
			attackCooldown: 25,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   1,
				ENERGY:    1,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 30,
	},

	// long-range
}

var elites_units = map[string]UnitData {
	// Assault
	"ELITE_COMMANDOS": {
		Cost:          25,
		DisplayedChar: 'e',
		Class:         CLASS_ASSAULT,
		Factions:      "elites",
		maxHitpoints:  2,
		NumInSquad:    4,
		Weapon: weapon{
			attackRating:   2,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   5,
				ENERGY:    5,
				EXPLOSION: 4,
			},
		},
		movementCooldown: 10,
	},
	"TANK_HEAVY": {
		Cost:          25,
		DisplayedChar: 'H',
		Class:         CLASS_ASSAULT,
		Factions:      "elites",
		maxHitpoints:  15,
		Weapon: weapon{
			attackRating:   4,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 25,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   5,
				ENERGY:    3,
				EXPLOSION: 4,
			},
		},
		movementCooldown: 20,
	},
	"INDESTRUCTIBLE": {
		Cost:          150,
		DisplayedChar: '@',
		Class:         CLASS_ASSAULT,
		Factions:      "elites",
		maxHitpoints:  10,
		Weapon: weapon{
			attackRating:   3,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   9,
				ENERGY:    9,
				EXPLOSION: 9,
			},
		},
		movementCooldown: 25,
	},
	// long-range
	"RAILGUN": {
		Cost:          35,
		DisplayedChar: 'r',
		Class:         CLASS_LONGRANGE,
		Factions:      "elites",
		maxHitpoints:  5,
		Weapon: weapon{
			attackRating:   10,
			attackType:     ENERGY,
			attackRange:    15,
			attackCooldown: 40,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   1,
				ENERGY:    1,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 30,
	},
}

var guerillas_units = map[string]UnitData {
	// Assault
	"INFANTRY": {
		Cost:          9,
		DisplayedChar: 'i',
		Class:         CLASS_ASSAULT,
		Factions:      "guerillas",
		maxHitpoints:  2,
		NumInSquad:    6,
		Weapon: weapon{
			attackRating:   1,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   1,
				ENERGY:    1,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 10,
	},
	// Support
	// long-range
	"MISSILE": {
		Cost:          15,
		DisplayedChar: 'M',
		Class:         CLASS_LONGRANGE,
		Factions:      "guerillas",
		maxHitpoints:  5,
		Weapon: weapon{
			attackRating:   10,
			attackType:     EXPLOSION,
			attackRange:    8,
			attackCooldown: 25,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   1,
				ENERGY:    1,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 30,
	},
}

var zerg_units = map[string]UnitData {
	// Assault
	"ZERGLING": {
		Cost:          1,
		DisplayedChar: 'z',
		Class:         CLASS_ASSAULT,
		Factions:      "zerg",
		maxHitpoints:  1,
		NumInSquad:    4,
		Weapon: weapon{
			attackRating:   1,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   1,
				ENERGY:    1,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 7,
	},
	"ROACH": {
		Cost:          8,
		DisplayedChar: 'r',
		Class:         CLASS_ASSAULT,
		Factions:      "zerg",
		maxHitpoints:  3,
		NumInSquad:    1,
		Weapon: weapon{
			attackRating:   2,
			attackType:     EXPLOSION,
			attackRange:    1,
			attackCooldown: 20,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   3,
				ENERGY:    2,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 15,
	},
	// Support
	"HYDRALISK": {
		Cost:          10,
		DisplayedChar: 'h',
		Class:         CLASS_SUPPORT,
		Factions:      "zerg",
		maxHitpoints:  2,
		NumInSquad:    0,
		Weapon: weapon{
			attackRating:   3,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   1,
				ENERGY:    2,
				EXPLOSION: 3,
			},
		},
		movementCooldown: 25,
	},
	// long-range
}

