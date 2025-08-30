package cpu

import (
	"context"
	"runtime"
)

var running = make(chan struct{}, runtime.NumCPU())

func DoLimitRun(ctx context.Context, fn func() error) error {
	select {
	case running <- struct{}{}:
		defer func() {
			<-running
		}()
		return fn()
	case <-ctx.Done():
		return ctx.Err()
	}
}
