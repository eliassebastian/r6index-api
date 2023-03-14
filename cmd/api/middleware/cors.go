package middleware

import (
	"bytes"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Cors(allowedOrigin string) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Header("Vary", "Origin")
		c.Header("Vary", "Access-Control-Request-Method")

		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			if bytes.Equal(c.Request.Method(), []byte("OPTIONS")) && c.Request.Header.Get("Access-Control-Request-Method") != "" {
				c.Header("Access-Control-Allow-Methods", "OPTIONS")
				c.Header("Access-Control-Allow-Headers", "Content-Type")
				c.Header("Access-Control-Max-Age", "43200")
				c.AbortWithStatus(consts.StatusOK)
				return
			}
		}
	}
}
