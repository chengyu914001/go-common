package sysconst

import (
	"os"
)

var serviceName string

func init() {
	var err error
	serviceName, err = os.Hostname()
	if err != nil {
		panic(err)
	}
}

func GetServiceName() string {
	return serviceName
}
