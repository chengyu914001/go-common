package id

import (
	"github.com/nrednav/cuid2"
)

func GeneratePrimaryKetID() string {
	return cuid2.Generate()
}
