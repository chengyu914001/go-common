package id

import (
	"github.com/nrednav/cuid2"
)

func GenerateTraceID() string {
	return cuid2.Generate()
}
