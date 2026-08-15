package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"

	ardan "github.com/ardanlabs/conf"
	"github.com/bvv_utils/goutils/config"
	"github.com/bvv_utils/goutils/mistake"
)

const ttrue = "true"

type ProjectPsqlDb struct {
	*sql.DB
}

// these methods are part of general ProjectDbI interface

func (p *ProjectPsqlDb) GetVersion() string {
	var version string
	err := p.QueryRow("select version()").Scan(&version)
	if err != nil {
		return "psql version unknown"
	}
	return fmt.Sprintf("psql version %s", version)
}

func (p *ProjectPsqlDb) CloseDb() error {
	return p.DB.Close()
}

func (p *ProjectPsqlDb) ConfigureDb() error {
	var d config.PsqlDbConfig
	err := ardan.Parse(os.Args, "", &d)
	if err != nil {
		return err
	}
	p.DB, err = newDb(
		d.Dsn,
		d.SslEnvRequired,
		d.SslEnvModeType,
		d.SslModeDisable)
	if err != nil {
		return err
	}
	if a := d.MaxConn; a > 0 {
		p.DB.SetMaxOpenConns(a)
	}
	return nil
}

func newDb(dsn, sslEnvRequired, sslEnvModeType, sslModeDisable string) (*sql.DB, error) {
	dsn, err := prepareDSN(dsn, sslEnvRequired, sslEnvModeType, sslModeDisable)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", mistake.ErrDbPing, err)
	}

	c, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", mistake.ErrDbConnect, err)
	}

	err = c.Ping()
	if err != nil {
		return nil, fmt.Errorf("%v: %v", mistake.ErrDbConnect, err)
	}

	return c, nil
}

func prepareDSN(dsn string,
	sslEnvRequiredValue, sslEnvModeTypeValue, sslModeDisable any,
) (string, error) {
	defaultTo := func(envName string, envValue interface{}, defaultValue string) error {
		if envValue == "" {
			err := os.Setenv(envName, defaultValue)
			if err != nil {
				return fmt.Errorf("SSL connect erro %v", err)
			}
		}
		return nil
	}
	if err := defaultTo(SslEnvRequired, sslEnvRequiredValue, sslEnvRequiredDefault); err != nil {
		return "", err
	}

	if err := defaultTo(SslEnvModeType, sslEnvModeTypeValue, SslModeDisableDefault); err != nil {
		return "", err
	}

	sslRequired := os.Getenv(SslEnvRequired)
	envSslModeType := os.Getenv(SslEnvModeType)
	if sslRequired == ttrue && envSslModeType == sslModeDisable {
		return "", fmt.Errorf("SSL connect error: %v", errSSLRequired)
	}

	if sslRequired == ttrue && (envSslModeType == sslModeVerifyCa || envSslModeType == sslModeVerifyFull) {
		err := checkDsnSSL(dsn)
		if err != nil {
			return "", fmt.Errorf("SSL connect error %v", err)
		}
	}
	if !strings.Contains(dsn, "sslmode=") {
		return fmt.Sprintf("%v sslmode=%s", dsn, envSslModeType), nil
	} else {
		return dsn, nil
	}
}
