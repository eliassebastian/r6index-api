package models

import (
	"time"

	ubisoft "github.com/eliassebastian/r6index-api/pkg/ubisoft/models"
)

type AliasCache struct {
	Name       string
	LastUpdate time.Time
}

type Alias struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type ProfileCache struct {
	LastUpdate int64
	Aliases    *[]Alias
}

type PlayerFound struct {
	Id string `json:"profileId"`
}

type Player struct {
	//Id         string                          `json:"id"`
	ProfileId  string                          `json:"profileId"`
	UserId     string                          `json:"userId"`
	Platform   string                          `json:"platform"`
	Nickname   string                          `json:"nickname"`
	LastUpdate int64                           `json:"lastUpdate"`
	Aliases    *[]Alias                        `json:"aliases"`
	Xp         int32                           `json:"xp"`
	Level      int16                           `json:"level"`
	Summary    *ubisoft.SummaryTeamRoles       `json:"summary"`
	RankedOne  *[]ubisoft.RankedOutputModel    `json:"rankedOne"`
	RankedTwo  *[]ubisoft.RankedTwoOutputModel `json:"rankedTwo"`
	Weapons    *ubisoft.WeaponTeamRoles        `json:"weapons"`
	Maps       *ubisoft.MapsTeamRoles          `json:"maps"`
	Operators  *ubisoft.OperatorTeamRoles      `json:"operators"`
	Trends     *ubisoft.TrendOutput            `json:"trends"`
}

type PlayerUpdate struct {
	ProfileId  string                          `json:"profileId"`
	Nickname   string                          `json:"nickname,omitempty"`
	LastUpdate int64                           `json:"lastUpdate"`
	Aliases    *[]Alias                        `json:"aliases,omitempty"`
	Xp         int32                           `json:"xp"`
	Level      int16                           `json:"level"`
	Summary    *ubisoft.SummaryTeamRoles       `json:"summary"`
	RankedOne  *[]ubisoft.RankedOutputModel    `json:"rankedOne"`
	RankedTwo  *[]ubisoft.RankedTwoOutputModel `json:"rankedTwo"`
	Weapons    *ubisoft.WeaponTeamRoles        `json:"weapons"`
	Maps       *ubisoft.MapsTeamRoles          `json:"maps"`
	Operators  *ubisoft.OperatorTeamRoles      `json:"operators"`
	Trends     *ubisoft.TrendOutput            `json:"trends"`
}
