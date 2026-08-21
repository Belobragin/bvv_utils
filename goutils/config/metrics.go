package config

import (
	"strconv"

	"github.com/belobragin/bvv_utils/goutils/mistake"
)

type MetricConfigI interface {
	GetMetricPort() string
}
type MetricConfig struct {
	// port exposed for prometheus metrics
	MetricsPort string `conf:"env:METRICS_PORT"`
}

func (s *MetricConfig) GetMetricPort() string {
	return s.MetricsPort
}

func (c *MetricConfig) ValidateMetricConfig() error {
	if i, err := strconv.Atoi(c.GetMetricPort()); err != nil {
		return err
	} else if i < 1001 || i > 64000 {
		return mistake.ErrIncorrectMetricPortValue
	}
	return nil
}
