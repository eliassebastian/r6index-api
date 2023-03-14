package ubisoft

type SummaryDetail struct {
	Value float64 `json:"value"`
	//P     float64 `json:"p"`
}

type Summary struct {
	// Type                   string                 `json:"type"`
	// StatsType              string                 `json:"statsType"`
	// StatsDetail            string                 `json:"statsDetail"`
	// SeasonYear             string                 `json:"seasonYear"`
	// SeasonNumber           string                 `json:"seasonNumber"`
	// MatchesPlayed          int           `json:"matchesPlayed"`
	// RoundsPlayed           int           `json:"roundsPlayed"`
	// MinutesPlayed          int           `json:"minutesPlayed"`
	// MatchesWon             int           `json:"matchesWon"`
	// MatchesLost            int           `json:"matchesLost"`
	// RoundsWon              int           `json:"roundsWon"`
	// RoundsLost             int           `json:"roundsLost"`
	// Kills                  int           `json:"kills"`
	// Assists                int           `json:"assists"`
	// Death                  int           `json:"death"`
	// Headshots              int           `json:"headshots"`
	// MeleeKills             int           `json:"meleeKills"`
	// TeamKills              int           `json:"teamKills"`
	OpeningKills       int `json:"openingKills"`
	OpeningDeaths      int `json:"openingDeaths"`
	Trades             int `json:"trades"`
	OpeningKillTrades  int `json:"openingKillTrades"`
	OpeningDeathTrades int `json:"openingDeathTrades"`
	// Revives                int           `json:"revives"`
	// DistanceTravelled      int           `json:"distanceTravelled"`
	// WinLossRatio           float64       `json:"winLossRatio"`
	// KillDeathRatio         SummaryDetail `json:"killDeathRatio"`
	HeadshotAccuracy       SummaryDetail `json:"headshotAccuracy"`
	KillsPerRound          SummaryDetail `json:"killsPerRound"`
	RoundsWithAKill        SummaryDetail `json:"roundsWithAKill"`
	RoundsWithMultiKill    SummaryDetail `json:"roundsWithMultiKill"`
	RoundsWithOpeningKill  SummaryDetail `json:"roundsWithOpeningKill"`
	RoundsWithOpeningDeath SummaryDetail `json:"roundsWithOpeningDeath"`
	RoundsWithKOST         SummaryDetail `json:"roundsWithKOST"`
	RoundsSurvived         SummaryDetail `json:"roundsSurvived"`
	RoundsWithAnAce        SummaryDetail `json:"roundsWithAnAce"`
	RoundsWithClutch       SummaryDetail `json:"roundsWithClutch"`
	TimeAlivePerMatch      float64       `json:"timeAlivePerMatch"`
	TimeDeadPerMatch       float64       `json:"timeDeadPerMatch"`
	DistancePerRound       float64       `json:"distancePerRound"`
}
type SummaryTeamRoles struct {
	All []Summary `json:"all"`
}
type SummaryGameMode struct {
	TeamRoles SummaryTeamRoles `json:"teamRoles"`
}
type SummaryGameModes struct {
	Ranked SummaryGameMode `json:"ranked"`
}
type SummaryOutputModel struct {
	GameModes SummaryGameModes `json:"gameModes"`
}
