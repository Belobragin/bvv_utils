package config

type MetricsConfig struct {
	// port exposed for prometheus metrics
	MetricsPort string `conf:"env:METRICS_PORT"`
}
