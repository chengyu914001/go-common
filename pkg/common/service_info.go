package common

import (
	"os"
)

var (
	serviceName string
)

func init() {
	serviceName = os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		panic("service name not set")
	}
}

func GetServiceName() string {
	return serviceName
}
