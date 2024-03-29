package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/client/retry"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/eliassebastian/r6index-api/cmd/api/middleware"
	"github.com/eliassebastian/r6index-api/cmd/api/models"
	"github.com/eliassebastian/r6index-api/pkg/auth"
	"github.com/eliassebastian/r6index-api/pkg/cache"
	"github.com/eliassebastian/r6index-api/pkg/meili"
	"github.com/eliassebastian/r6index-api/pkg/rabbitmq"
	"github.com/eliassebastian/r6index-api/pkg/utils"
)

type serverConfig struct {
	Authentication *auth.AuthStore
	Rabbit         *rabbitmq.RabbitConsumer
	Client         *client.Client
	Cache          *cache.CacheStore
	DB             *meili.MeiliSearchStore
}

//go:generate go run github.com/shamaton/msgpackgen
func main() {
	// Register generated resolver from msgpackgen
	models.RegisterGeneratedResolver()

	h := server.Default(
		server.WithHostPorts(utils.GetEnv("API_URL", "127.0.0.1:8080")),
		server.WithIdleTimeout(30*time.Second),
		server.WithExitWaitTime(5*time.Second),
	)

	origin := "*"
	if utils.GetEnv("ENV", "dev") == "prod" {
		origin = "https://r6index.app"
	}

	// Register cors middleware
	h.Use(middleware.Cors(origin))

	c, err := client.NewClient(
		client.WithResponseBodyStream(true),
		client.WithDialTimeout(1*time.Second),
		client.WithDialer(standard.NewDialer()),
		client.WithRetryConfig(
			retry.WithInitDelay(1*time.Second),
			retry.WithMaxAttemptTimes(5),
			retry.WithDelayPolicy(retry.BackOffDelayPolicy),
		),
	)

	if err != nil {
		log.Println(err.Error())
		return
	}

	c.SetRetryIfFunc(func(req *protocol.Request, resp *protocol.Response, err error) bool {
		return resp.StatusCode() != 200 || err != nil
	})

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	auth := auth.New()
	redis, err := cache.New(ctx)
	if err != nil {
		log.Println(err.Error())
		return
	}

	rabbit, err := rabbitmq.New(auth)
	if err != nil {
		log.Println(err.Error())
		return
	}

	m, err := meili.New()
	if err != nil {
		log.Println(err.Error())
		return
	}

	sc := &serverConfig{
		Authentication: auth,
		Client:         c,
		Rabbit:         rabbit,
		Cache:          redis,
		DB:             m,
	}

	routes(h, sc)

	// graceful shutdown function
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		rabbit.Close()

		stop()
		<-ctx.Done()
	})

	go rabbit.Consume(ctx)
	h.Spin()
}
