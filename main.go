package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/labstack/echo"
	"github.com/michimani/gotwi"
	"github.com/redis/go-redis/v9"
)

var (
	e      *echo.Echo
	rdb    *redis.Client
	logger *slog.Logger
)

func main() {
	e = echo.New()

	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()

	initRouting()

	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv("TWT_ACCESS_TOKEN"),
		OAuthTokenSecret:     os.Getenv("TWT_ACCESS_TOKEN_SECRET"),
	}

	c, err := gotwi.NewClient(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		psub := rdb.Subscribe(ctx, "__keyevent@0__:expired")
		defer psub.Close()

		ch := psub.Channel()
		for msg := range ch {
			expiredKey := msg.Payload
			logger.Info("key expired", "key", expiredKey)
			sendTweet(c, "test")
		}
	}()

	e.Logger.Fatal(e.Start(":8080"))
}
