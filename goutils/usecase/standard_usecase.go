package usecase

import (
	"net/http"
	"time"

	"github.com/belobragin/bvv_utils/goutils/config"
	"github.com/belobragin/bvv_utils/goutils/database"

	"go.uber.org/zap"
)

const (
	ServerReadHeaderTimeout = time.Duration(100 * time.Millisecond)
)

type StandardUseCaseI interface {
	GetDebug() bool
	GetDb() *database.ProjectPsqlDb
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

type StandardUseCaseRealization struct {
	log         *zap.Logger
	debug       bool
	db          *database.ProjectPsqlDb
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

func (s *StandardUseCaseRealization) GetDb() *database.ProjectPsqlDb {
	return s.db
}

func (s *StandardUseCaseRealization) GetDebug() bool {
	return s.debug
}

func (s *StandardUseCaseRealization) GetLog() *zap.Logger {
	return s.log
}

func (s *StandardUseCaseRealization) ErrCh() chan error {
	return s.errC
}

func (s *StandardUseCaseRealization) GetServiceName() string {
	return s.serviceName
}

func (s *StandardUseCaseRealization) GetApiPort() string {
	return s.apiPort
}

func (s *StandardUseCaseRealization) GetErrorMap() map[error]int {
	return s.errMap
}

func (s *StandardUseCaseRealization) GetMetricPort() string {
	return s.metricPort
}

func (s *StandardUseCaseRealization) GetAllStopChan() chan struct{} {
	return s.allStopChan
}

func (s *StandardUseCaseRealization) GetAppRouter() http.Handler {
	return s.AppRouter
}

func (s *StandardUseCaseRealization) GetMetricRouter() http.Handler {
	return s.MetricRouter
}

// run server:
func (s *StandardUseCaseRealization) NewApiServer() *http.Server {
	return &http.Server{
		Addr:              ":" + s.GetApiPort(),
		ReadHeaderTimeout: ServerReadHeaderTimeout,
		Handler:           s.GetAppRouter(),
	}
}

// run prometheus metrics server:
func (s *StandardUseCaseRealization) NewMetricServer() *http.Server {
	return &http.Server{
		Addr:              ":" + s.GetMetricPort(),
		ReadHeaderTimeout: ServerReadHeaderTimeout,
		Handler:           s.GetMetricRouter(),
	}
}

func NewStandardUseCase(
	c config.StandardConfigI,
	serviceName string,
	l *zap.Logger,
	ec chan error,
	projectErrMap map[error]int,
) (*StandardUseCaseRealization, error) {
	var u = new(StandardUseCaseRealization)
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

type StandardCorsRealization struct {
	useCors     bool
	allowOrigin string
	maxAge      int
}

func (s *StandardCorsRealization) GetUseCors() bool {
	return s.useCors
}

func (s *StandardCorsRealization) GetAllowOrigin() string {
	return s.allowOrigin
}

func (s *StandardCorsRealization) GetMaxAge() int {
	return s.maxAge
}

func NewStandardCors(
	c config.CorsConfigI,
) (*StandardCorsRealization, error) {
	var u = new(StandardCorsRealization)
	u.useCors = c.GetUseCors()
	u.allowOrigin = c.GetAllowOrigin()
	u.maxAge = c.GetMaxAge()
	return u, nil
}
