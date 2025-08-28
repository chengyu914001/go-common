package cpu

import (
	"context"
	"runtime"
)

var semphone = make(chan struct{}, runtime.NumCPU())

func DoLimitRun(ctx context.Context, f func() error) error {
	select {
	case semphone <- struct{}{}:
		defer func() {
			<-semphone
		}()
		return f()
	case <-ctx.Done():
		return ctx.Err()
	}
}
