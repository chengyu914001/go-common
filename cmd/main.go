package main

import (
	"context"

	"github.com/chengyu914001/go-common/pkg/log"

)


func main() {
	ctx := context.Background()
	log.GetLogger().Info(ctx, "Hello, World!")
}
