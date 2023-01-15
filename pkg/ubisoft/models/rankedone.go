package ubisoft

type RankedOutputModel struct {
	SeasonID  int            `json:"season_id"`
	SeasonTag string         `json:"season_tag"`
	Regions   []RankedSeason `json:"regions"`
}

type RankedModel struct {
	SeasonsPlayerSkillRecords []SeasonsPlayerSkillRecords `json:"seasons_player_skill_records"`
}

type SeasonsPlayerSkillRecords struct {
	SeasonID                  int                         `json:"season_id"`
	RegionsPlayerSkillRecords []RegionsPlayerSkillRecords `json:"regions_player_skill_records"`
}

type RegionsPlayerSkillRecords struct {
	Region                   string          `json:"region_id"`
	BoardsPlayerSkillRecords []RankedSeasons `json:"boards_player_skill_records"`
}

type RankedSeasons struct {
	Seasons []RankedSeason `json:"players_skill_records"`
}

type RankedSeason struct {
	Season      int16   `json:"season"`
	Rank        int16   `json:"rank"`
	MaxRank     int16   `json:"max_rank"`
	MaxMmr      float32 `json:"max_mmr"`
	Mmr         float32 `json:"mmr"`
	Deaths      int     `json:"deaths"`
	Kills       int     `json:"kills"`
	Abandons    int     `json:"abandons"`
	Wins        int     `json:"wins"`
	Losses      int     `json:"losses"`
	Leaderboard int     `json:"top_rank_position"`
	Region      string  `json:"region"`
}
