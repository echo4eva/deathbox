package main

import (
	"context"
	"fmt"
	"time"
)

func (app *Application) DoSomething(ctx context.Context) {
	app.rdb.Set(ctx, "heartbeat", time.Now(), time.Duration(10000*time.Second))
	ttl, _ := app.rdb.TTL(ctx, "heartbeat").Result()
	fmt.Println(ttl.Seconds())
}
