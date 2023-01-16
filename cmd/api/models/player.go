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
	ProfileId  string                          `json:"profileId"`
	UserId     string                          `json:"userId"`
	Platform   string                          `json:"platform"`
	Nickname   string                          `json:"nickname"`
	LastUpdate time.Time                       `json:"lastUpdate"`
	Aliases    *[]Alias                        `json:"aliases"`
	Xp         int32                           `json:"xp"`
	Level      int16                           `json:"level"`
	RankedOne  *[]ubisoft.RankedOutputModel    `json:"rankedOne"`
	RankedTwo  *[]ubisoft.RankedTwoOutputModel `json:"rankedTwo"`
}
