package cpu

import (
	"context"
	"runtime"
)

var running = make(chan struct{}, runtime.NumCPU())

func DoLimitRun(ctx context.Context, f func() error) error {
	select {
	case running <- struct{}{}:
		defer func() {
			<-running
		}()
		return f()
	case <-ctx.Done():
		return ctx.Err()
	}
}
