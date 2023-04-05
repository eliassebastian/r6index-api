package ubisoft

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/eliassebastian/r6index-api/pkg/auth"
	ubisoft "github.com/eliassebastian/r6index-api/pkg/ubisoft/models"
	"github.com/eliassebastian/r6index-api/pkg/utils"
)

func GetPlayerProfile(ctx context.Context, client client.Client, auth *auth.UbisoftSession, name, uuid, platform string) (*ubisoft.Profile, error) {

	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(profileUri(name, uuid, platform))
	requestHeaders(req, auth, false, false)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != consts.StatusOK {
		return nil, errors.New("ubisoft server profile response error (error code: #gpp1)")
	}

	var profile ubisoft.ProfileModel
	de := json.NewDecoder(res.BodyStream()).Decode(&profile)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rpp1)")
	}

	if len(profile.Profiles) == 0 {
		return nil, errors.New("profile not found")
	}

	for _, p := range profile.Profiles {
		if p.PlatformType == platform {
			if p.UserID != "" {
				return &p, nil
			}
		}
	}

	return nil, errors.New("profile not found")
}

func GetXpAndLevel(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string) (*ubisoft.XpAndLevel, error) {
	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(xpUri(uuid, platform))
	requestHeaders(req, auth, true, false)
	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	//ubisoft returned no content
	if res.StatusCode() == consts.StatusNoContent {
		return nil, nil
	}

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft xp and level response %v", res.StatusCode())
	}

	var xpLevel ubisoft.XpAndLevel
	de := json.NewDecoder(res.BodyStream()).Decode(&xpLevel)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rgxl1)")
	}

	return &xpLevel, nil
}

func GetPlaytime(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid string) {

}

func GetRankedOne(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string) (*[]ubisoft.RankedOutputModel, error) {
	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(rankedOneUri(uuid, platform, false))
	requestHeaders(req, auth, false, false)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	var rankedJson ubisoft.RankedModel
	de := json.NewDecoder(res.BodyStream()).Decode(&rankedJson)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rr11)")
	}

	var output []ubisoft.RankedOutputModel
	for _, season := range rankedJson.SeasonsPlayerSkillRecords {
		var regions []ubisoft.RankedSeason
		id := season.SeasonID

		for _, region := range season.RegionsPlayerSkillRecords {
			if len(region.BoardsPlayerSkillRecords) != 0 {
				regionInfo := region.BoardsPlayerSkillRecords[0].Seasons[0]
				if regionInfo.Wins+regionInfo.Losses+regionInfo.Abandons != 0 {
					rankInfoText := getRankInfo(id)
					regionInfo.MaxRankText = rankInfoText[regionInfo.MaxRank].Name
					regionInfo.RankText = rankInfoText[regionInfo.Rank].Name
					regions = append(regions, regionInfo)
				}
				if id > 17 {
					break
				}
			}
		}

		if len(regions) != 0 {
			seasonOut := ubisoft.RankedOutputModel{
				SeasonID:  id,
				SeasonTag: seasons[id].Code,
				Regions:   regions,
			}
			output = append(output, seasonOut)
		}
	}

	return &output, nil
}

func GetLastSeen(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string) (*time.Time, error) {
	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(rankedTwoCurrentSeason(uuid))
	requestHeaders(req, auth, false, false)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft last seen response %v", res.StatusCode())
	}

	var rankedJson ubisoft.RankedModel
	de := json.NewDecoder(res.BodyStream()).Decode(&rankedJson)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rr11)")
	}

	if len(rankedJson.SeasonsPlayerSkillRecords) == 0 {
		return nil, errors.New("no ranked seasons found")
	}

	return &rankedJson.SeasonsPlayerSkillRecords[0].RegionsPlayerSkillRecords[0].BoardsPlayerSkillRecords[0].Seasons[0].UpdateTime, nil
}

