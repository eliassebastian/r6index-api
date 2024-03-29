package ubisoft

type ProfileModel struct {
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	ProfileID      string `json:"profileId"`
	UserID         string `json:"userId,omitempty"`
	PlatformType   string `json:"platformType"`
	NameOnPlatform string `json:"nameOnPlatform"`
}
