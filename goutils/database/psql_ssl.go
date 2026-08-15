package database

import (
	"fmt"
	"strings"
)

const (
	SslEnvRequired        = "POSTGRESQL_SSL_REQUIRED"
	sslEnvRequiredDefault = "false"
	SslEnvModeType        = "POSTGRESQL_SSLMODE_TYPE"
	sslEnvKey             = "sslkey"
	sslEnvCert            = "sslcert"
	sslEnvRootCert        = "sslrootcert"
	sslModeRequire        = "require"
	sslModeVerifyCa       = "verify-ca"
	sslModeVerifyFull     = "verify-full"
	SslModeDisableDefault = "disable"
)

var (
	errSSLModeType = fmt.Errorf("dsn string, must contain %s, %s, %s",
		sslEnvKey, sslEnvCert, sslEnvRootCert)
	errSSLRequired = fmt.Errorf("environment \"%s\" must be one of value: %s, %s, %s",
		SslEnvModeType, sslModeRequire, sslModeVerifyCa, sslModeVerifyFull)
)

func checkDsnSSL(dsn string) error {
	if !strings.Contains(dsn, sslEnvKey) && !strings.Contains(dsn, sslEnvCert) && !strings.Contains(dsn, sslEnvRootCert) {
		return errSSLModeType
	}
	return nil
}
