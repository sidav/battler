package battler

var other_units = map[string]*UnitData{
	// Assault
	"TANK": {
		Cost:          10,
		DisplayedChar: 'T',
		Class:         CLASS_ASSAULT,
		Factions:      "elites,guerillas",
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
	"CHUCKNORRIS": {
		Cost:          1000,
		DisplayedChar: '&',
		Class:         CLASS_ASSAULT,
		Factions:      "chuck",
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
		Factions:      "elites,guerillas",
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

var elites_units = map[string]*UnitData{
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
	"ATTACK_ROBOT": {
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
	// Support
	"SHOCK": {
		Cost:          35,
		DisplayedChar: 's',
		Class:         CLASS_SUPPORT,
		Factions:      "elites",
		maxHitpoints:  2,
		NumInSquad:    3,
		Weapon: weapon{
			attackRating:   2,
			attackType:     ENERGY,
			attackRange:    4,
			attackCooldown: 35,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   2,
				ENERGY:    2,
				EXPLOSION: 2,
			},
		},
		movementCooldown: 30,
	},
	// long-range
	"RAILGUN": {
		Cost:          35,
		DisplayedChar: 'R',
		Class:         CLASS_LONGRANGE,
		Factions:      "elites",
		maxHitpoints:  5,
		Weapon: weapon{
			attackRating:   10,
			attackType:     ENERGY,
			attackRange:    15,
			attackCooldown: 50,
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
	"GAUSS": {
		Cost:          35,
		DisplayedChar: 'G',
		Class:         CLASS_LONGRANGE,
		Factions:      "elites",
		maxHitpoints:  3,
		Weapon: weapon{
			attackRating:   7,
			attackType:     KINETIC,
			attackRange:    12,
			attackCooldown: 40,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   1,
				ENERGY:    1,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 50,
	},
}

var arm_units = map[string]*UnitData{
	// Assault
	"PEEWEE": {
		Cost:          25,
		DisplayedChar: 'p',
		Class:         CLASS_ASSAULT,
		Factions:      "arm",
		maxHitpoints:  7,
		NumInSquad:    1,
		Weapon: weapon{
			attackRating:   3,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 5,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   4,
				ENERGY:    4,
				EXPLOSION: 4,
			},
		},
		movementCooldown: 10,
	},
	"ZEUS": {
		Cost:          25,
		DisplayedChar: 'Z',
		Class:         CLASS_ASSAULT,
		Factions:      "arm",
		maxHitpoints:  10,
		Weapon: weapon{
			attackRating:   7,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 25,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   7,
				ENERGY:    6,
				EXPLOSION: 7,
			},
		},
		movementCooldown: 20,
	},
	"COMMANDER": {
		Cost:          250,
		DisplayedChar: '@',
		Class:         CLASS_ASSAULT,
		Factions:      "arm",
		maxHitpoints:  25,
		Weapon: weapon{
			attackRating:   100,
			attackType:     ENERGY,
			attackRange:    1,
			attackCooldown: 50,
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
	// Support
	"FIDO": {
		Cost:          50,
		DisplayedChar: 'F',
		Class:         CLASS_SUPPORT,
		Factions:      "arm",
		maxHitpoints:  4,
		Weapon: weapon{
			attackRating:   6,
			attackType:     ENERGY,
			attackRange:    5,
			attackCooldown: 35,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   4,
				ENERGY:    4,
				EXPLOSION: 4,
			},
		},
		movementCooldown: 45,
	},
	// long-range
	"ANNIHILATOR": {
		Cost:          350,
		DisplayedChar: 'A',
		Class:         CLASS_LONGRANGE,
		Factions:      "arm",
		maxHitpoints:  5,
		Weapon: weapon{
			attackRating:   25,
			attackType:     ENERGY,
			attackRange:    20,
			attackCooldown: 150,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   1,
				ENERGY:    1,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 100,
	},
}

var guerillas_units = map[string]*UnitData{
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
				EXPLOSION: 0,
			},
		},
		movementCooldown: 10,
	},
	"SABOTEUR": {
		Cost:          100,
		DisplayedChar: 's',
		Class:         CLASS_ASSAULT,
		Factions:      "guerillas",
		maxHitpoints:  5,
		Weapon: weapon{
			attackRating:   25,
			attackType:     EXPLOSION,
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
		movementCooldown: 1,
	},
	"DEVASTATOR": {
		Cost:          75,
		DisplayedChar: 'D',
		Class:         CLASS_ASSAULT,
		Factions:      "guerillas",
		maxHitpoints:  25,
		Weapon: weapon{
			attackRating:   9,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   7,
				ENERGY:    5,
				EXPLOSION: 3,
			},
		},
		movementCooldown: 35,
	},
	// Support
	"SCORCHER": {
		Cost:          35,
		DisplayedChar: 'S',
		Class:         CLASS_SUPPORT,
		Factions:      "guerillas",
		maxHitpoints:  4,
		Weapon: weapon{
			attackRating:   6,
			attackType:     EXPLOSION,
			attackRange:    4,
			attackCooldown: 25,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   2,
				ENERGY:    2,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 30,
	},
	"LASER": {
		Cost:          45,
		DisplayedChar: 'L',
		Class:         CLASS_SUPPORT,
		Factions:      "guerillas",
		maxHitpoints:  7,
		Weapon: weapon{
			attackRating:   6,
			attackType:     ENERGY,
			attackRange:    4,
			attackCooldown: 25,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   2,
				ENERGY:    2,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 30,
	},
	// long-range
	"MISSILE": {
		Cost:          45,
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
	"SCUD": {
		Cost:          75,
		DisplayedChar: 'S',
		Class:         CLASS_LONGRANGE,
		Factions:      "guerillas",
		maxHitpoints:  3,
		Weapon: weapon{
			attackRating:   16,
			attackType:     EXPLOSION,
			attackRange:    5,
			attackCooldown: 50,
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

var zerg_units = map[string]*UnitData{
	// Assault
	"ZERGLING": {
		Cost:          1,
		DisplayedChar: 'z',
		Class:         CLASS_ASSAULT,
		Factions:      "zerg",
		maxHitpoints:  1,
		NumInSquad:    9,
		Weapon: weapon{
			attackRating:   1,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 5,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   2,
				ENERGY:    2,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 4,
	},
	"ROACH": {
		Cost:          8,
		DisplayedChar: 'r',
		Class:         CLASS_ASSAULT,
		Factions:      "zerg",
		maxHitpoints:  3,
		NumInSquad:    1,
		Weapon: weapon{
			attackRating:   3,
			attackType:     EXPLOSION,
			attackRange:    1,
			attackCooldown: 20,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   3,
				ENERGY:    2,
				EXPLOSION: 4,
			},
		},
		movementCooldown: 15,
	},
	"ULTRALISK": {
		Cost:          100,
		DisplayedChar: 'U',
		Class:         CLASS_ASSAULT,
		Factions:      "zerg",
		maxHitpoints:  25,
		NumInSquad:    0,
		Weapon: weapon{
			attackRating:   6,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 10,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   7,
				ENERGY:    7,
				EXPLOSION: 7,
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
		maxHitpoints:  3,
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
	"RAVAGER": {
		Cost:          15,
		DisplayedChar: 'G',
		Class:         CLASS_LONGRANGE,
		Factions:      "zerg",
		maxHitpoints:  3,
		NumInSquad:    0,
		Weapon: weapon{
			attackRating:   2,
			attackType:     KINETIC,
			attackRange:    5,
			attackCooldown: 15,
		},
		Armor: armor{
			values: map[int]int{
				KINETIC:   3,
				ENERGY:    3,
				EXPLOSION: 3,
			},
		},
		movementCooldown: 25,
	},
}
