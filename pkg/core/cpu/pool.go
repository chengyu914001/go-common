package cpu

import (
	"context"
	"runtime"
)

var cpuPool = make(chan struct{}, runtime.NumCPU())

func DoLimitRun(ctx context.Context, f func() error) error {
	select {
	case cpuPool <- struct{}{}:
		defer func() {
			<-cpuPool
		}()
		return f()
	case <-ctx.Done():
		return ctx.Err()
	}
}
