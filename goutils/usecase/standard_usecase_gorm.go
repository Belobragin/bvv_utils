package usecase

import (
	"net/http"

	"github.com/belobragin/bvv_utils/goutils/config"
	"github.com/belobragin/bvv_utils/goutils/database"

	"go.uber.org/zap"
)

type StandardUseCaseGormI interface {
	GetDebug() bool
	GetDb() *database.ProjectPsqlGormDb
	GetLog() *zap.Logger
	GetApiPort() string
	GetMetricPort() string
	GetServiceName() string
	GetUseCors() bool
	GetAllowOrigin() string
	GetMaxAge() int
	ErrCh() chan error
	GetAllStopChan() chan struct{}
	GetAppRouter() http.Handler
	NewApiServer() *http.Server
	NewMetricServer() *http.Server
	GetErrorMap() map[error]int
}

type StandardUseCaseGormRealization struct {
	log         *zap.Logger
	debug       bool
	db          *database.ProjectPsqlGormDb
	errC        chan error
	serviceName string
	apiPort     string
	metricPort  string
	allStopChan chan struct{}
	errMap      map[error]int
	// AppRouter value is set in each project
	AppRouter http.Handler
	// MetricRouter value is set in each project
	MetricRouter http.Handler
}

func (s *StandardUseCaseGormRealization) GetDb() *database.ProjectPsqlGormDb {
	return s.db
}

func (s *StandardUseCaseGormRealization) GetDebug() bool {
	return s.debug
}

func (s *StandardUseCaseGormRealization) GetLog() *zap.Logger {
	return s.log
}

func (s *StandardUseCaseGormRealization) ErrCh() chan error {
	return s.errC
}

func (s *StandardUseCaseGormRealization) GetServiceName() string {
	return s.serviceName
}

func (s *StandardUseCaseGormRealization) GetApiPort() string {
	return s.apiPort
}

func (s *StandardUseCaseGormRealization) GetErrorMap() map[error]int {
	return s.errMap
}

func (s *StandardUseCaseGormRealization) GetMetricPort() string {
	return s.metricPort
}

func (s *StandardUseCaseGormRealization) GetAllStopChan() chan struct{} {
	return s.allStopChan
}

func (s *StandardUseCaseGormRealization) GetAppRouter() http.Handler {
	return s.AppRouter
}

func (s *StandardUseCaseGormRealization) GetMetricRouter() http.Handler {
	return s.MetricRouter
}

// run server:
func (s *StandardUseCaseGormRealization) NewApiServer() *http.Server {
	return &http.Server{
		Addr:              ":" + s.GetApiPort(),
		ReadHeaderTimeout: ServerReadHeaderTimeout,
		Handler:           s.GetAppRouter(),
	}
}

// run prometheus metrics server:
func (s *StandardUseCaseGormRealization) NewMetricServer() *http.Server {
	return &http.Server{
		Addr:              ":" + s.GetMetricPort(),
		ReadHeaderTimeout: ServerReadHeaderTimeout,
		Handler:           s.GetMetricRouter(),
	}
}

func NewStandardGormUseCase(
	c config.StandardConfigI,
	serviceName string,
	l *zap.Logger,
	ec chan error,
	projectErrMap map[error]int,
) (*StandardUseCaseGormRealization, error) {
	var u = new(StandardUseCaseGormRealization)
	u.log = l
	u.serviceName = serviceName
	u.errC = ec
	u.debug = c.GetDebug()
	u.apiPort = c.GetApiPort()
	u.metricPort = c.GetMetricPort()
	u.allStopChan = make(chan struct{}, 1)
	u.errMap = make(map[error]int)
	u.errMap = projectErrMap
	return u, nil
}
