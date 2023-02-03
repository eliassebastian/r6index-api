package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/eliassebastian/r6index-api/cmd/api/controllers"
)

func routes(s *server.Hertz, sc *serverConfig) {
	ic := controllers.NewIndexController(sc.Authentication, sc.Client, sc.Cache, sc.DB)
	uc := controllers.NewUpdateController(sc.Authentication, sc.Client, sc.Cache, sc.DB)
	pc := controllers.NewPlayerController(sc.Cache, sc.DB)
	ssc := controllers.NewSearchController(sc.DB)

	h := s.Group("/v1")

	// test v1/ping
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "PONG")
	})

	h.POST("/index", ic.RequestHandler)
	h.GET("/player", pc.RequestHandler)
	h.GET("/search", ssc.RequestHandler)
	h.POST("/update", uc.RequestHandler)
}
