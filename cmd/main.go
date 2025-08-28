package main

import (
	"context"
	"time"

	"github.com/chengyu914001/go-common/pkg/core/id"
	"github.com/chengyu914001/go-common/pkg/core/log"
)

func main() {
	ctx := context.Background()
	ctx = log.SetStrKetVals(
		ctx,
		[2]string{"trace_id", id.NewTraceID()},
	)
	log.Info(ctx, "Hello, World!")

	time.Sleep(20 * time.Millisecond)
}
