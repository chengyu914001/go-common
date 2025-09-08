package passord

import (
	"sync"
)

const (
	memory      uint32 = 65536
	time        uint32 = 1
	parallelism uint8  = 2
	saltLen     uint8  = 16
	keyLen      uint32 = 32
)

var (
	lock = sync.Mutex{}
)
