package id

import (
	"github.com/nrednav/cuid2"
	"github.com/oklog/ulid/v2"
)

func GenerateInnerID() string {
	return ulid.Make().String()
}

func GeneratePublicID() string {
	return cuid2.Generate()
}
