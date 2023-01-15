package ubisoft

import (
	"fmt"
)

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

func rankedOneUri(uuid, platform string, xplay bool) string {
	if xplay {
		platform = "xplay"
	}

	spaceId := PlatformSpaceId[platform]
	sandboxId := PlatformSandboxId[platform]

	return fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/%s/sandboxes/%s/r6karma/player_skill_records?board_ids=pvp_ranked&region_ids=emea,ncsa,apac,global&season_ids=-6,-7,-8,-9,-10,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,-26,-27&profile_ids=%s", spaceId, sandboxId, uuid)
}
