package id

import (
	"github.com/nrednav/cuid2"
)

func NewTraceID() string {
	return cuid2.Generate()
}