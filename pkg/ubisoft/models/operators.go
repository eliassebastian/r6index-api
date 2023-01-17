package ubisoft

type OperatorDetail struct {
	Value float64 `json:"value"`
}

type Operator struct {
	//Type                   string         `json:"type"`
	//StatsType              string         `json:"statsType"`
	StatsDetail string `json:"statsDetail"`
	//SeasonYear             string         `json:"seasonYear"`
	//SeasonNumber           string         `json:"seasonNumber"`
	MatchesPlayed          int            `json:"matchesPlayed"`
	RoundsPlayed           int            `json:"roundsPlayed"`
	MinutesPlayed          int            `json:"minutesPlayed"`
	MatchesWon             int            `json:"matchesWon"`
	MatchesLost            int            `json:"matchesLost"`
	RoundsWon              int            `json:"roundsWon"`
	RoundsLost             int            `json:"roundsLost"`
	Kills                  int            `json:"kills"`
	Assists                int            `json:"assists"`
	Death                  int            `json:"death"`
	Headshots              int            `json:"headshots"`
	MeleeKills             int            `json:"meleeKills"`
	TeamKills              int            `json:"teamKills"`
	OpeningKills           int            `json:"openingKills"`
	OpeningDeaths          int            `json:"openingDeaths"`
	Trades                 int            `json:"trades"`
	OpeningKillTrades      int            `json:"openingKillTrades"`
	OpeningDeathTrades     int            `json:"openingDeathTrades"`
	Revives                int            `json:"revives"`
	DistanceTravelled      int            `json:"distanceTravelled"`
	WinLossRatio           float64        `json:"winLossRatio"`
	KillDeathRatio         OperatorDetail `json:"killDeathRatio"`
	HeadshotAccuracy       OperatorDetail `json:"headshotAccuracy"`
	KillsPerRound          OperatorDetail `json:"killsPerRound"`
	RoundsWithAKill        OperatorDetail `json:"roundsWithAKill"`
	RoundsWithMultiKill    OperatorDetail `json:"roundsWithMultiKill"`
	RoundsWithOpeningKill  OperatorDetail `json:"roundsWithOpeningKill"`
	RoundsWithOpeningDeath OperatorDetail `json:"roundsWithOpeningDeath"`
	RoundsWithKOST         OperatorDetail `json:"roundsWithKOST"`
	RoundsSurvived         OperatorDetail `json:"roundsSurvived"`
	RoundsWithAnAce        OperatorDetail `json:"roundsWithAnAce"`
	RoundsWithClutch       OperatorDetail `json:"roundsWithClutch"`
	TimeAlivePerMatch      float64        `json:"timeAlivePerMatch"`
	TimeDeadPerMatch       float64        `json:"timeDeadPerMatch"`
	DistancePerRound       float64        `json:"distancePerRound"`
}

type OperatorTeamRoles struct {
	Attacker []Operator `json:"attacker"`
	Defender []Operator `json:"defender"`
}

type OperatorGameMode struct {
	//Type      string    `json:"type"`
	TeamRoles OperatorTeamRoles `json:"teamRoles"`
}

type OperatorGameModes struct {
	Ranked OperatorGameMode `json:"ranked"`
}

type OperatorOutputModel struct {
	GameModes OperatorGameModes `json:"gameModes"`
}
