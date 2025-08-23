package common

type EnvModeType byte

const (
	ENV_MODE_NO_SET EnvModeType = iota
	ENV_MODE_LOCAL
	ENV_MODE_DEV
	ENV_MODE_PROD
)

var env EnvModeType = ENV_MODE_NO_SET

func GetEnvMode() EnvModeType {
	if env == ENV_MODE_NO_SET {
		panic("env mode not set")
	}
	return env
}

func SetEnvMode(mode string) {
	switch mode {
	case "local":
		env = ENV_MODE_LOCAL
	case "dev":
		env = ENV_MODE_DEV
	case "prod":
		env = ENV_MODE_PROD
	}
}
