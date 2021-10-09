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
	Factions         string
	Weapon           weapon
	Armor            armor
	maxHitpoints     int
	NumInSquad       int
	movementCooldown int
}

var UNITS_DATA map[string]*UnitData
