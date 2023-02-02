package ubisoft

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/eliassebastian/r6index-api/pkg/auth"
	ubisoft "github.com/eliassebastian/r6index-api/pkg/ubisoft/models"
	"github.com/eliassebastian/r6index-api/pkg/utils"
)

//import "fmt"

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

	var profile ubisoft.ProfileModel
	de := json.NewDecoder(res.BodyStream()).Decode(&profile)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rpp1)")
	}

	if len(profile.Profiles) == 0 {
		return nil, errors.New("profile not found")
	}

	if profile.Profiles[0].UserID == "" {
		return nil, errors.New("ubisoft server error - userid not found")
	}

	return &profile.Profiles[0], nil
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

func GetRankedTwo(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string) (*[]ubisoft.RankedTwoOutputModel, error) {
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
		})
	}

	return &output, nil
}

func GetWeapons(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string, xplay bool) (*ubisoft.WeaponTeamRoles, error) {
	if xplay && platform != "uplay" {
		platform = "xplay"
	}

	platform = PlatformModernStats[platform]

	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(weaponsUri(uuid, platform, -120, xplay))
	requestHeaders(req, auth, false, true)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft response %v", res.StatusCode())
	}

	var op map[string]interface{}
	de := json.NewDecoder(res.BodyStream()).Decode(&op)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rw1)")
	}

	//dirty - find new solution
	weapons := op["profileData"].(map[string]interface{})[uuid].(map[string]interface{})["platforms"].(map[string]interface{})[platform]

	var result ubisoft.WeaponsOutputModel
	a, b := utils.Transcode(weapons, &result)

	if a != nil || b != nil {
		return nil, errors.New("error decoding response (error code: #rw2)")
	}

	return &result.GameModes.Ranked.TeamRoles, nil
}

func GetMaps(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string, xplay bool) (*ubisoft.MapsTeamRoles, error) {
	if xplay && platform != "uplay" {
		platform = "xplay"
	}

	platform = PlatformModernStats[platform]

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

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft response %v", res.StatusCode())
	}

	var op map[string]interface{}
	de := json.NewDecoder(res.BodyStream()).Decode(&op)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rm1)")
	}

	//dirty - find new solution
	maps := op["profileData"].(map[string]interface{})[uuid].(map[string]interface{})["platforms"].(map[string]interface{})[platform]

	var result ubisoft.MapsOutputModel
	a, b := utils.Transcode(maps, &result)

	if a != nil || b != nil {
		return nil, errors.New("error decoding response (error code: #rm2)")
	}

	return &result.GameModes.Ranked.TeamRoles, nil
}

func GetOperators(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string, xplay bool) (*ubisoft.OperatorTeamRoles, error) {
	if xplay && platform != "uplay" {
		platform = "xplay"
	}

	platform = PlatformModernStats[platform]

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

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft response %v", res.StatusCode())
	}

	var op map[string]interface{}
	de := json.NewDecoder(res.BodyStream()).Decode(&op)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rw1)")
	}

	//dirty - find new solution
	operators := op["profileData"].(map[string]interface{})[uuid].(map[string]interface{})["platforms"].(map[string]interface{})[platform]

	var result ubisoft.OperatorOutputModel
	a, b := utils.Transcode(operators, &result)

	if a != nil || b != nil {
		return nil, errors.New("error decoding response (error code: #rw2)")
	}

	return &result.GameModes.Ranked.TeamRoles, nil
}

func GetTrends(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string, xplay bool) (*ubisoft.TrendOutput, error) {
	if xplay && platform != "uplay" {
		platform = "xplay"
	}

	platform = PlatformModernStats[platform]

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

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft response %v", res.StatusCode())
	}

	var op map[string]interface{}
	de := json.NewDecoder(res.BodyStream()).Decode(&op)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rt1)")
	}

	//dirty - find new solution
	//trends := op["profileData"].(map[string]interface{})[uuid].(map[string]interface{})["platforms"].(map[string]interface{})[platform].(map[string]interface{})["gameModes"].(map[string]interface{})["ranked"].(map[string]interface{})["teamRoles"].(map[string]interface{})["all"].([]interface{})
	trends := op["profileData"].(map[string]interface{})[uuid].(map[string]interface{})["platforms"].(map[string]interface{})[platform].(map[string]interface{})["gameModes"].(map[string]interface{})["ranked"].(map[string]interface{})["teamRoles"].(map[string]interface{})["all"].([]interface{})

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

func GetSummary(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid, platform string, xplay bool) (*ubisoft.SummaryTeamRoles, error) {
	if xplay && platform != "uplay" {
		platform = "xplay"
	}

	platform = PlatformModernStats[platform]

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

	if res.StatusCode() != consts.StatusOK {
		return nil, fmt.Errorf("failed to receive ubisoft response %v", res.StatusCode())
	}

	var op map[string]interface{}
	de := json.NewDecoder(res.BodyStream()).Decode(&op)
	if de != nil {
		return nil, errors.New("error decoding response (error code: #rs1)")
	}

	//dirty - find new solution
	summary := op["profileData"].(map[string]interface{})[uuid].(map[string]interface{})["platforms"].(map[string]interface{})[platform]

	var result ubisoft.SummaryOutputModel
	a, b := utils.Transcode(summary, &result)

	if a != nil || b != nil {
		return nil, errors.New("error decoding response (error code: #rs2)")
	}

	return &result.GameModes.Ranked.TeamRoles, nil
}
