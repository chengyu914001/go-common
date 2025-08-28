package id

import (
	"github.com/nrednav/cuid2"
)

func GetPublicID() string {
	return cuid2.Generate()
}
