package ubisoft

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/eliassebastian/r6index-api/cmd/api/models"
	"github.com/eliassebastian/r6index-api/pkg/auth"
	ubisoft "github.com/eliassebastian/r6index-api/pkg/ubisoft/models"
)

//import "fmt"

func GetPlayerProfile(ctx context.Context, client client.Client, auth *auth.UbisoftSession, output *models.Player, name, uuid, platform string) error {

	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(res)

	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(profileUri(name, uuid, platform))
	requestHeaders(req, auth, false)

	err := client.DoRedirects(ctx, req, res, 1)
	if err != nil {
		return err
	}

	var profile ubisoft.ProfileModel
	de := json.NewDecoder(res.BodyStream()).Decode(&profile)
	if de != nil {
		return errors.New("error decoding response")
	}

	if len(profile.Profiles) == 0 {
		return errors.New("profile not found")
	}

	output.ProfileId = profile.Profiles[0].ProfileID
	output.UserId = profile.Profiles[0].UserID
	output.Nickname = profile.Profiles[0].NameOnPlatform
	output.Platform = profile.Profiles[0].PlatformType
	output.LastUpdate = time.Now().UTC()

	return nil
}
