package controllers

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/eliassebastian/r6index-api/cmd/api/models"
	"github.com/eliassebastian/r6index-api/cmd/api/responses"
	"github.com/eliassebastian/r6index-api/pkg/auth"
	"github.com/eliassebastian/r6index-api/pkg/cache"
	"github.com/eliassebastian/r6index-api/pkg/meili"
	"github.com/eliassebastian/r6index-api/pkg/ubisoft"
	"github.com/eliassebastian/r6index-api/pkg/utils"
	"golang.org/x/sync/errgroup"
)

type UpdateController struct {
	auth   *auth.AuthStore
	client *client.Client
	cache  *cache.CacheStore
	db     *meili.MeiliSearchStore
}

func NewUpdateController(a *auth.AuthStore, c *client.Client, cs *cache.CacheStore, db *meili.MeiliSearchStore) *UpdateController {
	return &UpdateController{
		auth:   a,
		client: c,
		cache:  cs,
		db:     db,
	}
}

func (uc *UpdateController) RequestHandler(ctx context.Context, c *app.RequestContext) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	startTime := time.Now()
	platform := c.PostForm("platform")
	uuid := c.PostForm("id")

	if !utils.IsValidPlatform(platform) {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, "invalid platform provided"))
		return
	}

	if !utils.IsValidUUID(uuid) {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, "invalid player uuid provided"))
		return
	}

	profileCache := new(models.ProfileCache)
	err := uc.cache.GetCache(ctx, uuid, profileCache)
	if err != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, "internal cache error"))
		return
	}

	if !(startTime.UTC().Sub(time.Unix(profileCache.LastUpdate, 0)) > 1*time.Hour) {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, "last update occurred too soon"))
		return
	}

	us := uc.auth.Read()
	profile, pe := ubisoft.GetPlayerProfile(ctx, *uc.client, us, "", uuid, platform)
	if pe != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, "ubisoft response error"))
		return
	}

	output := &models.PlayerUpdate{
		ProfileId: profile.ProfileID,
	}

	aliases := *profileCache.Aliases
	lastAlias := aliases[len(aliases)-1]
	if lastAlias.Name != profile.NameOnPlatform && startTime.UTC() != lastAlias.Date {
		aliases = append(aliases, models.Alias{Name: profile.NameOnPlatform, Date: startTime.UTC()})
		output.Aliases = &aliases
		output.Nickname = profile.NameOnPlatform
	}

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		xp, err := ubisoft.GetXpAndLevel(ctx, *uc.client, us, profile.ProfileID, platform)

		if err != nil {
			return err
		}

		output.Xp = xp.Xp
		output.Level = xp.Level

		return err
	})

	group.Go(func() error {
		s, err := ubisoft.GetRankedOne(ctx, *uc.client, us, profile.ProfileID, platform)

		if err != nil {
			return err
		}

		output.RankedOne = s

		return err
	})

	group.Go(func() error {
		s, err := ubisoft.GetRankedTwo(ctx, *uc.client, us, profile.UserID, platform)

		if err != nil {
			return err
		}

		output.RankedTwo = s

		return err
	})

	group.Go(func() error {
		w, err := ubisoft.GetWeapons(ctx, *uc.client, us, profile.UserID, platform, true)

		if err != nil {
			return err
		}

		output.Weapons = w

		return err
	})

	group.Go(func() error {
		w, err := ubisoft.GetMaps(ctx, *uc.client, us, profile.UserID, platform, true)

		if err != nil {
			return err
		}

		output.Maps = w

		return err
	})

	group.Go(func() error {
		w, err := ubisoft.GetOperators(ctx, *uc.client, us, profile.UserID, platform, true)

		if err != nil {
			return err
		}

		output.Operators = w

		return err
	})

	group.Go(func() error {
		w, err := ubisoft.GetTrends(ctx, *uc.client, us, profile.UserID, platform, true)

		if err != nil {
			return err
		}

		output.Trends = w

		return err
	})

	group.Go(func() error {
		w, err := ubisoft.GetSummary(ctx, *uc.client, us, profile.UserID, platform, true)

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

	output.LastUpdate = startTime.UTC().Unix()

	_, de := uc.db.DB.Index(platform).UpdateDocuments(output)
	if de != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, "internal error (db1)"))
		return
	}

	c.JSON(consts.StatusAccepted, responses.Success(startTime, output))
}
