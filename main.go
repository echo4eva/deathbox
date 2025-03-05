package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
)

type Application struct {
	e   *echo.Echo
	rdb *redis.Client
}

func main() {
	e := echo.New()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	app := Application{
		e:   e,
		rdb: rdb,
	}
	app.InitRouting(e)

	ctx := context.Background()

	go func() {
		psub := app.rdb.Subscribe(ctx, "__keyevent@0__:expired")
		defer psub.Close()

		ch := psub.Channel()
		for msg := range ch {
			expiredKey := msg.Payload
			fmt.Printf("Key expired: %s\n", expiredKey)
		}
	}()

	app.e.Logger.Fatal(e.Start(":8080"))
}