func GetRankedTwo(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string) (*ubisoft.RankedTwoOutputModel, error) {
	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(rankedTwoUri(uuid, platform))
	requestHeaders(req, auth, true, false)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft ranked 2.0 response %v", res.StatusCode())
	}

	var resJson ubisoft.RankedTwoModel
	de := json.NewDecoder(res.BodyStream()).Decode(&resJson)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rr21)")
	}

	if len(resJson.PlatformFamiliesFullProfiles) == 0 {
		return nil, nil
	}

	if len(resJson.PlatformFamiliesFullProfiles[0].BoardIdsFullProfiles) == 0 {
		return nil, nil
	}

	if len(resJson.PlatformFamiliesFullProfiles[0].BoardIdsFullProfiles[3].FullProfiles) == 0 {
		return nil, nil
	}

	var output []ubisoft.RankedTwoOutputModel

	for _, season := range resJson.PlatformFamiliesFullProfiles[0].BoardIdsFullProfiles[3].FullProfiles {
		rankInfo := getRankInfo(season.Profile.SeasonID)
		output = append(output, ubisoft.RankedTwoOutputModel{
			MaxRank:         season.Profile.MaxRank,
			SeasonID:        season.Profile.SeasonID,
			Rank:            season.Profile.Rank,
			MaxRankPoints:   season.Profile.MaxRankPoints,
			RankPoints:      season.Profile.RankPoints,
			TopRankPosition: season.Profile.TopRankPosition,
			Abandons:        season.SeasonStatistics.MatchOutcomes.Abandons,
			Losses:          season.SeasonStatistics.MatchOutcomes.Losses,
			Wins:            season.SeasonStatistics.MatchOutcomes.Wins,
			Deaths:          season.SeasonStatistics.Deaths,
			Kills:           season.SeasonStatistics.Kills,
			MaxRankText:     rankInfo[season.Profile.MaxRank].Name,
			RankText:        rankInfo[season.Profile.Rank].Name,
			WinLoseRatio:    float32(season.SeasonStatistics.MatchOutcomes.Wins) / float32(season.SeasonStatistics.MatchOutcomes.Losses),
			KillDeathRatio:  float32(season.SeasonStatistics.Kills) / float32(season.SeasonStatistics.Deaths),
			SeasonName:      seasons[season.Profile.SeasonID].Name,
		})
	}

	return &output[0], nil
}

func GetWeapons(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string, xplay bool) (*[]ubisoft.Weapon, error) {
	if xplay && platform != "uplay" {
		platform = "xplay"
	}

	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(weaponsUri(uuid, platform, xplay))
	requestHeaders(req, auth, false, true)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	//ubisoft returned no content
	if res.StatusCode() == consts.StatusNoContent {
		return nil, nil
	}

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft weapons response %v", res.StatusCode())
	}

	var op map[string]interface{}
	de := json.NewDecoder(res.BodyStream()).Decode(&op)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rw1)")
	}

	//dirty - find new solution
	weapons := op["profileData"].(map[string]interface{})[uuid].(map[string]interface{})["platforms"].(map[string]interface{})[PlatformModernStats[platform]]

	var result ubisoft.WeaponsOutputModel
	a, b := utils.Transcode(weapons, &result)

	//error encoding response
	if a != nil {
		return nil, errors.New("error decoding response (error code: #rw2)")
	}

	//ubisoft returned no data - handle error by returning nil
	if b != nil {
		return nil, nil
	}

	//result.GameModes.Ranked.TeamRoles.All.WeaponSlots.
	var output []ubisoft.Weapon
	for _, wt := range result.GameModes.Ranked.TeamRoles.All.WeaponSlots.SecondaryWeapons.WeaponTypes {
		for _, w := range wt.Weapons {
			w.WeaponType = wt.WeaponType
			w.WeaponCategory = "secondary"
			output = append(output, w)
		}
	}

	for _, wt := range result.GameModes.Ranked.TeamRoles.All.WeaponSlots.PrimaryWeapons.WeaponTypes {
		for _, w := range wt.Weapons {
			w.WeaponType = wt.WeaponType
			w.WeaponCategory = "primary"
			output = append(output, w)
		}
	}

	return &output, nil
}

func GetMaps(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string, xplay bool) (*[]ubisoft.Map, error) {
	if xplay && platform != "uplay" {
		platform = "xplay"
	}

	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(mapUri(uuid, platform, xplay))
	requestHeaders(req, auth, false, true)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	//ubisoft returned no content
	if res.StatusCode() == consts.StatusNoContent {
		return nil, nil
	}

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft maps response %v", res.StatusCode())
	}

	var op map[string]interface{}
	de := json.NewDecoder(res.BodyStream()).Decode(&op)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rm1)")
	}

	//dirty - find new solution
	maps := op["profileData"].(map[string]interface{})[uuid].(map[string]interface{})["platforms"].(map[string]interface{})[PlatformModernStats[platform]]

	var result ubisoft.MapsOutputModel
	a, b := utils.Transcode(maps, &result)

	//error encoding response
	if a != nil {
		return nil, errors.New("error decoding response (error code: #rw2)")
	}

	//ubisoft returned no data - handle error by returning nil
	if b != nil {
		return nil, nil
	}

	if len(result.GameModes.Ranked.TeamRoles.All) == 0 {
		return nil, nil
	}

	return &result.GameModes.Ranked.TeamRoles.All, nil
}

