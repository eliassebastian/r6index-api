package ubisoft

import (
	"fmt"

	"github.com/eliassebastian/r6index-api/pkg/utils"
)

const (
	CURRENTSEASON   = "Y8S1"
	SEASONSTARTDATE = "20230307"
)

func profileUri(name, uuid, platform string) string {
	if name != "" {
		return fmt.Sprintf("https://public-ubiservices.ubi.com/v3/profiles?namesOnPlatform=%s&platformType=%s", name, platform)
	}

	return fmt.Sprintf("https://public-ubiservices.ubi.com/v3/users/%s/profiles?platformType=%s", uuid, platform)
}

func xpUri(uuid, platform string) string {
	//platform = PlatformSpaceId[platform]
	return fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/0d2ae42d-4c27-4cb7-af6c-2099062302bb/title/r6s/rewards/public_profile?profile_id=%s", uuid)
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

func rankedTwoCurrentSeason(uuid string) string {
	return fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/0d2ae42d-4c27-4cb7-af6c-2099062302bb/sandboxes/OSBOR_XPLAY_LNCH_A/r6karma/player_skill_records?board_ids=pvp_ranked&region_ids=global&season_ids=-1&profile_ids=%s", uuid)
}

func weaponsUri(uuid, platform string, xplay bool) string {
	spaceId := PlatformSpaceId[platform]
	platform = PlatformModernStats[platform]
	return fmt.Sprintf("https://prod.datadev.ubisoft.com/v1/users/%s/playerstats?spaceId=%s&view=current&aggregation=weapons&gameMode=ranked&platformGroup=%s&teamRole=all&startDate=%s&endDate=%s", uuid, spaceId, platform, SEASONSTARTDATE, utils.GetDate(-1))
}

func mapUri(uuid, platform string, xplay bool) string {
	spaceId := PlatformSpaceId[platform]
	platform = PlatformModernStats[platform]
	return fmt.Sprintf("https://prod.datadev.ubisoft.com/v1/users/%s/playerstats?spaceId=%s&gameMode=ranked&platformGroup=%s&view=seasonal&aggregation=maps&teamRole=all&seasons=%s", uuid, spaceId, platform, CURRENTSEASON)
}

func operatorUri(uuid, platform string, xplay bool) string {
	spaceId := PlatformSpaceId[platform]
	platform = PlatformModernStats[platform]
	return fmt.Sprintf("https://prod.datadev.ubisoft.com/v1/users/%s/playerstats?spaceId=%s&gameMode=ranked&platformGroup=%s&view=seasonal&aggregation=operators&teamRole=Attacker,Defender&seasons=%s", uuid, spaceId, platform, CURRENTSEASON)
}

func trendsUri(uuid, platform string, xplay bool) string {
	spaceId := PlatformSpaceId[platform]
	platform = PlatformModernStats[platform]
	return fmt.Sprintf("https://prod.datadev.ubisoft.com/v1/users/%s/playerstats?spaceId=%s&view=current&aggregation=movingpoint&trendType=days&gameMode=ranked&platformGroup=%s&teamRole=all&startDate=%s&endDate=%s", uuid, spaceId, platform, SEASONSTARTDATE, utils.GetDate(-1))
}

func summaryUri(uuid, platform string, xplay bool) string {
	spaceId := PlatformSpaceId[platform]
	platform = PlatformModernStats[platform]
	return fmt.Sprintf("https://prod.datadev.ubisoft.com/v1/users/%s/playerstats?platformGroup=%s&view=seasonal&aggregation=summary&gameMode=ranked&seasons=%s&spaceId=%s&teamRole=all", uuid, platform, CURRENTSEASON, spaceId)
}
