package main

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/eliassebastian/r6index-api/cmd/api/controllers"
	"github.com/eliassebastian/r6index-api/cmd/api/responses"
	"github.com/eliassebastian/r6index-api/pkg/utils"
)

func routes(s *server.Hertz, sc *serverConfig) {
	ic := controllers.NewIndexController(sc.Authentication, sc.Client, sc.Cache, sc.DB)

	h := s.Group("/v1")

	// test v1/ping
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "PONG")
	})

	// query database for player
	h.GET("/find", func(ctx context.Context, c *app.RequestContext) {

		startTime := time.Now()

		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		playerName := c.Query("name")
		playerId := c.Query("uid")
		//platform := c.Query("platform")

		//
		if playerName != "" && playerId != "" {
			c.JSON(consts.StatusBadRequest, responses.Error(startTime, "name and uid params both used"))
		}

		if playerName == "" && playerId == "" {
			c.JSON(consts.StatusBadRequest, responses.Error(startTime, "both name and uid params are empty"))
		}

		if playerId != "" && !utils.IsValidUUID(playerId) {
			c.JSON(consts.StatusBadRequest, responses.Error(startTime, "invalid player uuid provided"))
		}

	})

	h.POST("/index", ic.RequestHandler)

	//h.GET("/player")

	//h.POST("/update")
}
