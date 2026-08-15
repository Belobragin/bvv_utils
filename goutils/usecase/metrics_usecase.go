package usecase

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type MetricsStandardUseCaseI interface {
	GetMetricsPort() string
}

// use single register for all custom metrics
type MetricClientI interface {
	// general:
	GetLog() *zap.Logger
	GetCustomRegistry() *prometheus.Registry
}

type MetricsStandardUseCaseRealization struct {
	MetricsPort string
}

func (s *MetricsStandardUseCaseRealization) GetMetricsPort() string {
	return s.MetricsPort
}

type MetricClientStruct struct {
	MetricClientI
}

func (m *MetricClientStruct) CustomMetricsHandler() http.Handler {
	if mm := m.GetCustomRegistry(); mm == nil {
		return nil
	} else {
		return promhttp.HandlerFor(mm, promhttp.HandlerOpts{Registry: mm})
	}
}

func (m *MetricClientStruct) GetStandardMetrics() http.Handler {
	return promhttp.Handler()
}
