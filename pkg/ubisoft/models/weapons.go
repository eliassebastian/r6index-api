package ubisoft

type WeaponsOutputModel struct {
	GameModes WeaponGameModes `json:"gameModes"`
}

type WeaponGameModes struct {
	Ranked WeaponRanked `json:"ranked"`
}

type WeaponRanked struct {
	//Type      string          `json:"type"`
	TeamRoles WeaponTeamRoles `json:"teamRoles"`
}

type WeaponTeamRoles struct {
	All WeaponTeamRole `json:"all"`
	// Attacker WeaponTeamRole `json:"attacker"`
	// Defender WeaponTeamRole `json:"defender"`
}

type WeaponTeamRole struct {
	WeaponSlots WeaponSlots `json:"weaponSlots"`
}

type WeaponSlots struct {
	SecondaryWeapons WeaponType `json:"secondaryWeapons"`
	PrimaryWeapons   WeaponType `json:"primaryWeapons"`
}

type WeaponType struct {
	WeaponTypes []WeaponTypes `json:"weaponTypes"`
}

type WeaponTypes struct {
	WeaponType string   `json:"weaponType"`
	Weapons    []Weapon `json:"weapons"`
}

type Weapon struct {
	WeaponName          string  `json:"weaponName"`
	RoundsPlayed        int     `json:"roundsPlayed"`
	RoundsWon           int     `json:"roundsWon"`
	RoundsLost          int     `json:"roundsLost"`
	Kills               int     `json:"kills"`
	Headshots           int     `json:"headshots"`
	HeadshotAccuracy    float64 `json:"headshotAccuracy"`
	RoundsWithAKill     float64 `json:"roundsWithAKill"`
	RoundsWithMultiKill float64 `json:"roundsWithMultiKill"`
	WeaponType          string  `json:"weaponType"`
	WeaponCategory      string  `json:"weaponCategory"`
}
