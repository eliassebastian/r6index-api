package ubisoft

import (
	"fmt"

	"github.com/eliassebastian/r6index-api/pkg/utils"
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

func xpUri(uuid, platform string) string {
	platform = PlatformSpaceId[platform]
	return fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/%s/title/r6s/rewards/public_profile?profile_id=%s", platform, uuid)
}

func rankedOneUri(uuid, platform string, xplay bool) string {
	if xplay {
		platform = "xplay"
	}

	spaceId := PlatformSpaceId[platform]
	sandboxId := PlatformSandboxId[platform]

	return fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/%s/sandboxes/%s/r6karma/player_skill_records?board_ids=pvp_ranked&region_ids=emea,ncsa,apac,global&season_ids=-1,-2,-3,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,-21,-22&profile_ids=%s", spaceId, sandboxId, uuid)
}

func rankedTwoUri(uuid, platform string) string {
	platform = PlatformRankedTwo[platform]

	return fmt.Sprintf("https://public-ubiservices.ubi.com/v2/spaces/0d2ae42d-4c27-4cb7-af6c-2099062302bb/title/r6s/skill/full_profiles?profile_ids=%s&platform_families=%s", uuid, platform)
}

func weaponsUri(uuid, platform string, date int, xplay bool) string {
	var dateS string

	if date != 0 {
		if date < -120 {
			date = -120
		}

		dateStart := utils.GetDate(date)
		dateEnd := utils.GetDate(-1)

		if xplay {
			dateStart = "20221206"
		}

		dateS = fmt.Sprintf("&startDate=%s&endDate=%s", dateStart, dateEnd)
	}

	spaceId := PlatformSpaceId[platform]

	return fmt.Sprintf("https://prod.datadev.ubisoft.com/v1/users/%s/playerstats?spaceId=%s&view=current&aggregation=weapons&gameMode=ranked&platformGroup=%s&teamRole=all,attacker,defender%s", uuid, spaceId, platform, dateS)
}

func mapUri(uuid, platform string, xplay bool) string {
	//todo: move to const
	currentSeason := "Y7S4"
	spaceId := PlatformSpaceId[platform]

	return fmt.Sprintf("https://prod.datadev.ubisoft.com/v1/users/%s/playerstats?spaceId=%s&gameMode=ranked&platformGroup=%s&view=current&aggregation=maps&teamRole=all&seasons=%s", uuid, spaceId, platform, currentSeason)
}

func operatorUri(uuid, platform string, xplay bool) string {
	currentSeason := "Y7S4"
	spaceId := PlatformSpaceId[platform]

	return fmt.Sprintf("https://prod.datadev.ubisoft.com/v1/users/%s/playerstats?spaceId=%s&gameMode=ranked&platformGroup=%s&view=current&aggregation=operators&teamRole=Attacker,Defender&seasons=%s", uuid, spaceId, platform, currentSeason)
}

func trendsUri(uuid, platform string, xplay bool) string {
	spaceId := PlatformSpaceId[platform]
	dateS := fmt.Sprintf("&startDate=%s&endDate=%s", utils.GetDate(-30), utils.GetDate(-1))
	return fmt.Sprintf("https://prod.datadev.ubisoft.com/v1/users/%s/playerstats?spaceId=%s&view=current&aggregation=movingpoint&trendType=days&gameMode=ranked&platformGroup=%s&teamRole=all%s", uuid, spaceId, platform, dateS)
}

func summaryUri(uuid, platform string, xplay bool) string {
	currentSeason := "Y7S4"
	spaceId := PlatformSpaceId[platform]
	return fmt.Sprintf("https://prod.datadev.ubisoft.com/v1/users/%s/playerstats?spaceId=%s&view=current&aggregation=summary&gameMode=ranked&platformGroup=%s&teamRole=all&seasons=%s", uuid, spaceId, platform, currentSeason)
}
