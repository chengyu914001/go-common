package common

import (
	"os"
)

type EnvModeType byte

const (
	ENV_MODE_LOCAL = iota + 1
	ENV_MODE_DEV
	ENV_MODE_PROD
)

var envMode EnvModeType

func init() {
	switch os.Getenv("ENV_MODE") {
	case "local":
		envMode = ENV_MODE_LOCAL
	case "dev":
		envMode = ENV_MODE_DEV
	case "prod":
		envMode = ENV_MODE_PROD
	default:
		panic("invalid env mode")
	}
}

func GetEnvMode() EnvModeType {
	return envMode
}
