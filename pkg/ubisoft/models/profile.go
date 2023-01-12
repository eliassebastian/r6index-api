package ubisoft

type ProfileModel struct {
	Profiles []Profiles `json:"profiles"`
}

type Profiles struct {
	//platform specific uuid
	ProfileID string `json:"profileId"`
	//user uuid - note: stays the same across different platforms
	UserID       string `json:"userId"`
	PlatformType string `json:"platformType"`
	//platform ID
	IDOnPlatform   string `json:"idOnPlatform"`
	NameOnPlatform string `json:"nameOnPlatform"`
}
