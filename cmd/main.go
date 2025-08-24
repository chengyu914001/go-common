package main

import (
	"context"

	"github.com/chengyu914001/go-common/pkg/log"
)

func main() {
	ctx := context.Background()
	ctx = log.SetTraceID(ctx, "123")
	log.Info(ctx, "Hello, World!")
}
