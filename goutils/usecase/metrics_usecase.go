package usecase

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/belobragin/bvv_utils/goutils/mistake"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// metrics list
const (
	ApiCallMetric = "status_code_api_call_metric"
)

var (
	NewApiCallMetrica = newApiCallMetrica()
	NewEventMetrica   = newEventMetrica()

	// this variable for check, if metrica ever exist
	AllMetrica = []*prometheus.CounterVec{NewApiCallMetrica, NewEventMetrica}
)

// use single register for all custom metrics
type MetricI interface {
	GetCounterVec() []*prometheus.CounterVec
	SetStatusCodeApiCallMetrica(ServiceApiCallMetricaLabelsI) error
	// private:
	getCustomRegistry() *prometheus.Registry
}

type MetricRealization struct {
	p string
	c []*prometheus.CounterVec
	r *prometheus.Registry
}

func (l *MetricRealization) GetCounterVec() []*prometheus.CounterVec {
	return l.c
}

func (l *MetricRealization) getCustomRegistry() *prometheus.Registry {
	return l.r
}

func (s *MetricRealization) SetStatusCodeApiCallMetrica(l ServiceApiCallMetricaLabelsI) error {
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
	counterVec *prometheus.CounterVec) (MetricI, error) {
	var m = new(MetricRealization)

	if len(p) > 0 {
		m.p = p
	} else {
		return nil, mistake.ErrMetricPort
	}
	m.r = prometheus.NewRegistry()
	if !slices.Contains(AllMetrica, counterVec) {
		return nil, mistake.ErrInvalidMetricRegister
	}
	m.c = append(m.c, counterVec)
	m.r.Register(counterVec)
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

// base and custom metrics:
func AddCustomMetricRoutes(mux *http.ServeMux, metricClient MetricI) {
	mux.Handle(StandardMetricRoute, promhttp.Handler())
	mux.Handle(
		CustomMetricRoute,
		promhttp.HandlerFor(
			metricClient.getCustomRegistry(),
			promhttp.HandlerOpts{Registry: metricClient.getCustomRegistry()}))
}

type ServiceApiCallMetricaLabelsI interface {
	GetModellabel() string
	GetMethodLabel() string
	GetStatusCodeLabel() string
}

type ServiceApiCallMetricaLabels struct {
	Model  string
	Method string
	Code   int
}

func (s *ServiceApiCallMetricaLabels) GetModellabel() string {
	return s.Model
}
func (s *ServiceApiCallMetricaLabels) GetMethodLabel() string {
	return s.Method
}
func (s *ServiceApiCallMetricaLabels) GetStatusCodeLabel() string {
	return fmt.Sprintf("%d", s.Code)
}

// for Kafka, grpc etc. add other coubterVecs:
func newApiCallMetrica() *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_call_status_code_response_total",
			Help: "Number of http status code responses for the method for the model in the service",
		},
		[]string{"code", "method", "model"},
	)
}

type EventMetricaLabelsI interface {
	GetModellabel() string
	GetEventTypeLabel() string
	GetResultLabel() string
}
type EventMetricaLabels struct {
	Model     string
	EventType string
	Result    string
}

func (s *EventMetricaLabels) GetModellabel() string {
	return s.Model
}
func (s *EventMetricaLabels) GetEventTypeLabel() string {
	return s.EventType
}
func (s *EventMetricaLabels) GetResultLabel() string {
	return s.Result
}

func newEventMetrica() *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "broker_event_call_total",
			Help: "Number of broker requests succes/fail",
		},
		[]string{"result", "event_type", "model"},
	)
}
