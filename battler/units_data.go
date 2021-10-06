package battler

const (
	CLASS_ASSAULT = iota
	CLASS_LONGRANGE
	CLASS_SUPPORT

	KINETIC = iota
	ENERGY
	EXPLOSION
)

type UnitData struct {
	DisplayedChar    rune
	Cost             int
	Name             string
	Class            int
	weapon           weapon
	armor            armor
	maxHitpoints     int
	NumInSquad       int
	movementCooldown int
}

var UNITS_DATA = map[string]UnitData{
	// Assault
	"INFANTRY": {
		Cost:          1,
		DisplayedChar: 'i',
		Class:         CLASS_ASSAULT,
		maxHitpoints:  2,
		NumInSquad:    9,
		weapon: weapon{
			attackRating:   1,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		armor: armor{
			values: map[int]int{
				KINETIC:   1,
				ENERGY:    1,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 10,
	},
	"ELITE_COMMANDOS": {
		Cost:          25,
		DisplayedChar: 'e',
		Class:         CLASS_ASSAULT,
		maxHitpoints:  2,
		NumInSquad:    4,
		weapon: weapon{
			attackRating:   2,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		armor: armor{
			values: map[int]int{
				KINETIC:   7,
				ENERGY:    5,
				EXPLOSION: 4,
			},
		},
		movementCooldown: 10,
	},
	"TANK": {
		Cost:          10,
		DisplayedChar: 'T',
		Class:         CLASS_ASSAULT,
		maxHitpoints:  10,
		weapon: weapon{
			attackRating:   3,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		armor: armor{
			values: map[int]int{
				KINETIC:   3,
				ENERGY:    3,
				EXPLOSION: 3,
			},
		},
		movementCooldown: 20,
	},
	"TANK_HEAVY": {
		Cost:          25,
		DisplayedChar: 'H',
		Class:         CLASS_ASSAULT,
		maxHitpoints:  15,
		weapon: weapon{
			attackRating:   4,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 25,
		},
		armor: armor{
			values: map[int]int{
				KINETIC:   5,
				ENERGY:    3,
				EXPLOSION: 4,
			},
		},
		movementCooldown: 20,
	},
	"JUGGERNAUT": {
		Cost:          50,
		DisplayedChar: 'J',
		Class:         CLASS_ASSAULT,
		maxHitpoints:  100,
		weapon: weapon{
			attackRating:   5,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		armor: armor{
			values: map[int]int{
				KINETIC:   2,
				ENERGY:    2,
				EXPLOSION: 2,
			},
		},
		movementCooldown: 25,
	},
	"INDESTRUCTIBLE": {
		Cost:          150,
		DisplayedChar: '@',
		Class:         CLASS_ASSAULT,
		maxHitpoints:  10,
		weapon: weapon{
			attackRating:   3,
			attackType:     KINETIC,
			attackRange:    1,
			attackCooldown: 20,
		},
		armor: armor{
			values: map[int]int{
				KINETIC:   9,
				ENERGY:    9,
				EXPLOSION: 9,
			},
		},
		movementCooldown: 25,
	},
	// Support
	"ARTY": {
		Cost:          10,
		DisplayedChar: 'a',
		Class:         CLASS_SUPPORT,
		maxHitpoints:  5,
		weapon: weapon{
			attackRating:   5,
			attackType:     EXPLOSION,
			attackRange:    5,
			attackCooldown: 25,
		},
		armor: armor{
			values: map[int]int{
				KINETIC:   1,
				ENERGY:    1,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 30,
	},
	// long-range
	"MISSILE": {
		Cost:          15,
		DisplayedChar: 'M',
		Class:         CLASS_LONGRANGE,
		maxHitpoints:  5,
		weapon: weapon{
			attackRating:   10,
			attackType:     EXPLOSION,
			attackRange:    8,
			attackCooldown: 25,
		},
		armor: armor{
			values: map[int]int{
				KINETIC:   1,
				ENERGY:    1,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 30,
	},
	"RAILGUN": {
		Cost:          35,
		DisplayedChar: 'r',
		Class:         CLASS_LONGRANGE,
		maxHitpoints:  5,
		weapon: weapon{
			attackRating:   10,
			attackType:     ENERGY,
			attackRange:    15,
			attackCooldown: 40,
		},
		armor: armor{
			values: map[int]int{
				KINETIC:   1,
				ENERGY:    1,
				EXPLOSION: 1,
			},
		},
		movementCooldown: 30,
	},
}
