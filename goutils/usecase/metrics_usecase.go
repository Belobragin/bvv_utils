package usecase

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/belobragin/bvv_utils/goutils/metrica"
	"github.com/belobragin/bvv_utils/goutils/mistake"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type BaseMetricI interface {
	GetCounterVec() []*prometheus.CounterVec
	SetStatusCodeApiCallMetrica(metrica.ServiceApiCallMetricaLabelsI) error
}

// use single register for all custom metrics
type MetricI interface {
	BaseMetricI
	GetMetricPort() string
	NewCustomMetricServer(time.Duration) *http.Server
	// private:
	getCustomRegistry() *prometheus.Registry
}

type MetricRealization struct {
	p string
	c []*prometheus.CounterVec
	r *prometheus.Registry
}

func (l *MetricRealization) GetMetricPort() string {
	return l.p
}

func (l *MetricRealization) GetCounterVec() []*prometheus.CounterVec {
	return l.c
}

func (l *MetricRealization) getCustomRegistry() *prometheus.Registry {
	return l.r
}

// run prometheus metrics server:
func (s *MetricRealization) NewCustomMetricServer(timeout time.Duration) *http.Server {
	if s == nil {
		s = new(MetricRealization)
	}
	return &http.Server{
		Addr:              ":" + s.GetMetricPort(),
		ReadHeaderTimeout: timeout,
		Handler:           CustomMetricRouter(s),
	}
}

func (s *MetricRealization) SetStatusCodeApiCallMetrica(l metrica.ServiceApiCallMetricaLabelsI) error {
	u := s.GetCounterVec()
	if len(u) == 0 {
		return mistake.ErrApiMetricVectorNull
	}
	t, e := u[0].GetMetricWithLabelValues(l.GetStatusCodeLabel(), l.GetMethodLabel(), l.GetModellabel())
	if e != nil {
		return mistake.NewAddErr(mistake.ErrMetricaLabel, e)
	}
	t.Inc()
	return nil
}
func NewMetricRealization(
	p string,
	counterVecs ...*prometheus.CounterVec) (*MetricRealization, error) {
	var m = new(MetricRealization)

	if len(p) > 0 {
		m.p = p
	} else {
		return nil, mistake.ErrMetricPort
	}
	m.r = prometheus.NewRegistry()
	for _, v := range counterVecs {
		if !slices.Contains(metrica.AllMetrica, v) {
			return nil, mistake.ErrInvalidMetricRegister
		}
		m.c = append(m.c, v)
		m.r.Register(v)
	}
	return m, nil
}

const (
	MetricPrefix = "metrics"
)

var (
	StandardMetricRoute = fmt.Sprintf("/%s/%s", MetricPrefix, "standard")
	CustomMetricRoute   = fmt.Sprintf("/%s/%s", MetricPrefix, "custom")
)

// base metrics only:
func AddDefaultMetricRoutes(mux *http.ServeMux, metricClient MetricI) {
	mux.Handle(StandardMetricRoute, promhttp.Handler())
}

func DefaultMetricRouter(
	useCase MetricI,
) http.Handler {
	mux := http.NewServeMux()
	AddDefaultMetricRoutes(mux, useCase)

	return mux
}

// base and custom metrics:
func AddCustomMetricRoutes(mux *http.ServeMux, metricClient MetricI) {
	mux.Handle(StandardMetricRoute, promhttp.Handler())
	mux.Handle(
		CustomMetricRoute,
		promhttp.HandlerFor(
			metricClient.getCustomRegistry(),
			promhttp.HandlerOpts{Registry: metricClient.getCustomRegistry()}))
}

func CustomMetricRouter(
	useCase MetricI,
) http.Handler {
	mux := http.NewServeMux()
	AddCustomMetricRoutes(mux, useCase)

	return mux
}
