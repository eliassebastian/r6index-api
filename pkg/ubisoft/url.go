package ubisoft

import "fmt"

func profileUri(name, uuid, platform string) string {
	if name != "" {
		return fmt.Sprintf("https://public-ubiservices.ubi.com/v3/profiles?namesOnPlatform=%s&platformType=%s", name, platform)
	}

	return fmt.Sprintf("https://public-ubiservices.ubi.com/v3/users/%s/profiles?platformType=%s}", uuid, platform)
}

func playtimeUri(uuid string) string {

	return fmt.Sprintf("https://public-ubiservices.ubi.com/v1/profiles/stats?profileIds=%s&spaceId=0d2ae42d-4c27-4cb7-af6c-2099062302bb&statNames=PPvPTimePlayed,PPvETimePlayed,PTotalTimePlayed,PClearanceLevel", uuid)

}

func xpUri(uuid string) string {
	return fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/0d2ae42d-4c27-4cb7-af6c-2099062302bb/title/r6s/rewards/public_profile?profile_id=%s", uuid)
}
