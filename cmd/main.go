package main

import (
	"context"
	"time"

	"github.com/chengyu914001/go-common/pkg/log"
)

func main() {
	ctx := context.Background()
	ctx = log.SetStrVal(
		ctx,
		[2]string{"trace_id", "123"},
		[2]string{"ip", "127.0.0.1"},
		[2]string{"user_id", "1"},
	)
	log.Error(ctx, "Hello, World!")

	time.Sleep(100 * time.Millisecond)
}
