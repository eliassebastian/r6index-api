package controllers

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/eliassebastian/r6index-api/cmd/api/models"
	"github.com/eliassebastian/r6index-api/cmd/api/responses"
	"github.com/eliassebastian/r6index-api/cmd/api/validation"
	"github.com/eliassebastian/r6index-api/pkg/auth"
	"github.com/eliassebastian/r6index-api/pkg/cache"
	"github.com/eliassebastian/r6index-api/pkg/ubisoft"
	"golang.org/x/sync/errgroup"
)

type IndexController struct {
	auth   *auth.AuthStore
	client *client.Client
	cache  *cache.CacheStore
}

func NewIndexController(a *auth.AuthStore, c *client.Client, cs *cache.CacheStore) *IndexController {
	return &IndexController{
		auth:   a,
		client: c,
		cache:  cs,
	}
}

func (ic *IndexController) RequestHandler(ctx context.Context, c *app.RequestContext) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	startTime := time.Now()

	platform := c.PostForm("platform")
	name := c.PostForm("name")
	uuid := c.PostForm("id")

	if err := validation.All(platform, name, uuid); err != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, err.Error()))
		return
	}

	//fetch ubisoft session
	us := ic.auth.Read()
	profile, err := ubisoft.GetPlayerProfile(ctx, *ic.client, us, name, uuid, platform)
	if err != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, err.Error()))
		return
	}

	output := &models.Player{
		ProfileId:  profile.ProfileID,
		UserId:     profile.UserID,
		Nickname:   profile.NameOnPlatform,
		Platform:   profile.PlatformType,
		LastUpdate: startTime.UTC(),
		Aliases: &[]models.Alias{{
			Name: profile.NameOnPlatform,
			Date: startTime.UTC(),
		}},
	}

	currentAlias := new(models.AliasCache)
	ce := ic.cache.Once(profile.ProfileID, currentAlias, &models.AliasCache{Name: profile.NameOnPlatform})
	//ce := ic.cache.Set(ctx, profile.ProfileID, profile.NameOnPlatform, 1*time.Hour)
	if ce != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, ce.Error()))
		return
	}

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		xp, err := ubisoft.GetXpAndLevel(ctx, *ic.client, us, uuid)

		output.Xp = xp.Xp
		output.Level = xp.Level

		return err
	})

	group.Go(func() error {
		s, err := ubisoft.GetRankedOne(ctx, *ic.client, us, uuid, platform)

		output.RankedOne = s

		return err
	})

	if err := group.Wait(); err != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, err.Error()))
		return
	}

	c.JSON(consts.StatusOK, responses.Success(startTime, output))

	//fetch different stats
	//insert into db
	//convert to readable json
	//return response
}
