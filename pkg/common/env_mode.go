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

var env EnvModeType

func init() {
	switch os.Getenv("ENV_MODE") {
	case "local":
		env = ENV_MODE_LOCAL
	case "dev":
		env = ENV_MODE_DEV
	case "prod":
		env = ENV_MODE_PROD
	default:
		panic("invalid env mode")
	}
}

func GetEnvMode() EnvModeType {
	return env
}
