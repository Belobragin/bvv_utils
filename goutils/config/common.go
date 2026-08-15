package config

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
