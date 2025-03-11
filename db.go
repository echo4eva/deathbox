package main

import (
	"context"
	"time"
)

func registerHeartbeat(ctx context.Context, device string) {
	duration := time.Duration(7 * 24 * time.Hour)

	currentTTL, _ := rdb.TTL(ctx, "heartbeat").Result()
	rdb.Set(ctx, "heartbeat", time.Now(), duration)
	newTTL, _ := rdb.TTL(ctx, "heartbeat").Result()
	expires := time.Now().Add(duration)

	logger.Info(
		"heartbeat detected",
		"device", device,
		"current", currentTTL.Seconds(),
		"refresh", newTTL.Seconds(),
		"expires", expires.Format(time.UnixDate),
	)
}
