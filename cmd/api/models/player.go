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

//go:generate go run github.com/shamaton/msgpackgen
type ProfileCache struct {
	LastUpdate int64
	Aliases    *[]Alias
}

type PlayerFound struct {
	Message  string `json:"message"`
	Nickname string `json:"nickname"`
	Id       string `json:"profileId"`
}

type PlayerUpdated struct {
	Id         string `json:"profileId"`
	LastUpdate int64  `json:"lastUpdate"`
}

type Player struct {
	//Id         string                          `json:"id"`
	ProfileId  string           `json:"profileId"`
	UserId     string           `json:"userId"`
	Platform   string           `json:"platform"`
	Nickname   string           `json:"nickname"`
	FirstIndex int64            `json:"firstIndex"`
	LastSeen   int64            `json:"lastSeen"`
	LastUpdate int64            `json:"lastUpdate"`
	Aliases    *[]Alias         `json:"aliases"`
	Xp         int32            `json:"xp"`
	Level      int16            `json:"level"`
	Summary    *ubisoft.Summary `json:"summary"`
	//RankedOne  *[]ubisoft.RankedOutputModel    `json:"rankedOne"`
	RankedTwo *ubisoft.RankedTwoOutputModel `json:"ranked"`
	Weapons   *[]ubisoft.Weapon             `json:"weapons"`
	Maps      *[]ubisoft.Map                `json:"maps"`
	Operators *[]ubisoft.Operator           `json:"operators"`
	Trends    *ubisoft.TrendOutput          `json:"trends"`
}

type PlayerUpdate struct {
	ProfileId  string           `json:"profileId"`
	Nickname   string           `json:"nickname,omitempty"`
	LastUpdate int64            `json:"lastUpdate"`
	LastSeen   int64            `json:"lastSeen"`
	Aliases    *[]Alias         `json:"aliases,omitempty"`
	Xp         int32            `json:"xp"`
	Level      int16            `json:"level"`
	Summary    *ubisoft.Summary `json:"summary,omitempty"`
	//RankedOne  *[]ubisoft.RankedOutputModel    `json:"rankedOne,omitempty"`
	RankedTwo *ubisoft.RankedTwoOutputModel `json:"ranked,omitempty"`
	Weapons   *[]ubisoft.Weapon             `json:"weapons,omitempty"`
	Maps      *[]ubisoft.Map                `json:"maps,omitempty"`
	Operators *[]ubisoft.Operator           `json:"operators,omitempty"`
	Trends    *ubisoft.TrendOutput          `json:"trends,omitempty"`
}
