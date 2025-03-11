package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/labstack/echo"
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

	go func() {
		psub := rdb.Subscribe(ctx, "__keyevent@0__:expired")
		defer psub.Close()

		ch := psub.Channel()
		for msg := range ch {
			expiredKey := msg.Payload
			fmt.Printf("Key expired: %s\n", expiredKey)
		}
	}()

	e.Logger.Fatal(e.Start(":8080"))
}
