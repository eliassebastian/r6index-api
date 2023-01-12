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
