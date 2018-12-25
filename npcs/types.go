package npcs

// Power describes the attack
type Power struct {
	Attack  int
	Defense int
}

// Location describes where
type Location struct {
	X float64
	Y float64
	Z float64
}

// NonPlayerCharacter represent metadata for an in-game creature
type NonPlayerCharacter struct {
	Name  string
	Speed int
	HP    int
	Power Power
	Loc   Location
}
