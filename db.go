package main

import (
	"context"
	"time"
)

func registerHeartbeat(ctx context.Context, device string) {
	currentTTL, _ := rdb.TTL(ctx, "heartbeat").Result()
	rdb.Set(ctx, "heartbeat", time.Now(), time.Duration(10000*time.Second))
	newTTL, _ := rdb.TTL(ctx, "heartbeat").Result()
	expires := time.Now().Add(10000 * time.Second)

	logger.Info(
		"heartbeat detected",
		"device", device,
		"current", currentTTL.Seconds(),
		"refresh", newTTL.Seconds(),
		"expires", expires.Format(time.UnixDate),
	)
}
