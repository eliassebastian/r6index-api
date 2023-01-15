package ubisoft

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/eliassebastian/r6index-api/pkg/auth"
	ubisoft "github.com/eliassebastian/r6index-api/pkg/ubisoft/models"
)

//import "fmt"

func GetPlayerProfile(ctx context.Context, client client.Client, auth *auth.UbisoftSession, name, uuid, platform string) (*ubisoft.Profile, error) {
	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(profileUri(name, uuid, platform))
	requestHeaders(req, auth, false)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	var profile ubisoft.ProfileModel
	de := json.NewDecoder(res.BodyStream()).Decode(&profile)
	if de != nil {
		return nil, errors.New("error decoding response")
	}

	if len(profile.Profiles) == 0 {
		return nil, errors.New("profile not found")
	}

	return &profile.Profiles[0], nil
}

func GetXpAndLevel(ctx context.Context, client client.Client, auth *auth.UbisoftSession, uuid string) (*ubisoft.XpAndLevel, error) {
	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(xpUri(uuid))
	requestHeaders(req, auth, true)
	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	var xpLevel ubisoft.XpAndLevel
	de := json.NewDecoder(res.BodyStream()).Decode(&xpLevel)
	if de != nil {
		return nil, errors.New("error decoding response")
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
	requestHeaders(req, auth, false)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return nil, err
	}

	var rankedJson ubisoft.RankedModel
	de := json.NewDecoder(res.BodyStream()).Decode(&rankedJson)
	if de != nil {
		return nil, errors.New("error decoding response")
	}

	validSeasons := rankedJson.SeasonsPlayerSkillRecords[5:]
	//log.Println(validSeasons)

	var output []ubisoft.RankedOutputModel
	for _, season := range validSeasons {
		//log.Println(season)
		var regions []ubisoft.RankedSeason

		for _, region := range season.RegionsPlayerSkillRecords {
			if len(region.BoardsPlayerSkillRecords) != 0 {
				regionInfo := region.BoardsPlayerSkillRecords[0].Seasons[0]
				if regionInfo.Wins+regionInfo.Losses+regionInfo.Abandons != 0 {
					regions = append(regions, regionInfo)
				}
			}
		}

		if len(regions) != 0 {
			seasonOut := ubisoft.RankedOutputModel{
				SeasonID:  season.SeasonID,
				SeasonTag: seasons[season.SeasonID].Code,
				Regions:   regions,
			}
			output = append(output, seasonOut)
		}
	}

	return &output, nil
}
