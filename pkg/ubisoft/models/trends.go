package ubisoft

type TrendType struct {
	High    float64            `json:"high"`
	Average float64            `json:"average"`
	Low     float64            `json:"low"`
	Trend   map[string]float64 `json:"trend"`
	Actuals map[string]float64 `json:"actuals"`
}

type TrendTypeOutput struct {
	Name    string    `json:"name"`
	High    float64   `json:"high"`
	Average float64   `json:"average"`
	Low     float64   `json:"low"`
	Trend   []float64 `json:"trend"`
	Actuals []float64 `json:"actuals"`
}

type TrendOutput struct {
	MovingPoints float64           `json:"movingPoints"`
	Trends       []TrendTypeOutput `json:"trends"`
	// WinLossRatio           TrendTypeOutput `json:"winLossRatio"`
	// KillDeathRatio         TrendTypeOutput `json:"killDeathRatio"`
	// HeadshotAccuracy       TrendTypeOutput `json:"headshotAccuracy"`
	// KillsPerRound          TrendTypeOutput `json:"killsPerRound"`
	// RoundsWithAKill        TrendTypeOutput `json:"roundsWithAKill"`
	// RoundsWithMultiKill    TrendTypeOutput `json:"roundsWithMultiKill"`
	// RoundsWithOpeningKill  TrendTypeOutput `json:"roundsWithOpeningKill"`
	// RoundsWithOpeningDeath TrendTypeOutput `json:"roundsWithOpeningDeath"`
	// RoundsWithKOST         TrendTypeOutput `json:"roundsWithKOST"`
	// RoundsSurvived         TrendTypeOutput `json:"roundsSurvived"`
	// RatioTimeAlivePerMatch TrendTypeOutput `json:"ratioTimeAlivePerMatch"`
	// DistancePerRound       TrendTypeOutput `json:"distancePerRound"`
}

type Trend struct {
	Type                   string    `json:"type"`
	StatsType              string    `json:"statsType"`
	StatsDetail            string    `json:"statsDetail"`
	MovingPoints           int       `json:"movingPoints"`
	WinLossRatio           TrendType `json:"winLossRatio"`
	KillDeathRatio         TrendType `json:"killDeathRatio"`
	HeadshotAccuracy       TrendType `json:"headshotAccuracy"`
	KillsPerRound          TrendType `json:"killsPerRound"`
	RoundsWithAKill        TrendType `json:"roundsWithAKill"`
	RoundsWithMultiKill    TrendType `json:"roundsWithMultiKill"`
	RoundsWithOpeningKill  TrendType `json:"roundsWithOpeningKill"`
	RoundsWithOpeningDeath TrendType `json:"roundsWithOpeningDeath"`
	RoundsWithKOST         TrendType `json:"roundsWithKOST"`
	RoundsSurvived         TrendType `json:"roundsSurvived"`
	RatioTimeAlivePerMatch TrendType `json:"ratioTimeAlivePerMatch"`
	DistancePerRound       TrendType `json:"distancePerRound"`
}

type TrendsTeamRoles struct {
	All []Trend `json:"all"`
}

type TrendsRanked struct {
	//Type      string    `json:"type"`
	TrendTeamRoles TrendsTeamRoles `json:"teamRoles"`
}

type TrendsGameModes struct {
	Ranked TrendsRanked `json:"ranked"`
}

type TrendsOutputModel struct {
	GameModes TrendsGameModes `json:"gameModes"`
}
