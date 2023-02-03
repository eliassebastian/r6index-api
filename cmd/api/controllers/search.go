package controllers

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/eliassebastian/r6index-api/cmd/api/responses"
	"github.com/eliassebastian/r6index-api/pkg/meili"
	"github.com/eliassebastian/r6index-api/pkg/utils"
)

type SearchController struct {
	db *meili.MeiliSearchStore
}

func NewSearchController(db *meili.MeiliSearchStore) *SearchController {
	return &SearchController{
		db: db,
	}
}

func (sc *SearchController) RequestHandler(ctx context.Context, c *app.RequestContext) {
	// ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	// defer cancel()

	startTime := time.Now()
	platform := c.Query("p")
	query := c.Query("q")

	if !utils.IsValidPlatform(platform) {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, "invalid platform provided"))
		return
	}

	if query == "" {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, "incomplete query prompt provided"))
		return
	}

	resp, err := sc.db.Search(platform, query)
	if err != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, err.Error()))
		return
	}

	c.JSON(consts.StatusOK, responses.Success(startTime, resp))
}
