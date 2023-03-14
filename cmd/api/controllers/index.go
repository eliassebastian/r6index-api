package controllers

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/route/param"
	"github.com/eliassebastian/r6index-api/cmd/api/models"
	"github.com/eliassebastian/r6index-api/cmd/api/responses"
	"github.com/eliassebastian/r6index-api/cmd/api/validation"
	"github.com/eliassebastian/r6index-api/pkg/auth"
	"github.com/eliassebastian/r6index-api/pkg/cache"
	"github.com/eliassebastian/r6index-api/pkg/meili"
	"github.com/eliassebastian/r6index-api/pkg/ubisoft"
	"golang.org/x/sync/errgroup"
)

type IndexController struct {
	auth   *auth.AuthStore
	client *client.Client
	cache  *cache.CacheStore
	db     *meili.MeiliSearchStore
}

func NewIndexController(a *auth.AuthStore, c *client.Client, cs *cache.CacheStore, db *meili.MeiliSearchStore) *IndexController {
	return &IndexController{
		auth:   a,
		client: c,
		cache:  cs,
		db:     db,
	}
}

type requestParams struct {
	Platform string `json:"platform"`
	Name     string `json:"name"`
	Id       string `json:"id"`
}

func (ic *IndexController) RequestHandler(ctx context.Context, c *app.RequestContext) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	startTime := time.Now()
	var req requestParams

	//validate request params
	err := binding.BindAndValidate(&c.Request, &req, param.Params{})
	if err != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(time.Now(), err.Error()))
		return
	}

	if err := validation.All(req.Platform, req.Name, req.Id); err != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, err.Error()))
		return
	}

	//fetch ubisoft session
	us := ic.auth.Read()
	profile, err := ubisoft.GetPlayerProfile(ctx, *ic.client, us, req.Name, req.Id, req.Platform)
	if err != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, err.Error()))
		return
	}

	//check if player exists in db
	var playerFound models.PlayerFound
	pfe := ic.db.FindPlayer(req.Platform, profile.ProfileID, &playerFound)
	//player found
	if pfe == nil {
		c.JSON(consts.StatusOK, responses.Success(startTime, &playerFound))
		return
	}

	playerFound.Id = profile.ProfileID
	playerFound.Name = profile.NameOnPlatform

	output := &models.Player{
		//Id:         profile.ProfileID,
		ProfileId:  profile.ProfileID,
		UserId:     profile.UserID,
		Nickname:   profile.NameOnPlatform,
		Platform:   profile.PlatformType,
		FirstIndex: startTime.UTC().Unix(),
		LastUpdate: startTime.UTC().Unix(),
		Aliases: &[]models.Alias{{
			Name: profile.NameOnPlatform,
			Date: startTime.UTC(),
		}},
	}

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		xp, err := ubisoft.GetXpAndLevel(ctx, *ic.client, us, profile.ProfileID, req.Platform)

		if err != nil {
			return err
		}

		output.Xp = xp.Xp
		output.Level = xp.Level

		return err
	})

	group.Go(func() error {
		s, err := ubisoft.GetLastSeen(ctx, *ic.client, us, profile.ProfileID, req.Platform)

		if err != nil {
			return err
		}

		output.LastSeeen = s

		return err
	})

	group.Go(func() error {
		s, err := ubisoft.GetRankedTwo(ctx, *ic.client, us, profile.UserID, req.Platform)

		if err != nil {
			return err
		}

		output.RankedTwo = s

		return err
	})

	group.Go(func() error {
		w, err := ubisoft.GetWeapons(ctx, *ic.client, us, profile.UserID, req.Platform, true)

		if err != nil {
			return err
		}

		output.Weapons = w

		return err
	})

	group.Go(func() error {
		w, err := ubisoft.GetMaps(ctx, *ic.client, us, profile.UserID, req.Platform, true)

		if err != nil {
			return err
		}

		output.Maps = w

		return err
	})

	group.Go(func() error {
		w, err := ubisoft.GetOperators(ctx, *ic.client, us, profile.UserID, req.Platform, true)

		if err != nil {
			return err
		}

		output.Operators = w

		return err
	})

	group.Go(func() error {
		w, err := ubisoft.GetTrends(ctx, *ic.client, us, profile.UserID, req.Platform, true)

		if err != nil {
			return err
		}

		output.Trends = w

		return err
	})

	group.Go(func() error {
		w, err := ubisoft.GetSummary(ctx, *ic.client, us, profile.UserID, req.Platform, true)

		if err != nil {
			return err
		}

		output.Summary = w

		return err
	})

	if err := group.Wait(); err != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, err.Error()))
		return
	}

	// _, de := ic.db.DB.Index(platform).UpdateDocuments(output)
	// if de != nil {
	// 	c.JSON(consts.StatusBadRequest, responses.Error(startTime, "internal error (db1)"))
	// 	return
	// }

	//playerFound
	c.JSON(consts.StatusAccepted, responses.Success(startTime, output))
}
