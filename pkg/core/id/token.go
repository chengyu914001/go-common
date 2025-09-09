package id

import (
	"github.com/nrednav/cuid2"
)

func GenerateTokenID() string {
	return cuid2.Generate()
}
