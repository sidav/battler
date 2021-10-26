package battler

const (
	KINETIC = "KINETIC"
	ENERGY = "ENERGY"
	EXPLOSION = "EXPLOSION"
)

type UnitData struct {
	DisplayedChar    string
	Cost             int
	Name             string
	Class            string
	Factions         string
	Weapon           Weapon `json:"Weapon"`
	Armor            Armor  `json:"Armor"`
	MaxHitpoints     int
	NumInSquad       int
	MovementCooldown int
}

func (u *UnitData) normalize(code string) {
	if u.Name == "" {
		u.Name = code
	}
	if u.Weapon.PercentToHit == 0 {
		u.Weapon.PercentToHit = 30
	}
	if u.Armor.PercentToBlock == 0 {
		u.Armor.PercentToBlock = 30
	}
}

var UNITS_DATA map[string]*UnitData
