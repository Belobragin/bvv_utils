package metrica

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// metrics list
const (
	ApiCallMetric = "status_code_api_call_metric"
)

var (
	NewApiCallMetrica = newApiCallMetrica()
	NewEventMetrica   = newEventMetrica()

	AllMetrica = []*prometheus.CounterVec{NewApiCallMetrica, NewEventMetrica}
)

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
			Name: "kafka_event_call_total",
			Help: "Number of nats requests succes/fail",
		},
		[]string{"result", "event_type", "model"},
	)
}
