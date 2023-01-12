package models

import "time"

type Player struct {
	ProfileId  string    `json:"profileId"`
	UserId     string    `json:"userId"`
	Platform   string    `json:"platform"`
	Nickname   string    `json:"nickname"`
	LastUpdate time.Time `json:"lastUpdate"`
}
