package main

import (
	"context"

	"github.com/chengyu914001/go-common/pkg/log"
)

func main() {
	ctx := context.Background()
	ctx = log.SetDymanicValues(ctx, map[string]string{"trace_id": "123", "ip": "127.0.0.1", "user_id": "1"})
	log.Error(ctx, "Hello, World!")
}
