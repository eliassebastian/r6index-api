package ubisoft

type MapDetail struct {
	Value float64 `json:"value"`
}

type Map struct {
	//Type                   string    `json:"type"`
	//StatsType              string    `json:"statsType"`
	StatsDetail string `json:"statsDetail"`
	//SeasonYear             string    `json:"seasonYear"`
	//SeasonNumber           string    `json:"seasonNumber"`
	MatchesPlayed          int       `json:"matchesPlayed"`
	RoundsPlayed           int       `json:"roundsPlayed"`
	MinutesPlayed          int       `json:"minutesPlayed"`
	MatchesWon             int       `json:"matchesWon"`
	MatchesLost            int       `json:"matchesLost"`
	RoundsWon              int       `json:"roundsWon"`
	RoundsLost             int       `json:"roundsLost"`
	Kills                  int       `json:"kills"`
	Assists                int       `json:"assists"`
	Death                  int       `json:"death"`
	Headshots              int       `json:"headshots"`
	MeleeKills             int       `json:"meleeKills"`
	TeamKills              int       `json:"teamKills"`
	OpeningKills           int       `json:"openingKills"`
	OpeningDeaths          int       `json:"openingDeaths"`
	Trades                 int       `json:"trades"`
	OpeningKillTrades      int       `json:"openingKillTrades"`
	OpeningDeathTrades     int       `json:"openingDeathTrades"`
	Revives                int       `json:"revives"`
	DistanceTravelled      int       `json:"distanceTravelled"`
	WinLossRatio           float64   `json:"winLossRatio"`
	KillDeathRatio         MapDetail `json:"killDeathRatio"`
	HeadshotAccuracy       MapDetail `json:"headshotAccuracy"`
	KillsPerRound          MapDetail `json:"killsPerRound"`
	RoundsWithAKill        MapDetail `json:"roundsWithAKill"`
	RoundsWithMultiKill    MapDetail `json:"roundsWithMultiKill"`
	RoundsWithOpeningKill  MapDetail `json:"roundsWithOpeningKill"`
	RoundsWithOpeningDeath MapDetail `json:"roundsWithOpeningDeath"`
	RoundsWithKOST         MapDetail `json:"roundsWithKOST"`
	RoundsSurvived         MapDetail `json:"roundsSurvived"`
	RoundsWithAnAce        MapDetail `json:"roundsWithAnAce"`
	RoundsWithClutch       MapDetail `json:"roundsWithClutch"`
	TimeAlivePerMatch      float64   `json:"timeAlivePerMatch"`
	TimeDeadPerMatch       float64   `json:"timeDeadPerMatch"`
	DistancePerRound       float64   `json:"distancePerRound"`
}

type MapsTeamRoles struct {
	All []Map `json:"all"`
	// Attacker []Map `json:"attacker"`
	// Defender []Map `json:"defender"`
}

type MapsGameMode struct {
	//Type      string        `json:"type"`
	TeamRoles MapsTeamRoles `json:"teamRoles"`
}

type MapsGameModes struct {
	Ranked MapsGameMode `json:"ranked"`
}

type MapsOutputModel struct {
	GameModes MapsGameModes `json:"gameModes"`
}
