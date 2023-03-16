package ubisoft

type RankedTwoOutputModel struct {
	MaxRank         int16   `json:"max_rank"`
	SeasonID        int16   `json:"season"`
	Rank            int16   `json:"rank"`
	MaxRankPoints   int     `json:"max_rank_points"`
	RankPoints      int     `json:"rank_points"`
	TopRankPosition int     `json:"top_rank_position"`
	Abandons        int     `json:"abandons"`
	Losses          int     `json:"losses"`
	Wins            int     `json:"wins"`
	Deaths          int     `json:"deaths"`
	Kills           int     `json:"kills"`
	RankText        string  `json:"rank_text"`
	MaxRankText     string  `json:"max_rank_text"`
	KillDeathRatio  float32 `json:"kill_death_ratio"`
	WinLoseRatio    float32 `json:"win_lose_ratio"`
}

type RankedTwoModel struct {
	PlatformFamiliesFullProfiles []RankedTwoPlatformFamilies `json:"platform_families_full_profiles"`
}

type RankedTwoPlatformFamilies struct {
	BoardIdsFullProfiles []RankedTwoBoards `json:"board_ids_full_profiles"`
	PlatformFamily       string            `json:"platform_family"`
}

type RankedTwoBoards struct {
	BoardID      string                 `json:"board_id"`
	FullProfiles []RankedTwoFullProfile `json:"full_profiles"`
}

type RankedTwoFullProfile struct {
	Profile          RankedTwoProfile          `json:"profile"`
	SeasonStatistics RankedTwoSeasonStatistics `json:"season_statistics"`
}

type RankedTwoProfile struct {
	BoardID         string `json:"board_id"`
	ID              string `json:"id"`
	MaxRank         int16  `json:"max_rank"`
	MaxRankPoints   int    `json:"max_rank_points"`
	PlatformFamily  string `json:"platform_family"`
	Rank            int16  `json:"rank"`
	RankPoints      int    `json:"rank_points"`
	SeasonID        int16  `json:"season_id"`
	TopRankPosition int    `json:"top_rank_position"`
}

type RankedTwoSeasonStatistics struct {
	Deaths        int                    `json:"deaths"`
	Kills         int                    `json:"kills"`
	MatchOutcomes RankedTwoMatchOutcomes `json:"match_outcomes"`
}

type RankedTwoMatchOutcomes struct {
	Abandons int `json:"abandons"`
	Losses   int `json:"losses"`
	Wins     int `json:"wins"`
}
