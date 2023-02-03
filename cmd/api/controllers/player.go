package controllers

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/eliassebastian/r6index-api/cmd/api/models"
	"github.com/eliassebastian/r6index-api/cmd/api/responses"
	"github.com/eliassebastian/r6index-api/pkg/cache"
	"github.com/eliassebastian/r6index-api/pkg/meili"
	"github.com/eliassebastian/r6index-api/pkg/utils"
)

type PlayerController struct {
	cache *cache.CacheStore
	db    *meili.MeiliSearchStore
}

func NewPlayerController(cs *cache.CacheStore, db *meili.MeiliSearchStore) *PlayerController {
	return &PlayerController{
		cache: cs,
		db:    db,
	}
}

func (pc *PlayerController) RequestHandler(ctx context.Context, c *app.RequestContext) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	startTime := time.Now()
	platform := c.Query("p")
	id := c.Query("id")

	if !utils.IsValidPlatform(platform) {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, "invalid platform provided"))
		return
	}

	if !utils.IsValidUUID(id) {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, "invalid uuid provided"))
		return
	}

	player := new(models.Player)
	err := pc.db.FetchPlayer(platform, id, player)
	if err != nil {
		log.Println(err.Error())
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, "internal db error"))
		return
	}

	ce := pc.cache.SetOnce(ctx, id, &models.ProfileCache{LastUpdate: player.LastUpdate, Aliases: player.Aliases})
	if ce != nil {
		log.Println(ce.Error())
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, "internal cache error"))
		return
	}

	c.JSON(consts.StatusOK, responses.Success(startTime, player))
}
