package config

type PsqlDbConfig struct {
	Dsn            string `conf:"mask,env:DB_DSN"`
	SslEnvRequired string `conf:"mask,env:POSTGRESQL_SSL_REQUIRED"`
	SslEnvModeType string `conf:"mask,env:POSTGRESQL_SSLMODE_TYPE"`
	SslModeDisable string `conf:"mask,env:POSTGRESQL_SSL_MODEL_DISABLE"`
	MaxConn        int    `conf:"default:2,env:SERVICE_MAX_DB_CONN"`
}
