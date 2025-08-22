package common

type EnvModeType byte

const (
	NO_SET EnvModeType = 1 << iota
	LOCAL_ENV
	DEV_ENV
	PROD_ENV
)

var env EnvModeType = NO_SET

func GetEnvMode() EnvModeType {
	if env == NO_SET {
		panic("env mode not set")
	}

	return env
}

func SetEnvMode(mode string) {
	switch mode {
	case "local":
		env = LOCAL_ENV
	case "dev":
		env = DEV_ENV
	case "prod":
		env = PROD_ENV
	}
}
