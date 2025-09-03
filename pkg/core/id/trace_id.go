package id

import (
	"github.com/nrednav/cuid2"
)

func GetTraceID() string {
	return cuid2.Generate()
}
