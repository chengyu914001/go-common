package sysconst

import (
	"os"
)

var (
	serviceName string
	issuerName  string
)

func init() {
	var err error
	serviceName, err = os.Hostname()
	if err != nil {
		panic(err)
	}

	issuerName = os.Getenv("ISSUER_NAME")
	if issuerName == "" {
		panic("ISSUER_NAME is not set")
	}
}

func GetServiceName() string {
	return serviceName
}

func GetIssuerName() string {
	return issuerName
}
