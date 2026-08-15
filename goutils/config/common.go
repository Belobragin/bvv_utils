package config

import (
	"slices"
	"strconv"

	"github.com/bvv_utils/goutils/mistake"
)

type StandardConfig struct {
	App struct {
		// PromoHost string `conf:"env:PROMO_HOST"`
		PromoPort string `conf:"env:PROMO_PORT"`
	}
	Db struct {
		Dsn string `conf:"mask,env:DB_DSN"`
	}
	/* LogLevel can be:
	 "DEBUG", "WARN", "ERROR", "PANIC", "DPANIC", "FATAL"
	- goto https://pkg.go.dev/go.uber.org/zap#pkg-constants
	*/
	LogLevel string `conf:"default:INFO,env:LOGGER_LEVEL"`
}

func (s *StandardConfig) GetApiPort() string {
	return s.App.PromoPort
}
func (s *StandardConfig) GetLogLevel() string {
	return s.LogLevel
}

func (c *StandardConfig) ValidateStandardConfig() error {
	if i, err := strconv.Atoi(c.GetApiPort()); err != nil {
		return err
	} else if i < 1001 || i > 64000 {
		return mistake.ErrIncorrectApiPortValue
	}
	if !slices.Contains([]string{"DEBUG", "WARN", "ERROR", "PANIC", "DPANIC", "FATAL"}, c.GetLogLevel()) {
		return mistake.ErrIncorrectLogLevel
	}
	return nil
}

// use this structure for parallel database access patterns:
type ParallelConfig struct {
	App struct {
		// PromoHost string `conf:"env:PROMO_HOST"`
		PromoPort string `conf:"env:PROMO_PORT"`
	}
	Db struct {
		Dsn     string `conf:"mask,env:DB_DSN"`
		MaxConn int    `conf:"default:2,env:SERVICE_MAX_DB_CONN"`
	}
	/* LogLevel can be:
	 "DEBUG", "WARN", "ERROR", "PANIC", "DPANIC", "FATAL"
	- goto https://pkg.go.dev/go.uber.org/zap#pkg-constants
	*/
	LogLevel string `conf:"default:INFO,env:LOGGER_LEVEL"`
}
