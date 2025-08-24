package main

import (
	"context"

	"github.com/chengyu914001/go-common/pkg/log"
)

func main() {
	ctx := context.Background()
	ctx = log.SetStrVal(ctx, "trace_id", "123")
	ctx = log.SetStrVal(ctx, "ip", "127.0.0.1")
	ctx = log.SetStrVal(ctx, "user_id", "1")
	log.Error(ctx, "Hello, World!")
}