func GetOperators(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string, xplay bool) (*[]ubisoft.Operator, error) {
	if xplay && platform != "uplay" {
		platform = "xplay"
	}

	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(operatorUri(uuid, platform, xplay))
	requestHeaders(req, auth, false, true)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	//ubisoft returned no content
	if res.StatusCode() == consts.StatusNoContent {
		return nil, nil
	}

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft operators response %v", res.StatusCode())
	}

	var op map[string]interface{}
	de := json.NewDecoder(res.BodyStream()).Decode(&op)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rw1)")
	}

	//dirty - find new solution
	operators := op["profileData"].(map[string]interface{})[uuid].(map[string]interface{})["platforms"].(map[string]interface{})[PlatformModernStats[platform]]

	var result ubisoft.OperatorOutputModel
	a, b := utils.Transcode(operators, &result)

	//error encoding response
	if a != nil {
		return nil, errors.New("error decoding response (error code: #rw2)")
	}

	//ubisoft returned no data - handle error by returning nil
	if b != nil {
		return nil, nil
	}

	//return single slice of operators
	var output []ubisoft.Operator
	if len(result.GameModes.Ranked.TeamRoles.Attacker) != 0 {
		for _, o := range result.GameModes.Ranked.TeamRoles.Attacker {
			o.OperatorSide = "attacker"
			output = append(output, o)
		}
	}

	if len(result.GameModes.Ranked.TeamRoles.Defender) != 0 {
		for _, o := range result.GameModes.Ranked.TeamRoles.Defender {
			o.OperatorSide = "defender"
			output = append(output, o)
		}
	}

	//ubisoft returned no data - handle error by returning nil
	if len(output) == 0 {
		return nil, nil
	}

	return &output, nil
}

func GetTrends(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string, xplay bool) (*ubisoft.TrendOutput, error) {
	if xplay && platform != "uplay" {
		platform = "xplay"
	}

	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(trendsUri(uuid, platform, xplay))
	requestHeaders(req, auth, false, true)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	//ubisoft returned no content
	if res.StatusCode() == consts.StatusNoContent {
		return nil, nil
	}

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft trends response %v", res.StatusCode())
	}

	var op map[string]interface{}
	de := json.NewDecoder(res.BodyStream()).Decode(&op)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rt1)")
	}

	//dirty - find new solution
	trends := op["profileData"].(map[string]interface{})[uuid].(map[string]interface{})["platforms"].(map[string]interface{})[PlatformModernStats[platform]].(map[string]interface{})["gameModes"].(map[string]interface{})["ranked"].(map[string]interface{})["teamRoles"].(map[string]interface{})["all"].([]interface{})

	if len(trends) == 0 {
		return nil, errors.New("failed to receive ubisoft response")
	}

	trendsTwo := trends[0].(map[string]interface{})
	trendsOutputArray := make([]ubisoft.TrendTypeOutput, 0, 12)

	for _, trend := range []string{"winLossRatio", "killDeathRatio", "headshotAccuracy", "killsPerRound", "roundsWithAKill", "roundsWithMultiKill", "roundsWithOpeningKill", "roundsWithOpeningDeath", "roundsWithKOST", "roundsSurvived", "ratioTimeAlivePerMatch", "distancePerRound"} {
		trendM := trendsTwo[trend].(map[string]interface{})

		trendMap := trendM["trend"].(map[string]interface{})
		actualMap := trendM["actuals"].(map[string]interface{})
		trendArray := make([]float64, 0, len(trendMap))
		actualArray := make([]float64, 0, len(actualMap))

		for i := 1; i <= len(trendMap); i++ {
			trendArray = append(trendArray, trendMap[strconv.Itoa(i)].(float64))
		}

		for i := 1; i <= len(actualMap); i++ {
			actualArray = append(actualArray, actualMap[strconv.Itoa(i)].(float64))
		}

		trendTypeOutput := ubisoft.TrendTypeOutput{
			Name:    trend,
			High:    trendM["high"].(float64),
			Average: trendM["average"].(float64),
			Low:     trendM["low"].(float64),
			Actuals: actualArray,
			Trend:   trendArray,
		}

		trendsOutputArray = append(trendsOutputArray, trendTypeOutput)
	}

	return &ubisoft.TrendOutput{MovingPoints: trendsTwo["movingPoints"].(float64), Trends: trendsOutputArray}, nil
}

func GetSummary(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string, xplay bool) (*ubisoft.Summary, error) {
	if xplay && platform != "uplay" {
		platform = "xplay"
	}

	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(summaryUri(uuid, platform, xplay))
	requestHeaders(req, auth, false, true)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	//ubisoft returned no content
	if res.StatusCode() == consts.StatusNoContent {
		return nil, nil
	}

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft summaries response %v", res.StatusCode())
	}

	var op map[string]interface{}
	de := json.NewDecoder(res.BodyStream()).Decode(&op)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rs1)")
	}

	//dirty - find new solution
	summary := op["profileData"].(map[string]interface{})[uuid].(map[string]interface{})["platforms"].(map[string]interface{})[PlatformModernStats[platform]]

	var result ubisoft.SummaryOutputModel
	a, b := utils.Transcode(summary, &result)

	//error encoding response
	if a != nil {
		return nil, errors.New("error decoding response (error code: #rw2)")
	}

	//ubisoft returned no data - handle error by returning nil
	if b != nil {
		return nil, nil
	}

	if len(result.GameModes.Ranked.TeamRoles.All) == 0 {
		return nil, nil
	}

	return &result.GameModes.Ranked.TeamRoles.All[0], nil
}
