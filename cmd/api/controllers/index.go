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
	"github.com/eliassebastian/r6index-api/pkg/ubisoft"
)

type IndexController struct {
	auth   *auth.AuthStore
	client *client.Client
}

func NewIndexController(a *auth.AuthStore, c *client.Client) *IndexController {
	return &IndexController{
		auth:   a,
		client: c,
	}
}

func (ic *IndexController) RequestHandler(ctx context.Context, c *app.RequestContext) {
	startTime := time.Now()

	platform := c.PostForm("platform")
	name := c.PostForm("name")
	uuid := c.PostForm("id")

	if err := validation.All(platform, name, uuid); err != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, err.Error()))
		return
	}

	output := &models.Player{}
	//fetch ubisoft session
	us := ic.auth.Read()
	err := ubisoft.GetPlayerProfile(ctx, *ic.client, us, output, name, uuid, platform)
	if err != nil {
		c.JSON(consts.StatusBadRequest, responses.Error(startTime, err.Error()))
		return
	}

	c.JSON(consts.StatusOK, responses.Success(startTime, output))

	//fetch from ubisoft profiles

	//check if valid account
	//fetch different stats
	//insert into db
	//convert to readable json
	//return response
}
