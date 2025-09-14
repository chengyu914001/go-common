package id

import (
	"github.com/nrednav/cuid2"
)

func GeneratePrimaryKeyID() string {
	return cuid2.Generate()
}
