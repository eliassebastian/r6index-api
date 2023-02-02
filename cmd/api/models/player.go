package models

import (
	"time"

	ubisoft "github.com/eliassebastian/r6index-api/pkg/ubisoft/models"
)

type AliasCache struct {
	Name string
}

type Alias struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type Player struct {
	//Id         string                          `json:"id"`
	ProfileId  string                          `json:"profileId"`
	UserId     string                          `json:"userId"`
	Platform   string                          `json:"platform"`
	Nickname   string                          `json:"nickname"`
	LastUpdate time.Time                       `json:"lastUpdate"`
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
