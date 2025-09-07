package main

import (
	"context"
	"time"

	"github.com/chengyu914001/go-common/pkg/core/log"
)

func main() {
	ctx := context.Background()
	ctx = log.SetDuration(ctx, "duration", time.Duration(time.Hour))
	ctx = log.SetTime(ctx, "test_time", time.Now())
	log.Info(ctx, "test")
	time.Sleep(10 * time.Millisecond)
}
